package model

type Order struct {
	Openid           string              `bson:"openid" json:"openid"`
	OrderId          string              `bson:"order_id" json:"order_id"`
	CmbOrderId       string              `bson:"cmbOrderId" json:"cmbOrderId"`
	RefundCmbOrderId string              `bson:"refund_cmbOrderId" json:"refund_cmbOrderId"`
	IsPay            int                 `bson:"is_pay" json:"is_pay"`
	Seller           string              `bson:"seller" json:"seller"`
	TotalFee         string              `bson:"total_fee" json:"total_fee"`
	SiteId           string              `bson:"siteId" json:"siteId"`
	Refund           string              `bson:"refund" json:"refund"`
	Mark             string              `bson:"mark" json:"mark"`
	Time             string              `bson:"time" json:"time"`
	Taskid           string              `bson:"taskId" json:"taskId"`
	Templates        []map[string]string `bson:"templates" json:"templates"`     // 需要下载的照片
	Print            []map[string]string `bson:"print" json:"print"`             // 需要打印的照片
	Videos           []map[string]string `bson:"videos" json:"videos"`           // 需要合成的视频
	TaskIds          []string            `bson:"taskIds" json:"taskIds"`         // 照片多模板支付
	CarTaskids       []string            `bson:"car_taskIds" json:"car_taskIds"` // 购物车多task支付
	Uid              string              `bson:"uid" json:"uid"`
	ActivityId       string              `bson:"activity_id" json:"activity_id"`
	PrintId          string              `bson:"print_id" json:"print_id"`
	Phone            string              `bson:"phone" json:"phone"`
	CashCode         string              `bson:"cash_code" json:"cash_code"`
	Test             string              `bson:"test" json:"test"`
}

//type VPhotoTemplate struct {
//	Mode      string `json:"mode"`
//	ImgUrl    string `json:"img_url"`
//	Mov       string `json:"mov"`
//	Video_url string `json:"video_url"`
//}
