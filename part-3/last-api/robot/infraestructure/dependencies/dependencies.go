package dependencies

import (
	"last-api/core"
	"last-api/robot/application"
	"last-api/robot/infraestructure"
	"last-api/robot/infraestructure/controllers"
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

func SaveRobot() *controllers.SaveRobotCtrl {
	ucSaveRobot := application.NewSaveRobot(&mySQL)
	return controllers.NewSaveRobotCtrl(ucSaveRobot)
}

func GetAllRobots() *controllers.GetAllRobotsCtrl {
	ucGetAllRobots := application.NewGetAllRobot(&mySQL)
	return controllers.NewGetAllRobotsCtrl(ucGetAllRobots)
}

func GetRobotByID() *controllers.GetRobotIDCtrl {
	ucGetRobotByID := application.NewGetByIdRobot(&mySQL)
	return controllers.NewGetRobotByIDCtrl(ucGetRobotByID)
}

func UpdateRobot() *controllers.UpdateRobotCtrl {
	ucUpdateRobot := application.NewUpdateRobot(&mySQL)
	return controllers.NewUpdateRobotCtrl(ucUpdateRobot)
}

func DeleteRobot() *controllers.DeleteRobotCtrl {
	ucDeleteRobot := application.NewDeleteRobot(&mySQL)
	return controllers.NewDeleteRobotCtrl(ucDeleteRobot)
}

