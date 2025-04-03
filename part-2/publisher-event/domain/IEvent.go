package domain

type IEventRabbit interface {
	Publish(rabbit *EventRabbit)(bool, error)
}