package dto

// ProductRequestDTO represents the data transfer object for a product request
type ProductRequestDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
