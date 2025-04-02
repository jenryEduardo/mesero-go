package application

import "consumer-event/domain"

type ConsumeRabbitMQ struct {
	repo domain.IRabbitMQ
}

func NewConsumeRabbit(repo domain.IRabbitMQ)*ConsumeRabbitMQ {
	return &ConsumeRabbitMQ{repo:repo}
}

func (c *ConsumeRabbitMQ) Execute()error{
	return c.repo.ConsumeTransaction()
}