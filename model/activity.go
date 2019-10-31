package model

type Activity struct {
	UserId       string                   `bson:"user_id" json:"user_id"`
	CompanyId    string                   `bson:"company_id" json:"company_id"`
	ProjectId    string                   `bson:"project_id" json:"project_id"`
	ActivityId   string                   `bson:"activity_id" json:"activity_id"`
	Account      map[string]interface{}   `bson:"account" json:"account"`
	ActivityName string                   `bson:"activity_name" json:"activity_name"`
	Banner       string                   `bson:"banner" json:"banner"`
	Group        string                   `bson:"group" json:"group"`
	Type         string                   `bson:"type" json:"type"`
	Address      string                   `bson:"address" json:"address"`
	Mark         string                   `bson:"mark" json:"mark"`
	Lon          string                   `bson:"lon" json:"lon"`
	Lat          string                   `bson:"lat" json:"lat"`
	Cutconfigs   []map[string]interface{} `bson:"cutconfigs" json:"cutconfigs"`
	Settings     interface{}              `bson:"settings" json:"settings"`
	CreateTime   string                   `bson:"create_time" json:"create_time"`
	UpdateTime   string                   `bson:"update_time" json:"update_time"`
	Visits       []string                 `bson:"visits" json:"visits"`
	Partners     []string                 `bson:"partners" json:"partners"`
	Logo         string                   `bson:"logo" json:"logo"`
	Mateurl      string                   `bson:"mateurl" json:"mateurl"`
	Keyword      string                   `bson:"keyword" json:"keyword"`
	Secret       string                   `bson:"secret" json:"secret"`
	Leadpage     string                   `bson:"leadpage" json:"leadpage"`
	Activitys    []map[string]interface{} `bson:"activitys" json:"activitys"`
	Report       map[string]interface{}   `bson:"report" json:"report"`
}
