package domain

type Iuser interface{
	Save(user *User)error
	GetUserById(id int)([]User,error)
	GetAllUser()([]User,error)
	UpdateUser(id int,user *User)error
	DeleteUser(id int)error
}