package model

type Task struct {
	ActivityId string      `bson:"activity_id" json:"activity_id"`
	Task       interface{} `bson:"task" json:"task"`
	State      string      `bson:"state" json:"state"`
	Del        string      `bson:"del" json:"del"`
	Msg        string      `bson:"msg" json:"msg"`
	Group      string      `bson:"group" json:"group"`
	WorkerId   string      `bson:"workerId" json:"workerId"`
	Likes      []string    `bson:"likes" json:"likes"`
	Collects   []string    `bson:"collects" json:"collects"`
	Visits     []string    `bson:"visits" json:"visits"`
	CreatedAt   string      `bson:"createdAt" json:"createdAt"`
	UpdatedAt   string      `bson:"updatedAt" json:"updatedAt"`
}
