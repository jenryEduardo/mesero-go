package adapters

import (
	"bytes"
	"consumer2/domain"
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
	conn, err := amqp.Dial("amqp://guest:guest@13.217.71.115:5672/")
	if err != nil {
		return nil, fmt.Errorf("Error conectando a RabbitMQ: %v", err)
	}

	return &RabbitMQRepository{conn: conn}, nil
}

func (r *RabbitMQRepository) ConsumeTransaction() error {
	ch, err := r.conn.Channel()
	if err != nil {
		log.Fatal("Error al abrir el canal: ", err)
	}
	defer ch.Close()

	// Declarar la cola de transacciones
	queue, err := ch.QueueDeclare("transactions", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Error declarando la cola: ", err)
	}

	msgs, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Error al consumir la cola: ", err)
	}

	fmt.Println(" Escuchando transacciones en RabbitMQ...")

	for msg := range msgs {
		var transaction domain.Transaction
		if err := json.Unmarshal(msg.Body, &transaction); err != nil {
			log.Println(" Error al procesar el mensaje: ", err)
			continue
		}

		fmt.Printf("Transacción recibida: %+v\n", transaction)

		// Enviar los datos al controlador para guardarlos en la BD
		err := r.SendTransactionToController(&transaction)
		if err != nil {
			log.Println(" Error al enviar la transacción al controlador: ", err)
		}
	}
	return nil
}

func (r *RabbitMQRepository) SendTransactionToController(transaction *domain.Transaction) error {
	url := "http://localhost:3001/transaccion/"

	body, err := json.Marshal(transaction)
	if err != nil {
		log.Println("Error serializando la transacción: ", err)
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println(" Error enviando la solicitud HTTP: ", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(" Error en la respuesta del servidor: ", resp.Status)
		return fmt.Errorf("error en la respuesta del servidor: %s", resp.Status)
	}

	fmt.Println("Transacción guardada correctamente en la base de datos")
	return nil
}
