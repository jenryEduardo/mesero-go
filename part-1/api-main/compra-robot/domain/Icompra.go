package domain


type ICompra interface{
	Save(compra *Compra)error
	Update(idCompra int,compra *Compra)error
	GetAll()([]Compra,error)
	GetById(id int)([]Compra,error)
	Delete(id int)error
}