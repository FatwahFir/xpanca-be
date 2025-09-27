package dto

import "time"

type ProductQuery struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Search   string `search:"search"`
}

type ProductResponse struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Category    string          `json:"category"`
	Price       int64           `json:"price"`
	Description string          `json:"description"`
	Images      []ImageResponse `json:"images"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}
