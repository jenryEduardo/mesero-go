package application

import "publisher/domain"

type PublishInRabbit struct {
	repo domain.IEventRabbit
}

func NewPublishInRabbit(repo domain.IEventRabbit) *PublishInRabbit {
	return &PublishInRabbit{repo: repo}
}

func (c *PublishInRabbit) Run(data *domain.EventRabbit) (bool, error) {
	return c.repo.Publish(data)
}