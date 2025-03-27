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





func (r *MySQLRepository)Save(pedido *domain.Pedido)error{
	query:=("INSERT INTO pedido(idMesa,nombre_cliente,status,total)(?,?,?,?)")
	_,err:=r.conn.DB.Exec(query,pedido.IdMesa,pedido.Nombre_cliente,pedido.Status,pedido.Total)
	if err!=nil{
		return err
	}
	return err
}

func (r *MySQLRepository)Update(id int,pedido *domain.Pedido)error{
	query:=("UPDATE pedido SET idMesa=?,nombre_cliente=?,status=?,total=? where idPedido=?")
	_,err:=r.conn.DB.Exec(query,id)
	if err!=nil{
		return err
	}

	return err
}


func (r *MySQLRepository)Delete(id int)error{
	query:=("DELETE FROM pedido WHERE idPedido=?")
	_,err:=r.conn.DB.Exec(query,id)
	if err!=nil{
		return nil
	}

	return nil
}

func (r *MySQLRepository) GetAll() ([]domain.Pedido, error) {
	query := "SELECT idPedido, idMesa, nombre_cliente, status, total FROM pedido"
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


func (r*MySQLRepository)GetById(id int)([]domain.Pedido,error){
	query:=("SELECT idPedido, idMesa, nombre_cliente, status, total FROM pedido WHERE idPedido=?")
	rows,err:=r.conn.DB.Query(query,id)

	if err!=nil{
		return nil,err
	}

	defer rows.Close()


	var pedidos []domain.Pedido

	for rows.Next(){
		var pedido domain.Pedido

		if err:=rows.Scan(pedido.IdPedido,pedido.IdMesa,pedido.Nombre_cliente,pedido.Status,pedido.Total);err!=nil{
			return nil,err
		}

		pedidos = append(pedidos, pedido)

	}

	if err = rows.Err();err!=nil{
		return nil,err
	}

	return pedidos,err

}