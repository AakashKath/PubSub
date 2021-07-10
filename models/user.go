package models

import (
	"fmt"

	"github.com/AakashKath/PubSub/lib"
)

type User struct {
	ID        int64  `gorm:"PRIMARY_KEY"`
	FirstName string `gorm:"type:varchar(255)"`
	LastName  string `gorm:"type:varchar(255)"`
	Email     string `gorm:"type:varchar(255)"`
}

// SelectByID returns a record from the "user" table using email(username)
func (u *User) SelectByID() error {
	db := lib.DB.Connection
	if err := db.Where("email = ?", u.Email).First(&u).Error; err != nil {
		fmt.Println("Error while selecting record", "table", "user", "error", err.Error())
		return err
	}
	return nil
}
