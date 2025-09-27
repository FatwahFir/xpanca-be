package dto

type ProductQuery struct {
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Search   string `search:"search"`
}
