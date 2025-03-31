package adapters

import (
	"consumer-event/domain"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMQRepository struct {
	conn *amqp.Connection
}

// Constructor para RabbitMQRepository
func NewRabbitMQRepository() (*RabbitMQRepository, error) {
	conn, err := amqp.Dial("amqp://diedgo:rabbit666@18.232.202.247:5672/")
	if err != nil {
		return nil, fmt.Errorf("Error conectando a RabbitMQ: %v", err)
	}

	return &RabbitMQRepository{conn: conn}, nil
}

func (r *RabbitMQRepository) ConsumeTransactions() error {
	ch, err := r.conn.Channel()
	if err != nil {
		return fmt.Errorf("Error abriendo canal: %v", err)
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare("pedido", true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("Error declarando la cola de pedidos: %v", err)
	}

	msgs, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
	if err != nil {
		return fmt.Errorf("Error al consumir mensajes: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			var pedido domain.RabbitMQ
			if err := json.Unmarshal(msg.Body, &pedido); err != nil {
				fmt.Println("Error al deserializar mensaje:", err)
				continue
			}

			fmt.Println("Pedido recibido:", pedido)

			// Simular procesamiento
			time.Sleep(2 * time.Second)

			// Enviar confirmación
			respCh, _ := r.conn.Channel()
			defer respCh.Close()

			responseQueue, _ := respCh.QueueDeclare("pedido_responses", true, false, false, false, nil)
			response := map[string]interface{}{
				"id":     pedido.IdPedido,
				"status": "success",
			}
			responseBody, _ := json.Marshal(response)

			err = respCh.Publish("", responseQueue.Name, false, false, amqp.Publishing{
				ContentType: "application/json",
				Body:        responseBody,
			})
			if err != nil {
				fmt.Println("Error al enviar respuesta:", err)
			}

			// Hacer POST sin enviar JSON
			req, err := http.NewRequest("POST", "http://localhost:3000/consume/", nil)
			req.Header.Set("Content-Type", "application/json")

			if err != nil {
				fmt.Println("Error creando la solicitud HTTP:", err)
				continue
			}

			client := &http.Client{Timeout: 10 * time.Second}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error al llamar al consumer API:", err)
				continue
			}
			defer resp.Body.Close()

			fmt.Println("Respuesta del consumer API:", resp.Status)
		}
	}()

	fmt.Println("[*] Esperando mensajes. Para salir, presiona CTRL+C")
	<-forever
	return nil
}
