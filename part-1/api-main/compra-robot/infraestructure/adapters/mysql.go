package adapters

import (
	"api-main/core"
	"api-main/compra-robot/domain"
)

type MySQLRepository struct {
	conn *core.Conn_MySQL
}

func NewMySQLRepository() *MySQLRepository {
	conn := core.GetDBPool()
	return &MySQLRepository{conn: conn}
}

func (r *MySQLRepository) Save(compra *domain.Compra) error {
	query := "INSERT INTO compra (idRobot, idUsuario) VALUES (?, ?)"
	_, err := r.conn.DB.Exec(query, compra.IdRobot, compra.IdUsuario)
	return err
}

func (r *MySQLRepository) Update(id int, compra domain.Compra) error {
	query := "UPDATE compra SET idRobot=?, idUsuario=? WHERE idCompra=?"
	_, err := r.conn.DB.Exec(query, compra.IdRobot, compra.IdUsuario, id)
	return err
}

func (r *MySQLRepository) Delete(id int) error {
	query := "DELETE FROM compra WHERE idCompra=?"
	_, err := r.conn.DB.Exec(query, id)
	return err
}

func (r *MySQLRepository) GetAll() ([]domain.Compra, error) {
	query := "SELECT idCompra, idRobot, idUsuario FROM compra"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var compras []domain.Compra
	for rows.Next() {
		var compra domain.Compra
		if err := rows.Scan(&compra.IdCompra, &compra.IdRobot, &compra.IdUsuario); err != nil {
			return nil, err
		}
		compras = append(compras, compra)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return compras, nil
}

func (r *MySQLRepository) GetById(id int) (*domain.Compra, error) {
	query := "SELECT idCompra, idRobot, idUsuario FROM compra WHERE idCompra=?"
	row := r.conn.DB.QueryRow(query, id)

	var compra domain.Compra
	if err := row.Scan(&compra.IdCompra, &compra.IdRobot, &compra.IdUsuario); err != nil {
		return nil, err
	}

	return &compra, nil
}
