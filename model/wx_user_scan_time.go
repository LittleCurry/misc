package model

type WxUserScanTime struct {
	Id         int    `xorm:"not null pk autoincr unique INT(11)"`
	ActivityId string `xorm:"default '' VARCHAR(50)"`
	WxOpenid   string `xorm:"default '' VARCHAR(50)"`
	WxUnionid  string `xorm:"default '' VARCHAR(50)"`
	CreateTime string `xorm:"default '1970-01-01 08:00:00' VARCHAR(45)"`
	UpdateTime string `xorm:"default '1970-01-01 08:00:00' VARCHAR(45)"`
}

/*
type WxUserScanTime struct {
	Id         int    `xorm:"not null pk autoincr INT(11)"`
	ActivityId string `xorm:"default '' VARCHAR(50)"`
	WxOpenid   string `xorm:"default '' VARCHAR(50)"`
	WxUnionid  string `xorm:"default '' VARCHAR(50)"`
	CreateTime string `xorm:"default '' VARCHAR(45)"`
	UpdateTime string `xorm:"default '' VARCHAR(45)"`
}
*/