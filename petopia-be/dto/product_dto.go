package dto

// ProductDTO represents the data transfer object for a product
type ProductDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
