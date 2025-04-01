package dependencies

import (
	"api-main/core"
	"api-main/detalles-del-pedido/application"
	"api-main/detalles-del-pedido/infraestructure"
	"api-main/detalles-del-pedido/infraestructure/controllers"
)

var (
	mySQL infraestructure.MySQLRepository
)

func Init() {
	dbConn := core.GetDBPool() // GetDBPool solo devuelve un valor
	if dbConn == nil {
		return
	}
	mySQL = *infraestructure.NewMySQLRepository(dbConn.DB) // Usar dbConn.DB
}


func SaveDetalle() *controllers.SaveDetallesCtrl {
	ucSaveD := application.NewSaveDetalles(&mySQL)
	return controllers.NewSaveDetallesCtrl(ucSaveD)
}

func GetAllDetalles() *controllers.GetAllDetallesCtrl {
	ucGetAll := application.NewGetAllDetalles(&mySQL)
	return controllers.NewGetAllDetallesCtrl(ucGetAll)
}

func GetByIdDetalles() *controllers.GetDetallesByIDCtrl {
	ucGetByID := application.NewGetDetalleByID(&mySQL)
	return controllers.NewGetRsIDCtrl(ucGetByID)
}

func DeleteDetalles() *controllers.DeleteDetallesCtrl {
	ucDelete := application.NewDeleteDetalles(&mySQL)
	return controllers.NewDeleteRSCtrl(ucDelete)
}

func UpdateDetalles() *controllers.UpdateDetallesCtrl {
	ucUpdate := application.NewUpdateDetalles(&mySQL)
	return controllers.NewUpdateDetallesCtrl(ucUpdate) 
}