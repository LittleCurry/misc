package model

type SoccorSiteSetting struct {
	SiteId       string         `bson:"siteId" json:"siteId"`
	ActivityId   string         `bson:"activity_id" json:"activity_id"`
	SiteName     string         `bson:"siteName" json:"siteName"`
	SiteType     string         `bson:"siteType" json:"siteType"`
	DeviceConfig []deviceConfig `bson:"deviceConfig" json:"deviceConfig"`
	Group        string         `bson:"group" json:"group"`
}
