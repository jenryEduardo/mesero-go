package domain


type Imesa interface{
	Save(mesa *Mesa)error
	Update(idMesa int,mesa Mesa)error
	Delete(id int)error
	GetById(id int)([]Mesa,error)
	GetAll()([]Mesa,error)
}