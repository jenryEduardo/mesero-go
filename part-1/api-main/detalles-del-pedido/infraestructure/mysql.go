package infraestructure

import (
	"database/sql"
	"api-main/detalles-del-pedido/domain"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) Save(detalles domain.DetallesPedido) error {
	query := "INSERT INTO detalles_pedido(idPedido, idProducto, cantidad, precio_unitario, subtotal) VALUES(?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, detalles.IdPedido, detalles.IdProducto, detalles.Cantidad, detalles.PrecioUnitario, detalles.Subtotal)
	if err != nil {
		return err
	}
	return err
}

func (r *MySQLRepository) GetAll() ([]domain.DetallesPedido, error) {
	query := "SELECT idDetallePedido, idPedido, idProducto, cantidad, precio_unitario, subtotal FROM detalles_pedido"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var detalles []domain.DetallesPedido
	for rows.Next() {
		var detallesP domain.DetallesPedido
		if err := rows.Scan(&detallesP.IdDetallePedido, &detallesP.IdPedido, &detallesP.IdProducto, &detallesP.Cantidad, &detallesP.PrecioUnitario, &detallesP.Subtotal); err != nil {
			return nil, err
		}
		detalles = append(detalles, detallesP)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return detalles, nil
}

func (r *MySQLRepository) GetById(idDetallePedido int) ([]domain.DetallesPedido, error) {
	query := "SELECT idDetallePedido, idPedido, idProducto, cantidad, precio_unitario, subtotal FROM detalles_pedido WHERE idDetallePedido=?"
	rows, err := r.db.Query(query, idDetallePedido)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var detalles []domain.DetallesPedido
	for rows.Next() {
		var detallesP domain.DetallesPedido
		if err := rows.Scan(&detallesP.IdDetallePedido, &detallesP.IdPedido, &detallesP.IdProducto,  &detallesP.Cantidad, &detallesP.PrecioUnitario, &detallesP.Subtotal); err != nil {
			return nil, err
		}
		detalles = append(detalles, detallesP)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return detalles, nil
}

func (r *MySQLRepository) Delete(idDetallePedido int) error {
	query := "DELETE FROM detalles_pedido WHERE idDetallePedido=?"
	_, err := r.db.Exec(query, idDetallePedido)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLRepository) Update(idDetallePedido int, detallesP domain.DetallesPedido) error {
	query := "UPDATE detalles_pedido SET idPedido=?, idProducto=?, cantidad=?, precio_unitario=?, subtotal=? WHERE idDetallePedido=?"
	_, err := r.db.Exec(query, detallesP.IdPedido, detallesP.IdProducto, detallesP.Cantidad, detallesP.PrecioUnitario, detallesP.Subtotal, idDetallePedido)
	if err != nil {
		return err
	}
	return err
}