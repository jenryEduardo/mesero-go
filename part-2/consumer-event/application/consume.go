package application

import "consumer2/domain"

type ConsumeRbbit struct {
	repo domain.Irabbitmq
}

func NewConsume(repo domain.Irabbitmq)*ConsumeRbbit{
	return &ConsumeRbbit{repo: repo}
}


func(c *ConsumeRbbit)Execute()error{
	return c.repo.ConsumeTransaction()
}