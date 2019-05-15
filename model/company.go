package model

type Company struct {
	Id          int    `xorm:"not null pk autoincr unique INT(11)"`
	CompanyId   string `xorm:"not null default '' unique VARCHAR(20)"`
	CompanyName string `xorm:"default '' VARCHAR(30)"`
	Mark        string `xorm:"default '' VARCHAR(50)"`
	Leader      string `xorm:"default '' VARCHAR(50)"`
	Phone       string `xorm:"default '' VARCHAR(11)"`
	CreateTime  string `xorm:"default '2006-01-02 15:04:05' VARCHAR(30)"`
}
