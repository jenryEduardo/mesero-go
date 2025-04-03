package domain


type Robot struct {
	IdRobot string `json:"idRobot" gorm:"primaryKey"`
	Alias   string `json:"alias"`
	//Quedaría solo idrobot, alias
}