	package adapters

	import (
		"consumer/domain"
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
		conn, err := amqp.Dial("amqp://guest:guest@13.217.71.115:5672/")
		if err != nil {
			return nil, fmt.Errorf("Error conectando a RabbitMQ: %v", err)
		}

		return &RabbitMQRepository{conn: conn}, nil
	}

	func (r *RabbitMQRepository) PublishTransaction(pedido *domain.RabbitMQ) (bool, error) {
		fmt.Println("Publicando pedido:", pedido)

		ch, err := r.conn.Channel()
		if err != nil {
			return false, fmt.Errorf("Error abriendo canal: %v", err)
		}
		defer ch.Close()

		_, err = ch.QueueDeclare("pedido", true, false, false, false, nil)
		if err != nil {
			return false, fmt.Errorf("Error declarando la cola de transacciones: %v", err)
		}

		body, err := json.Marshal(pedido)
		if err != nil {
			return false, fmt.Errorf("Error serializando JSON: %v", err)
		}

		err = ch.Publish(
			"",
			"pedido",
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
			respConn, _ := amqp.Dial("amqp://guest:guest@13.217.71.115:5672/")
			defer respConn.Close()

			respCh, _ := respConn.Channel()
			defer respCh.Close()

			queue, _ := respCh.QueueDeclare("pedido_responses", true, false, false, false, nil)
			msgs, _ := respCh.Consume(queue.Name, "", true, false, false, false, nil)

			for msg := range msgs {
				var response map[string]interface{}
				_ = json.Unmarshal(msg.Body, &response)

				fmt.Println("Mensaje recibido de transactions_responses:", response)

				if response["id"] == pedido.ID {
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
				req, err := http.NewRequest("POST", "http://localhost:3002/consumer/", nil)
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
			return false, fmt.Errorf("Timeout esperando respuesta")
		}
		
	}
