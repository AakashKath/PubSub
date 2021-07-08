package models

type Message struct {
	ID        int64 `gorm:"PRIMARY_KEY"`
	TopicID   int64
	Topic     Topic `gorm:"foreignKey:TopicID; column: topic"`
	MessageNo int64
	Content   string
}
