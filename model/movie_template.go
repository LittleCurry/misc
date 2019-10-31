package model

type MovieTemplate struct {
	Id           int    `xorm:"not null pk autoincr unique INT(11)"`
	MovieId      string `xorm:"default '' VARCHAR(20)"`
	TemplateName string `xorm:"default '' VARCHAR(50)"`
	Url          string `xorm:"default '' VARCHAR(30)"`
	CoverH       string `xorm:"default '' VARCHAR(100)"`
	CoverV       string `xorm:"default '' VARCHAR(100)"`
	CreateTime   string `xorm:"default '2006-01-02 15:04:05' VARCHAR(30)"`
}

type MovieAndActor struct {
	MovieTemplate `xorm:"extends"`
	Actor         `xorm:"extends"`
}
