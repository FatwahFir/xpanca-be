package domain

import "time"

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"index"`
	Category    string `gorm:"index"`
	Price       int64
	Description string
	Images      []Image
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
