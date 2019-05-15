package model

type Print struct {
	Id         int    `xorm:"not null pk autoincr unique INT(11)"`
	PrintId    string `xorm:"default '' VARCHAR(30)"`
	ActivityId string `xorm:"default '' VARCHAR(30)"`
	Address    string `xorm:"default '' VARCHAR(100)"`
	Lon        string `xorm:"default '' VARCHAR(30)"`
	Lat        string `xorm:"default '' VARCHAR(30)"`
	Mark       string `xorm:"default '' VARCHAR(200)"`
	CreateTime string `xorm:"default '1970-01-01 08:00:00' VARCHAR(45)"`
}
