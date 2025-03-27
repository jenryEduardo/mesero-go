package dependencies

import (
	"last-api/core"
	"last-api/detalles-del-pedido/application"
	"last-api/detalles-del-pedido/infraestructure"
	"last-api/detalles-del-pedido/infraestructure/controllers"
)

var (
	mySQL infraestructure.MySQLRepository
)

func Init() {
	db, err := core.InitMySQL()
	if err != nil {
		return
	}
	mySQL = *infraestructure.NewMySQLRepository(db)
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