package infraestructure

import (
	"api-main/core"
	"api-main/pedido/domain"
)

type MySQLRepository struct {
	conn *core.Conn_MySQL
}

func NewMySQLRepository() *MySQLRepository {
	conn := core.GetDBPool()
	return &MySQLRepository{conn: conn}
}

func (r *MySQLRepository) Save(pedido *domain.Pedido) (int64, error) {
	// Insertar el pedido en la tabla "pedido"
	//recuerda que el front debe mandar siempre status pendiente
	result, err := r.conn.DB.Exec("INSERT INTO pedido (idMesa, nombre_cliente, status, total) VALUES (?, ?, ?, ?)",
		&pedido.IdMesa, &pedido.Nombre_cliente, &pedido.Status, &pedido.Total)
	if err != nil {
		return 0, err
	}

	// Obtener el ID del pedido recién insertadox1
	pedidoID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Insertar los productos en la tabla "detalles_pedido"
	for _, detalle := range pedido.Detalles {
		_, err := r.conn.DB.Exec("INSERT INTO detalles_pedido (idPedido, idProducto, cantidad, subtotal) VALUES (?, ?, ?, ?)",
			pedidoID, detalle.IdProducto, detalle.Cantidad, detalle.Subtotal)
		if err != nil {
			return 0, err
		}
	}

	// Actualizar el total del pedido
	_, err = r.conn.DB.Exec("UPDATE pedido SET total = (SELECT SUM(subtotal) FROM detalles_pedido WHERE idPedido = ?) WHERE idPedido = ?", pedidoID, pedidoID)
	if err != nil {
		return 0, err
	}

	return pedidoID, err
}

func (r *MySQLRepository) AgregarNuevoProducto(idPedido int, detalle *domain.DetallesPedido) error {
	_, err := r.conn.DB.Exec("INSERT INTO detalles_pedido (idPedido, idProducto,nombre_producto, cantidad, precio_unitario, subtotal) VALUES (?,?, ?, ?, ?, ?)",
		idPedido, &detalle.IdProducto, &detalle.NombreProducto, &detalle.Cantidad, &detalle.PrecioUnitario, &detalle.Subtotal)

	if err != nil {
		return err
	}

	// Actualizar el total del pedido
	r.conn.DB.Exec("UPDATE pedido SET total = (SELECT SUM(subtotal) FROM detalles_pedido WHERE idPedido = ?) WHERE idPedido = ?", idPedido, idPedido)

	return err
}

func (r *MySQLRepository) ObtenerTotalPedido(idPedido int) (float64, error) {
	var total float64
	err := r.conn.DB.QueryRow("SELECT SUM(subtotal) FROM detalles_pedido WHERE id_pedido = ?", idPedido).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *MySQLRepository) Update(id int, pedido *domain.Pedido) error {
	query := ("UPDATE pedido SET idMesa=?,nombre_cliente=?,status=?,total=? where idPedido=?")
	_, err := r.conn.DB.Exec(query, &pedido.IdMesa, &pedido.Nombre_cliente, &pedido.Status, &pedido.Total, id)
	if err != nil {
		return err
	}

	return err
}

func (r *MySQLRepository) Delete(id int) error {
	query := ("DELETE FROM pedido WHERE idPedido=?")
	_, err := r.conn.DB.Exec(query, id)
	if err != nil {
		return nil
	}

	return nil
}

func (r *MySQLRepository) GetAll() ([]domain.Pedido, error) {
	query := "SELECT idPedido ,idMesa, nombre_cliente, status, total FROM pedido"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pedidos []domain.Pedido
	for rows.Next() {
		var pedido domain.Pedido
		if err := rows.Scan(&pedido.IdPedido, &pedido.IdMesa, &pedido.Nombre_cliente, &pedido.Status, &pedido.Total); err != nil {
			return nil, err
		}
		pedidos = append(pedidos, pedido)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pedidos, nil
}

func (r *MySQLRepository) GetById(id int) ([]domain.Pedido, error) {
	query := ("SELECT idPedido, idMesa, nombre_cliente, status, total FROM pedido WHERE idPedido=?")
	rows, err := r.conn.DB.Query(query, id)


	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var pedidos []domain.Pedido

	for rows.Next() {
		var pedido domain.Pedido

		if err := rows.Scan(&pedido.IdPedido, &pedido.IdMesa, &pedido.Nombre_cliente, &pedido.Status, &pedido.Total); err != nil {
			return nil, err
		}

		pedidos = append(pedidos, pedido)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pedidos, err

}
