package dto

import "time"

type ImageResponse struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ProductID   uint      `json:"-"`
	URL         string    `json:"url"`
	IsThumbnail bool      `json:"is_thumbnail"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
