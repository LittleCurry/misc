package model

type Actor struct {
	Id         int    `xorm:"not null pk autoincr unique INT(11)"`
	ActorId    string `xorm:"not null default '' VARCHAR(20)"`
	MovieId    string `xorm:"default '' VARCHAR(20)"`
	ActorName  string `xorm:"default '' VARCHAR(30)"`
	Gender     int    `xorm:"default '' INT(1)"`
	Avatar     string `xorm:"default '' VARCHAR(100)"`
	TrueName   string `xorm:"default '' VARCHAR(20)"`
	CreateTime string `xorm:"default '2006-01-02 15:04:05' VARCHAR(30)"`
}
