package model

type ChangeFace struct {
	Id           int     `xorm:"not null pk autoincr unique INT(11)"`
	ChangeFaceId string  `xorm:"not null default '' unique VARCHAR(20)"`
	Openid       string  `xorm:"default '' VARCHAR(30)"`
	MovieId      string  `xorm:"default '' VARCHAR(30)"`
	ActorIds     string  `xorm:"default '' VARCHAR(200)"`
	Rate         float64 `xorm:"default '' VARCHAR(7)"`
	FacePath     string  `xorm:"default '' VARCHAR(200)"`
	CreateTime   string  `xorm:"default '2006-01-02 15:04:05' VARCHAR(30)"`
	UpdateTime   string  `xorm:"default '2006-01-02 15:04:05' VARCHAR(30)"`
}
