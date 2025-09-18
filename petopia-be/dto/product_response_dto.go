package dto

import "time"

// ProductResponseDTO represents the data transfer object for a product response
type ProductResponseDTO struct {
	ID             string                 `json:"id"`
	ProductID      int                    `json:"product_id"`
	ProductName    string                 `json:"product_name"`
	Description    string                 `json:"description"`
	BrandID        int                    `json:"brand_id"`
	BrandName      string                 `json:"brand_name"`
	SellerID       int                    `json:"seller_id"`
	Category       string                 `json:"category"`
	ItemDimensions map[string]interface{} `json:"item_dimensions"`
	Price          float64                `json:"price"`
	Discount       float64                `json:"discount"`
	Availability   bool                   `json:"availability"`
	CreatedAt      time.Time              `json:"created_at"`
	UpdatedAt      time.Time              `json:"updated_at"`
}

// PaginatedResponse represents a paginated response with metadata
type PaginatedResponse struct {
	Items      interface{} `json:"items"`
	Total      int64       `json:"total"`
	Page       int64       `json:"page"`
	Limit      int64       `json:"limit"`
	TotalPages int64       `json:"total_pages"`
}
