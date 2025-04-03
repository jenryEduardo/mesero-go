package domain

type IHistorial interface {
	FindColor(idPedido int) (int,error)
	GetStatus() error
	UpdateStatus(idPedido int, nuevoStatus string) error
}