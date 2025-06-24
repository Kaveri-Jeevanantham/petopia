package dao

import "gorm.io/gorm"

// Product represents the data access object for a product
type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
}
