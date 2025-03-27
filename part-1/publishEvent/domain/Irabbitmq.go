package domain


type Irabbitmq interface{
	PublishTransaction(rabbit *RabbitMQ)(bool,error)
}