package infraestructure

import (
	"database/sql"
	"last/historial-de-entregas/domain"

)


type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db:db}
}

func (r *MySQLRepository) Save(historial domain.Historial) error {
	query:= "INSERT INTO historial_entrega(idRobot, estatus_entrega, percanses) VALUES(?,?,?)"
	_, err := r.db.Exec(query, historial.IdRobot, historial.Estatus_entrega, historial.Percanses)
	if err != nil {
		return err
	}
	return err
}

func (r *MySQLRepository) GetAll()([]domain.Historial, error) {
	query:= "SELECT id_historial, idRobot, estatus_entrega, percanses FROM historial_entrega"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	} 
	defer rows.Close()

	var historial []domain.Historial

	for rows.Next() {
		var h domain.Historial
		if err := rows.Scan(&h.IdHistorial, &h.IdRobot, &h.Estatus_entrega, &h.Percanses); err != nil {
			return nil, err
		}
		historial = append(historial, h)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return historial, nil
}

func (r *MySQLRepository) GetById(idHistorial int) ([]domain.Historial, error) {
	query := "SELECT id_historial, idRobot, estatus_entrega, percanses FROM historial_entrega WHERE id_historial=?"

	rows, err := r.db.Query(query, idHistorial)
	if err != nil {
		return nil,err
	}

	defer rows.Close()

	var historial []domain.Historial

	for rows.Next() {
		var h domain.Historial

		if err := rows.Scan(&h.IdHistorial, &h.IdRobot, &h.Estatus_entrega, &h.Percanses); err != nil {
			return nil, err
		}
		historial = append(historial, h)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return historial, nil
}

func (r *MySQLRepository) Update(idHistorial int, historial domain.Historial) error {
	query := "UPDATE historial_entrega SET idRobot=?, estatus_entrega=?, percanses=? WHERE id_historial=?"
	_, err := r.db.Exec(query, historial.IdRobot, historial.Estatus_entrega, historial.Percanses, idHistorial)
	if err != nil {
		return err
	}
	return err
}

func (r *MySQLRepository) Delete(idHistorial int) error{
	query := "DELETE FROM historial_entrega WHERE id_historial=?"
	_, err := r.db.Exec(query, idHistorial)
	if err != nil {
		return err
	}
	return nil
}