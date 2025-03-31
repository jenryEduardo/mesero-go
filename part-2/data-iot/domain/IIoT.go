package domain

type InterfaceIoT interface {
	Enviar(iot IoT) error
	Recibir()([]IoT,error)
}