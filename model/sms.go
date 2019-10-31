package model

type Sms struct {
	Id           int    `xorm:"not null pk autoincr unique INT(11)"`
	SmsId        string `xorm:"not null default '' unique VARCHAR(20)"`
	ActivityId   string `xorm:"default '' VARCHAR(20)"`
	ActivityName string `xorm:"default '' VARCHAR(30)"`
	Info         string `xorm:"default '' VARCHAR(200)"`
	IsResolved   string `xorm:"default 0 TINYINT(4)"`
	Schedule     string `xorm:"default '' VARCHAR(200)"`
	Reason       string `xorm:"default '' VARCHAR(200)"`
	CreateTime   string `xorm:"default '2006-01-02 15:04:05' VARCHAR(30)"`
	UpdateTime   string `xorm:"default '2006-01-02 15:04:05' VARCHAR(30)"`
}
