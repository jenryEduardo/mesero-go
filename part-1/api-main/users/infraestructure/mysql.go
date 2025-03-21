package infraestructure

import (
	"api-main/core"
	"api-main/users/domain"
)

type MySQLRepository struct {
	conn *core.Conn_MySQL
}


func NewMySQLRepository() *MySQLRepository {
	conn := core. GetDBPool()
	return &MySQLRepository{conn: conn}
}

func (r *MySQLRepository) Save(p *domain.User) error {
	query := "INSERT INTO usuarios (name,last_name,email,password) VALUES (?,?,?,?)"
	_, err := r.conn.DB.Exec(query,p.Name,p.Last_name,p.Email,p.Password)
	return err
}

func (r *MySQLRepository) GetUserById(id int) ([]domain.User, error) {
	query := "SELECT idUsuario, name, last_name, email, password FROM usuarios WHERE idUsuario = ?"
	rows, err := r.conn.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close() 

	var users []domain.User

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.Last_name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}


func (r *MySQLRepository) GetAllUser()([]domain.User,error) {
	query := "SELECT idUsuario, name, last_name, email, password FROM usuarios"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() 

	var users []domain.User

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.Last_name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}




func (r *MySQLRepository) UpdateUser(id int,p *domain.User) error {
	query := "UPDATE usuarios SET name=?,last_name=?,email=?,password=? where idUsuario=?"
	_, err := r.conn.DB.Exec(query,p.Name,p.Last_name,p.Email,p.Password,id)
	return err
}

func (r *MySQLRepository) DeleteUser(id int) error {
	query := "DELETE FROM usuarios WHERE id=?"
	_, err := r.conn.DB.Exec(query,id)
	if err!=nil{
		return err
	}
	return err
}

