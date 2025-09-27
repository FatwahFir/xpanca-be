package dto

type CartItemRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Qty       int  `json:"qty,omitempty"`
}

type CartItemResponse struct {
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Category  string  `json:"category"`
	Price     int64   `json:"price"`
	Thumbnail *string `json:"thumbnail,omitempty"`
	Qty       int     `json:"qty"`
	LineTotal int64   `json:"line_total"`
}

type CartResponse struct {
	Items    []CartItemResponse `json:"items"`
	Subtotal int64              `json:"subtotal"`
	Count    int                `json:"count"`
}
