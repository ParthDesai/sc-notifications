package stats

import (
	"log"

	"github.com/changer/khabar/db"
	"github.com/changer/khabar/utils"
)

type Stats struct {
	LastSeen    int64 `json:"last_seen"`
	TotalCount  int   `json:"total_count"`
	UnreadCount int   `json:"unread_count"`
	TotalUnread int   `json:"total_unread"`
}

func Save(dbConn *db.MConn, user string, appName string, org string) {
	stats_query := utils.M{
		"user":     user,
		"app_name": appName,
		"org":      org,
	}

	save_doc := utils.M{
		"user":      user,
		"app_name":  appName,
		"org":       org,
		"timestamp": utils.EpochNow(),
	}

	log.Println("hello Saving", stats_query)

	dbConn.Upsert(StatsCollection, stats_query, save_doc)
	return
}

func Get(dbConn *db.MConn, user string, appName string, org string) *Stats {
	var stats Stats

	stats_query := utils.M{}
	unread_query := utils.M{"is_read": false}
	unread_since_query := utils.M{"is_read": false}

	stats_query["user"] = user
	stats_query["app_name"] = appName
	stats_query["org"] = org

	unread_query["user"] = user
	unread_since_query["user"] = user

	if len(appName) > 0 {
		unread_query["app_name"] = appName
		unread_since_query["app_name"] = appName
	}

	if len(org) > 0 {
		unread_query["org"] = org
		unread_since_query["org"] = org
	}

	var last_seen LastSeen

	err := dbConn.GetOne(StatsCollection, stats_query, &last_seen)
	if err != nil {
		log.Println(err)
	}

	if last_seen.Timestamp > 0 {
		unread_since_query["created_on"] = utils.M{"$gt": last_seen.Timestamp}
	}

	stats.LastSeen = last_seen.Timestamp

	stats.TotalCount = dbConn.Count(StatsCollection, stats_query)
	stats.UnreadCount = dbConn.Count(StatsCollection, unread_since_query)
	stats.TotalUnread = dbConn.Count(StatsCollection, unread_query)

	return &stats
}
