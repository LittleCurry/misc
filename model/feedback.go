package model

type Feedback struct {
	Id           int    `xorm:"not null pk autoincr unique INT(11)"`
	FeedId       string `xorm:"not null default '' unique VARCHAR(30)"`
	WxOpenid     string `xorm:"default '' VARCHAR(30)"`
	ActivityId   string `xorm:"default '' VARCHAR(30)"`
	ActivityName string `xorm:"default '' VARCHAR(30)"`
	ProjectId    string `xorm:"default '' VARCHAR(30)"`
	ProjectName  string `xorm:"default '' VARCHAR(30)"`
	OrderId      string `xorm:"default '' VARCHAR(30)"`
	Taskid       string `xorm:"default '' VARCHAR(40)"`
	TotalFee     string `xorm:"default '' VARCHAR(30)"`
	Name         string `xorm:"default '' VARCHAR(20)"`
	Phone        string `xorm:"default '' VARCHAR(20)"`
	Email        string `xorm:"default '' VARCHAR(30)"`
	Problem      string `xorm:"default '' VARCHAR(200)"`
	Url          string `xorm:"default '' VARCHAR(200)"`
	Refund       string `xorm:"default '' VARCHAR(10)"`
	FixType      string `xorm:"default '' VARCHAR(10)"`
	FixInfo      string `xorm:"default '' VARCHAR(200)"`
	Other        string `xorm:"default '' VARCHAR(200)"`
	CreateTime   string `xorm:"default '2006-01-02 15:04:05' VARCHAR(30)"`
	UpdateTime   string `xorm:"default '2006-01-02 15:04:05' VARCHAR(30)"`
}
