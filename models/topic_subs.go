package models

type TopicSubs struct {
	ID             int64 `gorm:"PRIMARY_KEY"`
	TopicID        int64
	Topic          Topic `gorm:"foreignKey:TopicID"`
	SubscriptionID int64
	Subscription   Subscription `gorm:"foreignKey:SubscriptionID"`
}
