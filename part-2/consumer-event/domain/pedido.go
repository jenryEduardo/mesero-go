package domain

type Pedido struct {
	IdPedido      int     `json:"id_pedido"`
	IdMesa        int     `json:"id_mesa"`
	NombreCliente string  `json:"nombre_cliente"`
	Status        string  `json:"status"`
	Total         float64 `json:"total"`
}

type Response struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}