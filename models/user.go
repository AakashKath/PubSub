package models

type User struct {
	ID        int64  `gorm:"PRIMARY_KEY"`
	FirstName string `gorm:"type:varchar(255)"`
	LastName  string `gorm:"type:varchar(255)"`
	Email     string `gorm:"type:varchar(255)"`
}
