package controllers

import (
	"second/robot/application"
	"second/robot/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SaveRobotCtrl struct {
	uc *application.SaveRobot
}

func NewSaveRobotCtrl(uc *application.SaveRobot) *SaveRobotCtrl {
	return &SaveRobotCtrl{uc: uc}
}

func (ctrl *SaveRobotCtrl) Run(c *gin.Context) {
	var robot domain.Robot

	if err := c.ShouldBindJSON(&robot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.uc.Run(robot)

	if err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": true,
			"data": gin.H{
				"type": "robot",
				"idRobot": robot.IdRobot,
				"attributes": gin.H{
					"alias": robot.Alias,
				},
			},
		})
	}
}