package models

import (
	"fmt"

	"github.com/AakashKath/PubSub/lib"
)

type Topic struct {
	ID   int64  `gorm:"PRIMARY_KEY"`
	Name string `gorm:"type:varchar(255)"`
	// PublisherID int64
	// Publisher   User `gorm:"foreignKey:PublisherID"`
}

// Insert inserts a record to the "topic" table
func (t *Topic) Insert() error {
	db := lib.DB.Connection
	if err := db.Create(&t).Error; err != nil {
		fmt.Println("Error while inserting record", "table", "topic", "error", err.Error())
		return err
	}
	return nil
}

// Delete delete a record from the "topic" table
func (t *Topic) Delete() error {
	db := lib.DB.Connection
	if err := db.Where("name = ?", t.Name).Delete(&t).Error; err != nil {
		fmt.Println("Error while deleting record", "table", "topic", "error", err.Error())
		return err
	}
	return nil
}

// SelectByID returns a record from the "email" table using id
func (t *Topic) SelectByID() error {
	db := lib.DB.Connection
	if err := db.Where("name = ?", t.Name).First(&t).Error; err != nil {
		fmt.Println("Error while selecting record", "table", "topic", "error", err.Error())
		return err
	}
	return nil
}
