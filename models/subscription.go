package models

type Subscription struct {
	ID     int64 `gorm:"PRIMARY_KEY"`
	UserID int64
	User   User `gorm:"foreignKey:UserID; column:user"`
}
