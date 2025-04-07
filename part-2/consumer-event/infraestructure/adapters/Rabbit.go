package adapters

import (
	"consumer-event/domain"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/streadway/amqp"
)

type RabbitMQRepository struct {
	conn *amqp.Connection
}

// Constructor para RabbitMQRepository
func NewRabbitMQRepository() (*RabbitMQRepository, error) {
	conn, err := amqp.Dial("amqp://guest:guest@3.218.163.41:80/")
	if err != nil {
		return nil, fmt.Errorf("Error conectando a RabbitMQ: %v", err)
	}

	return &RabbitMQRepository{conn: conn}, nil
}

func (r *RabbitMQRepository) ConsumeTransaction() error {
	conn, err := amqp.Dial("amqp://guest:guest@3.218.163.41:80/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	// Declarar la cola de transacciones
	queue, err := ch.QueueDeclare("transactions", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Escuchando transacciones en RabbitMQ...")

	for msg := range msgs {
		var transaction domain.Transaction
		if err := json.Unmarshal(msg.Body, &transaction); err != nil {
			log.Println("❌ Error al procesar el mensaje:", err)
			continue
		}

		fmt.Printf("Transacción recibida y procesada: %+v\n", transaction)

		// Llamada HTTP a otro servidor para enviar el idPedido
		err := r.sendPedidoToExternalAPI(transaction.Idpedido)
		if err != nil {
			log.Println("❌ Error al enviar pedido a la API:", err)
			continue
		}

		// Confirmar la transacción después de que se haya enviado a la API
		r.ConfirmTransaction(transaction.Idpedido, "success")
	}
	return nil
}

// Enviar confirmación a la cola de respuestas
func (r *RabbitMQRepository) ConfirmTransaction(transactionID int, status string) {
	conn, err := amqp.Dial("amqp://guest:guest@3.218.163.41:80/")
	if err != nil {
		log.Println("Error conectando a RabbitMQ:", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Error abriendo canal:", err)
		return
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare("transactions_responses", true, false, false, false, nil)
	if err != nil {
		log.Println("Error declarando la cola de respuestas:", err)
		return
	}

	// Crear el mensaje de respuesta
	response := map[string]interface{}{
		"id":     transactionID,
		"status": status,
	}

	body, err := json.Marshal(response)
	if err != nil {
		log.Println("Error serializando respuesta:", err)
		return
	}

	// Publicar la respuesta en la cola
	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Println("Error enviando confirmación:", err)
		return
	}

	fmt.Println("Confirmación enviada:", transactionID)
}

// Realiza la solicitud HTTP al servidor externo con el idPedido
func (r *RabbitMQRepository) sendPedidoToExternalAPI(idPedido int) error {
	url := fmt.Sprintf("http://localhost:8081/pedidos/%d", idPedido)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error al crear la solicitud HTTP: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error al enviar la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error al obtener respuesta del servidor externo: %v", resp.Status)
	}

	fmt.Printf("Solicitud HTTP exitosa para idPedido %d\n", idPedido)
	return nil
}
