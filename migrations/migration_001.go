package migrations

import (
	"github.com/AakashKath/PubSub/lib"
	"github.com/AakashKath/PubSub/models"
)

func migrate001() {
	db := lib.DB.Connection
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Subscription{})
	db.AutoMigrate(&models.Topic{})
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.TopicSubs{})
}
