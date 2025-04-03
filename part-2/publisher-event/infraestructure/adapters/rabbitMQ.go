package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"
	"publisher/domain"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn *amqp.Connection
}

func NewRabbitMQ() (*RabbitMQ, error) {
	conn, err := amqp.Dial("amqp://diedgo:rabbit666@18.232.202.247:5672/")
	if err != nil {
		return nil, fmt.Errorf("Error conectando a RabbitMQ: %v", err)
	}
	return &RabbitMQ{conn:conn}, nil
}

func(r *RabbitMQ) Publish(status *domain.EventRabbit)(bool,error) {
	fmt.Println("Actualizando Status: ", status)
	ch, err := r.conn.Channel()
	if err != nil {
		return false, fmt.Errorf("Error abriendo el canal: %v", err)
	}

	body, err := json.Marshal(status)
	if err != nil {
		return false, fmt.Errorf("Error serializando JSON: %v", err)
	}

	err = ch.Publish(
		"",
		"status",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return false, fmt.Errorf("Error publicando mensaje: %v", err)
	}

	// Esperar respuesta en `transactions_responses`
	responseChan := make(chan bool)

	go func() {
		respConn, _ := amqp.Dial("amqp://diedgo:rabbit666@18.232.202.247:5672/")
		defer respConn.Close()

		respCh, _ := respConn.Channel()
		defer respCh.Close()

		queue, _ := respCh.QueueDeclare(
			"status_responses",
			true,
			false,
			false,
			false,
			nil)
		msgs, _ := respCh.Consume(
			queue.Name,
			"",
			true,
			false,
			false,
			false,
			nil)

		for msg := range msgs {
			var response map[string]interface{}
			_ = json.Unmarshal(msg.Body, &response)

			fmt.Println("Mensaje recibido de status_responses:", response)

			if response["id"] == status.Status {
				status, ok := response["status"].(string)
				if !ok {
					fmt.Println("Error: no se pudo leer el estado de la transacción.")
					responseChan <- false
				} else if status == "success" {
					responseChan <- true
				} else {
					responseChan <- false
				}
				close(responseChan) 
				break
			}
		}
	}()

	select {
	case success := <-responseChan:
		if success {
			fmt.Println("Confirmación recibida. La transferencia fue exitosa.")

			// Hacer POST sin enviar JSON
			req, err := http.NewRequest("POST", "http://localhost:3003/publisherStatus/", nil)
			req.Header.Set("Content-Type", "application/json") 

			if err != nil {
				fmt.Println("Error creando la solicitud:", err)
				return false, nil
			}

			client := &http.Client{Timeout: 10 * time.Second} // Cliente con timeout
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error al llamar al consumer:", err)
				return false, nil
			}
			defer resp.Body.Close()

			fmt.Println("Respuesta del consumer:", resp.Status)
			return true, nil // Cambio a `true` si todo sale bien
		} else {
			fmt.Println("Error en el procesamiento de la transacción.")
			return false, nil
		}
	case <-time.After(10 * time.Second):
		fmt.Println("No se recibió respuesta en el tiempo esperado.")
		return false, fmt.Errorf("timeout esperando respuesta")
	}
	
}