package infraestructure

import (
	"api-main/core"
	"api-main/pedido/producto/domain"
)

type MySQLRepository struct {
	conn *core.Conn_MySQL
}

func NewMySQLRepository() *MySQLRepository {
	conn := core.GetDBPool()
	return &MySQLRepository{conn: conn}
}


func (r *MySQLRepository)Save(producto *domain.Producto)error{
	query:=("INSERT INTO productos(nombre,precio,descripcion,tipo)VALUES(?,?,?,?)")
	_,err:=r.conn.DB.Exec(query,producto.Nombre,producto.Precio,producto.Descripcion,producto.Tipo)
	if err!=nil{
		return err
	}
	return err
}

func (r *MySQLRepository)Update(id int,producto *domain.Producto)error{
	query:=("UPDATE productos SET nombre=?,precio=?,descripcion=?,tipo=? where idProducto=?")
	_,err:=r.conn.DB.Exec(query,producto.Nombre,producto.Precio,producto.Descripcion,producto.Tipo,id)
	if err!=nil{
		return err
	}

	return err
}


func (r *MySQLRepository)Delete(id int)error{
	query:=("DELETE FROM productos WHERE idProducto=?")
	_,err:=r.conn.DB.Exec(query,id)
	if err!=nil{
		return nil
	}

	return nil
}

func (r *MySQLRepository) GetAll() ([]domain.Producto, error) {
	query := "SELECT idProducto,nombre,precio,descripcion,tipo FROM productos"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pedidos []domain.Producto
	for rows.Next() {
		var pedido domain.Producto
		if err := rows.Scan(&pedido.IdProducto,&pedido.Nombre,&pedido.Precio,&pedido.Descripcion,&pedido.Tipo); err != nil {
			return nil, err
		}
		pedidos = append(pedidos, pedido)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pedidos, nil
}


func (r*MySQLRepository)GetById(id int)([]domain.Producto,error){
	query:=("SELECT idProducto,nombre,precio, descripcion, tipo FROM productos WHERE idProducto=?")
	rows,err:=r.conn.DB.Query(query,id)

	if err!=nil{
		return nil,err
	}

	defer rows.Close()


	var pedidos []domain.Producto
	for rows.Next() {
		var pedido domain.Producto
		if err := rows.Scan(&pedido.IdProducto,&pedido.Nombre,&pedido.Precio,&pedido.Descripcion,&pedido.Tipo); err != nil {
			return nil, err
		}
		pedidos = append(pedidos, pedido)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pedidos, nil

}