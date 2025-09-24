package domain

import "time"

type Image struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ProductID   uint      `json:"-"`
	URL         string    `json:"url"`
	IsThumbnail bool      `json:"is_thumbnail"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"index"`
	Category    string    `json:"category" gorm:"index"`
	Price       int64     `json:"price"`
	Description string    `json:"description"`
	Images      []Image   `json:"images"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"size:64;uniqueIndex"`
	Password string `json:"-" gorm:"size:255"`
	Role     string `json:"role" gorm:"size:32"`
}
