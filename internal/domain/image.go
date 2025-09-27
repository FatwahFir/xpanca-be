package domain

import "time"

type Image struct {
	ID          uint `gorm:"primaryKey"`
	ProductID   uint
	URL         string
	IsThumbnail bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
