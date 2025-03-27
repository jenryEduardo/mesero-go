package domain

type Historial struct {
	IdHistorial int 		`json:"id_historial"`
	IdRobot int 			`json:"idRobot"`
	Estatus_entrega string 	`json:"estatus_entrega"`
	Percanses string 		`json:"percanses"`
}