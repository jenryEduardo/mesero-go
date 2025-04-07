package application

import "consumer/domain"

type saveInRbbitmq struct {
	repo domain.Irabbitmq
}


func NewRabbitSave(repo domain.Irabbitmq)*saveInRbbitmq{
	return &saveInRbbitmq{repo: repo}
}


func(c *saveInRbbitmq)Execute(id int)(bool,error){
	return c.repo.PublishTransaction(id)
}