package model

type PayAccount struct {
	Id          int    `xorm:"not null pk autoincr unique INT(11)"`
	Appid       string `xorm:"default '' VARCHAR(30)"`
	MchId       string `xorm:"default '' VARCHAR(30)"`
	PrivateKey  string `xorm:"default '' VARCHAR(100)"`
	Secret      string `xorm:"default '' VARCHAR(100)"`
	OrderPrefix string `xorm:"default '' VARCHAR(100)"`
	ProjectId   string `xorm:"default '' VARCHAR(30)"`
	CompanyId   string `xorm:"default '' VARCHAR(30)"`
	CompanyName string `xorm:"default '' VARCHAR(30)"`
	TemplateId  string `xorm:"default '' VARCHAR(100)"`
	JumpPage    string `xorm:"default '' VARCHAR(50)"`
	Mark        string `xorm:"default '' VARCHAR(30)"`
}
