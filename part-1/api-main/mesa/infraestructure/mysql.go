package infraestructure

import (
	"api-main/core"
	"api-main/mesa/domain"
)

type MySQLRepository struct {
	conn *core.Conn_MySQL
}

func NewMySQLRepository() *MySQLRepository {
	conn := core.GetDBPool()
	return &MySQLRepository{conn: conn}
}


func (r *MySQLRepository)Save(mesa *domain.Mesa)error{

	query := "INSERT INTO mesa(status) VALUES (?)"
	_,err:=r.conn.DB.Exec(query,mesa.Status)

	if err!=nil{
		return err
	}
	return err
}

func (r *MySQLRepository)Update(id int,mesa domain.Mesa)error{

	query:=("UPDATE mesa SET status=? WHERE idMesa=?")

	_,err:=r.conn.DB.Exec(query,mesa.Status,id)
return err
}

func (r *MySQLRepository)Delete(id int)error{
	query:=("DELETE FROM mesa WHERE idMesa=?")
	_,err:=r.conn.DB.Exec(query,id)
	return err
}


func (r *MySQLRepository) GetAll() ([]domain.Mesa, error) {
	query := "SELECT idMesa, status FROM mesa"
	rows, err := r.conn.DB.Query(query) // Usar Query en lugar de Exec
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mesas []domain.Mesa
	for rows.Next() {
		var mesa domain.Mesa
		if err := rows.Scan(&mesa.IdMesa, &mesa.Status); err != nil {
			return nil, err
		}
		mesas = append(mesas, mesa)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return mesas, nil
}

func (r *MySQLRepository) GetById(id int) ([]domain.Mesa, error) {
	query := "SELECT idMesa, status FROM mesa WHERE idMesa=?"
	rows, err := r.conn.DB.Query(query,id) // Usar Query en lugar de Exec
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mesas []domain.Mesa
	for rows.Next() {
		var mesa domain.Mesa
		if err := rows.Scan(&mesa.IdMesa, &mesa.Status); err != nil {
			return nil, err
		}
		mesas = append(mesas, mesa)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return mesas, nil
}
