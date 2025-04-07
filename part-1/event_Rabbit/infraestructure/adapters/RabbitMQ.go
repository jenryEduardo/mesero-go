package adapters

import (
	"encoding/json"
	"fmt"

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

func (r *RabbitMQRepository) PublishTransaction(id int) (bool, error) {
	fmt.Println("Publicando transacción:", id)

	ch, err := r.conn.Channel()
	if err != nil {
		return false, fmt.Errorf("Error abriendo canal: %v", err)
	}
	defer ch.Close()

	_, err = ch.QueueDeclare("transactions", true, false, false, false, nil)
	if err != nil {
		return false, fmt.Errorf("Error declarando la cola de transacciones: %v", err)
	}

	type Transaction struct {
		Idpedido int `json:"idpedido"`
	}
	body, _ := json.Marshal(Transaction{Idpedido: int(id)})
	if err != nil {
		return false, fmt.Errorf("Error serializando JSON: %v", err)
	}

	err = ch.Publish(
		"",
		"transactions",
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

	fmt.Println("Mensaje publicado correctamente en la cola 'transactions'")
	return true, nil
}
