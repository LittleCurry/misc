package model

type User struct {
	Id            int    `xorm:"not null pk autoincr unique INT(11)"`
	UserId        string `xorm:"not null default '' unique VARCHAR(20)"`
	Phone         string `xorm:"default '' VARCHAR(50)"`
	Passwd        string `xorm:"default '' VARCHAR(50)"`
	UserName      string `xorm:"not null default '' unique VARCHAR(20)"`
	Gender        int    `xorm:"default '' INT(1)"`
	Birthday      string `xorm:"default '' VARCHAR(50)"`
	Address       string `xorm:"default '' VARCHAR(30)"`
	WxOpenid      string `xorm:"default '' VARCHAR(50)"`
	WxUnionid     string `xorm:"default '' VARCHAR(50)"`
	Head          string `xorm:"default '' VARCHAR(30)"`
	Introducation string `xorm:"default '' VARCHAR(50)"`
	Del           int    `xorm:"default 0 TINYINT(1)"`
	CreateTime    string `xorm:"default '2006-01-02 15:04:05' VARCHAR(30)"`
}
