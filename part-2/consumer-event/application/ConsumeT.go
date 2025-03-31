package application

import "consumer-event/domain"

type ConsumeRabbitMQ struct {
	repo domain.IRabbitMQ
}

func NewConsumeRabbit(repo domain.IRabbitMQ)*ConsumeRabbitMQ {
	return &ConsumeRabbitMQ{repo:repo}
}

func (c *ConsumeRabbitMQ) Execute(data *domain.RabbitMQ)(bool,error){
	return c.repo.ConsumeTransaction(data)
}