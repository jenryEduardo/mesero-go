package application

import "last-api/robot/domain"

type SaveRobot struct {
	repo domain.IRobot
}

func NewSaveRobot(repo domain.IRobot) *SaveRobot {
	return &SaveRobot{repo: repo}
}

func (c *SaveRobot) Run(robot domain.Robot) error {
	return c.repo.Save(robot)
}