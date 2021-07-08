package models

type Topic struct {
	ID          int64 `gorm:"PRIMARY_KEY"`
	PublisherID int64
	Publisher   User `gorm:"foreignKey:PublisherID; column:publisher"`
}
