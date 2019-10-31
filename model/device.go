package model

type Device struct {
	Id         int    `xorm:"not null pk autoincr unique INT(11)"`
	DeviceId   string `xorm:"not null default '' unique VARCHAR(20)"`
	Type       string `xorm:"default '' VARCHAR(30)"`
	State      string `xorm:"default '' VARCHAR(30)"`
	Mark       string `xorm:"default '' VARCHAR(200)"`
	CreateTime string `xorm:"default '2006-01-02 15:04:05' VARCHAR(30)"`
	UpdateTime string `xorm:"default '2006-01-02 15:04:05' VARCHAR(30)"`
}
