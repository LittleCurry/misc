package model

type Order struct {
	Openid     string   `bson:"openid" json:"openid"`
	OrderId    string   `bson:"order_id" json:"order_id"`
	IsPay      int      `bson:"is_pay" json:"is_pay"`
	Seller     string   `bson:"seller" json:"seller"`
	TotalFee   string   `bson:"total_fee" json:"total_fee"`
	SiteId     string   `bson:"siteId" json:"siteId"`
	Time       string   `bson:"time" json:"time"`
	Taskid     string   `bson:"taskId" json:"taskId"`
	Templates  []string `bson:"templates" json:"templates"`
	Print      []string `bson:"print" json:"print"`
	Uid        string   `bson:"uid" json:"uid"`
	ActivityId string   `bson:"activity_id" json:"activity_id"`
	PrintId    string   `bson:"print_id" json:"print_id"`
}
