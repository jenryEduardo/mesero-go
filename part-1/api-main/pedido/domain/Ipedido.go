package domain

type Ipedido interface{
	Save(pedido *Pedido)error
	Update(id int,pedido *Pedido)error
	Delete(id int)error
	GetAll()([]Pedido,error)
	GetById(id int)([]Pedido,error)
	
}