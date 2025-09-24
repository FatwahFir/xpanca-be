package dto

type ProductQuery struct {
	Page     int
	PageSize int
	Name     string
	Category string
	Search   string
}
