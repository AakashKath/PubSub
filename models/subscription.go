package models

import (
	"fmt"

	"github.com/AakashKath/PubSub/lib"
)

type Subscription struct {
	ID      int64 `gorm:"PRIMARY_KEY"`
	TopicID int64
	Topic   Topic  `gorm:"foreignKey:TopicID"`
	Name    string `gorm:"type:varchar(255)"`
}

// Insert inserts a record to the "topic" table
func (s *Subscription) Insert() error {
	db := lib.DB.Connection
	if err := db.Create(&s).Error; err != nil {
		fmt.Println("Error while inserting record", "table", "subscription", "error", err.Error())
		return err
	}
	return nil
}

// Delete delete a record from the "topic" table
func (s *Subscription) Delete() error {
	db := lib.DB.Connection
	if err := db.Where("name = ?", s.Name).Delete(&s).Error; err != nil {
		fmt.Println("Error while deleting record", "table", "subscription", "error", err.Error())
		return err
	}
	return nil
}

// SelectByID returns a record from the "email" table using id
func (s *Subscription) SelectByID() error {
	db := lib.DB.Connection
	if err := db.Where("name = ?", s.Name).First(&s).Error; err != nil {
		fmt.Println("Error while selecting record", "table", "subscription", "error", err.Error())
		return err
	}
	return nil
}

type Subscriber struct {
	ID             int64
	SubscriptionID int64
	Subscription   Subscription `gorm:"foreignKey:SubscriptionID"`
	UserID         int64
	UserSub        User `gorm:"foreignKey:UserID"`
	IsSubscribed   bool `gorm:"default:True"`
}

// Insert inserts a record to the "topic" table
func (s *Subscriber) Insert() error {
	db := lib.DB.Connection
	if err := db.Create(&s).Error; err != nil {
		fmt.Println("Error while inserting record", "table", "subscriber", "error", err.Error())
		return err
	}
	return nil
}

// Delete delete a record from the "topic" table
func (s *Subscriber) Delete() error {
	db := lib.DB.Connection
	if err := db.Where("subscription_id = ?", s.SubscriptionID).Delete(&s).Error; err != nil {
		fmt.Println("Error while deleting record", "table", "subscriber", "error", err.Error())
		return err
	}
	return nil
}

// Delete delete a record from the "topic" table
func (s *Subscriber) SetSubscribe() error {
	db := lib.DB.Connection
	if err := db.Model(&s).Update("is_subscribed", "t").Error; err != nil {
		fmt.Println("Error while subscribing record", "table", "subscriber", "error", err.Error())
		return err
	}
	return nil
}

// Delete delete a record from the "topic" table
func (s *Subscriber) SetUnsubscribe() error {
	db := lib.DB.Connection
	if err := db.Model(&s).Update("is_subscribed", "f").Error; err != nil {
		fmt.Println("Error while unsubscribing record", "table", "subscriber", "error", err.Error())
		return err
	}
	return nil
}

// SelectByID returns a record from the "email" table using id
func (s *Subscriber) SelectByID() error {
	db := lib.DB.Connection
	if err := db.Where("subscription_id = ? and user_id=?", s.Subscription.ID, s.UserSub.ID).First(&s).Error; err != nil {
		fmt.Println("Error while selecting record", "table", "subscriber", "error", err.Error())
		return err
	}
	return nil
}
