package domain

type IRabbitMQ interface {
	ConsumeTransaction(rabbit *RabbitMQ)(bool,error)
}