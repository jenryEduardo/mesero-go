package domain


type Robot struct {
	IdRobot                int    `json:"idRobot"`
	IdPedido               int    `json:"idPedido"`
	IdCircuito             int    `json:"idCircuito"`
	Alias                  string `json:"alias"`
	Mantenimiento          bool   `json:"Mantenimiento"`
	Fecha_de_mantenimiento string `json:"fecha_De_Mantenimiento"`
}