package model

type GoodsCar struct {
	Id         int    `xorm:"not null pk autoincr unique INT(11)"`
	GoodsCarId string `xorm:"not null default '' unique VARCHAR(20)"`
	WxOpenid   string `xorm:"default '' VARCHAR(30)"`
	Taskid     string `xorm:"default '' VARCHAR(50)"`
	ActivityId string `xorm:"default '' VARCHAR(50)"`
	IsPay      string `xorm:"default '' VARCHAR(2)"`
	Total      int    `xorm:"default '' VARCHAR(30)"`
	CreateTime string `xorm:"default '2006-01-02 15:04:05' VARCHAR(45)"`
	UpdateTime string `xorm:"default '2006-01-02 15:04:05' VARCHAR(45)"`
}
