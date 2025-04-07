package domain


type Irabbitmq interface{
	PublishTransaction(id int)(bool,error)
}