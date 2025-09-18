package dto

// ProductRequestDTO represents the data transfer object for a product request
type ProductRequestDTO struct {
	ProductName    string                 `json:"product_name" binding:"required"`
	Description    string                 `json:"description"`
	BrandID        int                    `json:"brand_id"`
	BrandName      string                 `json:"brand_name"`
	SellerID       int                    `json:"seller_id"`
	Category       string                 `json:"category"`
	ItemDimensions map[string]interface{} `json:"item_dimensions"`
	Price          float64                `json:"price" binding:"required"`
	Discount       float64                `json:"discount"`
	Availability   bool                   `json:"availability"`
}

// ProductFilter represents the filtering criteria for products
type ProductFilter struct {
	Category     string
	SearchTerm   string
	Page         int64
	Limit        int64
	SortBy       string
	SortOrder    string
	Availability *bool
}
