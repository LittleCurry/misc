package model

type WxScanBindTask struct {
	Id                int    `xorm:"not null pk autoincr unique INT(11)"`
	ActivityId        string `xorm:"default '' VARCHAR(50)"`
	ActivityName      string `xorm:"default '' VARCHAR(50)"`
	WxOpenid          string `xorm:"default '' VARCHAR(50)"`
	Mode              string `xorm:"default '' VARCHAR(20)"`
	Taskid            string `xorm:"default '' VARCHAR(50)"`
	MulTemplateCreate int    `xorm:"default 0 TINYINT(1)"`
	IsPay             int    `xorm:"default 0 TINYINT(1)"`
	TemplateUrl       string `xorm:"default '' VARCHAR(200)"`
	CreateTime        string `xorm:"default '1970-01-01 08:00:00' VARCHAR(45)"`
	UpdateTime        string `xorm:"default '1970-01-01 08:00:00' VARCHAR(45)"`
}
