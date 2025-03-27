package domain

type IDetallePedido interface {
	Save(detalles DetallesPedido) error
	GetAll() ([]DetallesPedido, error)
	GetById(idDetalle int) ([]DetallesPedido, error)
	Delete(idDetalle int) error
	Update(idDetalle int, detalles DetallesPedido) error 
}