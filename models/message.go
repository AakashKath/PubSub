package models

import (
	"fmt"
	"time"

	"github.com/AakashKath/PubSub/lib"
)

type Message struct {
	ID        int64 `gorm:"PRIMARY_KEY"`
	TopicID   int64
	Topic     Topic `gorm:"foreignKey:TopicID"`
	CreatedAt time.Time
	Content   string
}

// Insert inserts a record to the "topic" table
func (m *Message) Insert() error {
	db := lib.DB.Connection
	if err := db.Create(&m).Error; err != nil {
		fmt.Println("Error while inserting record", "table", "message", "error", err.Error())
		return err
	}
	return nil
}

// Delete delete a record from the "topic" table
func (m *Message) Delete() error {
	db := lib.DB.Connection
	if err := db.Where("id = ?", m.ID).Delete(&m).Error; err != nil {
		fmt.Println("Error while deleting record", "table", "message", "error", err.Error())
		return err
	}
	return nil
}

// SelectByID returns a record from the "email" table using id
func (m *Message) SelectByID() error {
	db := lib.DB.Connection
	if err := db.Where("topic_id = ?", m.Topic.ID).Order("created_at desc").First(&m).Error; err != nil {
		fmt.Println("Error while selecting record", "table", "message", "error", err.Error())
		return err
	}
	return nil
}
