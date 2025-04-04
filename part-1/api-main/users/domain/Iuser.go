package domain

type Iuser interface{
	Save(user *User)error
	GetUserById(id int)([]User,error)
	GetAllUser()([]User,error)
	UpdateUser(id int,user *User)error
	DeleteUser(id int)error
	Login(email string,password string)(*User,string, error)
	GetUserByEmail(email string)(*User,error)	
}