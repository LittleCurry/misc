package model

type Datareport struct {
	TaskId   string `bson:"taskId" json:"taskId"`
	Visit    int    `bson:"visit" json:"visit"`
	Play     int    `bson:"play" json:"play"`
	Download int    `bson:"download" json:"download"`
	Share    int    `bson:"share" json:"share"`
	Like     int    `bson:"like" json:"like"`
}
