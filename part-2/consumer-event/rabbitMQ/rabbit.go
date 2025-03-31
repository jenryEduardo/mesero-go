package rabbitmq

import (
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"consumer-event/domain"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQ() (*RabbitMQ, error) {
	_ = godotenv.Load()
	rabbitURL := os.Getenv("RABBIT_URL")

	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{conn: conn, channel: ch}, nil
}

func (r *RabbitMQ) SetupExchangeAndQueues() {
	err := r.channel.ExchangeDeclare(
		"pedido_exchange",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Error declarando el exchange")

	_, err = r.channel.QueueDeclare(
		"pedidoQueue",
		true,
		false,
		false,
	    false,
		nil,
		)
	failOnError(err, "Error declarando la cola de pedidos")

	err = r.channel.QueueBind(
		"pedidoQueue",
		"pedido_routing_key",
		"pedido_exchange",
	   	false,
		nil,
		)
	failOnError(err, "Error vinculando la cola de pedidos")
}

func (r *RabbitMQ) PublishPedido(pedido domain.Pedido) error {
	body, err := json.Marshal(pedido)
	if err != nil {
		return err
	}

	err = r.channel.Publish(
		"pedido_exchange",
		"pedido_routing_key",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	return err
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
