package model

type WxBindActivity struct {
	Id         int    `xorm:"not null pk autoincr unique INT(11)"`
	WxOpenid   string `xorm:"default '' VARCHAR(50)"`
	ActivityId string `xorm:"default '' VARCHAR(50)"`
	ProjectId  string `xorm:"default '' VARCHAR(50)"`
	Phone      string `xorm:"default '' VARCHAR(20)"`
	CreateTime string `xorm:"default '1970-01-01 08:00:00' VARCHAR(45)"`
	UpdateTime string `xorm:"default '1970-01-01 08:00:00' VARCHAR(45)"`
}
