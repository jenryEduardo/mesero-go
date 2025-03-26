package domain


type Iproducto interface{
	Save(producto *Producto)error
	Update(id int,producto *Producto)error
	Delete(id int)error
	GetAll()([]Producto,error)
	GetById(id int)([]Producto,error)
}