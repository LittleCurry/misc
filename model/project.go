package model

type Project struct {
	Id            int    `xorm:"not null pk autoincr unique INT(11)"`
	ProjectId     string `xorm:"default '' VARCHAR(30)"`
	CompanyId     string `xorm:"default '' VARCHAR(30)"`
	ProjectName   string `xorm:"default '' VARCHAR(100)"`
	ProjectManger string `xorm:"default '' VARCHAR(100)"`
	Phone         string `xorm:"default '' VARCHAR(30)"`
	Passwd        string `xorm:"default '' VARCHAR(100)"`
	ChargeAccount string `xorm:"default '' VARCHAR(100)"`
	PartnerTo     string `xorm:"default '' VARCHAR(100)"`
	Banner        string `xorm:"default '' VARCHAR(200)"`
	Mark          string `xorm:"default '' VARCHAR(200)"`
	CreateTime    string `xorm:"default '' VARCHAR(30)"`
	UpdateTime    string `xorm:"default '' VARCHAR(30)"`
}

//type CompanyAndProject struct {
//	Project `xorm:"extends"`
//	Company `xorm:"extends"`
//}

type CompanyAndProject struct {
	Project            `xorm:"extends"`
	CompanyName string `xorm:"default '' VARCHAR(30)"`
}
