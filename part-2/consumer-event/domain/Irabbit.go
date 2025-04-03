package domain


type Irabbitmq interface{
	ConsumeTransaction()error
	SendTransactionToController(trasaccion *Transaction)error
}