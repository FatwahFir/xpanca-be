package domain

import "time"

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"uniqueIndex"`
	Items     []CartItem
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartItem struct {
	ID        uint `gorm:"primaryKey"`
	CartID    uint `gorm:"index"`
	ProductID uint `gorm:"index"`
	Qty       int
	Product   Product
	CreatedAt time.Time
	UpdatedAt time.Time
}
