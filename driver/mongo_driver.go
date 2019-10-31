package driver

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"strings"
)

var goalMongo *mgo.Database

func MongoInit(address string) {

	// address传进来 siiva:siiva0901@production.mongodb.rds.aliyuncs.com:3717/bt_prod
	// 要转换成 mongodb://siiva:siiva0901@production.mongodb.rds.aliyuncs.com:3717/admin
	// 在从admin切到bt_prod
	bt_prod := strings.Split(address, "/")[1]
	session, err1 := mgo.Dial("mongodb://"+strings.Replace(address, bt_prod, "admin", -1))
	if err1 != nil {
		fmt.Println("err1:", err1)
	}
	session.SetMode(mgo.Monotonic, true)
	goalMongo = session.DB(bt_prod)

}

func Mongo() *mgo.Database {
	return goalMongo
}