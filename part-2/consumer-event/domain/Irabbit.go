package domain


type Irabbitmq interface{
	ConsumeTransaction()error
	ConfirmTransaction(transaccionId int,status string)
}