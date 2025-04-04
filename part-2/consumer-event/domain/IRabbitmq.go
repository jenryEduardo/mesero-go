package domain

type IRabbitMQ interface {
	ConsumeTransaction()error
}