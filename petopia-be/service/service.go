package service

import (
	"context"
	"petopia-be/dto"
)

// Common welcome message function
func GetWelcomeMessage() string {
	return "Welcome to Petopia API!"
}

// ProductService defines the interface for product business operations
type ProductService interface {
	// Create a new product
	CreateProduct(ctx context.Context, requestDTO dto.ProductRequestDTO) (*dto.ProductResponseDTO, error)

	// Get a product by its ID
	GetProductByID(ctx context.Context, id string) (*dto.ProductResponseDTO, error)

	// Get a product by its product_id field
	GetProductByProductID(ctx context.Context, productID int) (*dto.ProductResponseDTO, error)

	// Get all products with pagination
	GetAllProducts(ctx context.Context, page, limit int64) (*dto.PaginatedResponse, error)

	// Get products by category
	GetProductsByCategory(ctx context.Context, category string) ([]dto.ProductResponseDTO, error)

	// Get available products with pagination
	GetAvailableProducts(ctx context.Context, page, limit int64) (*dto.PaginatedResponse, error)

	// Update a product
	UpdateProduct(ctx context.Context, id string, requestDTO dto.ProductRequestDTO) error

	// Delete a product
	DeleteProduct(ctx context.Context, id string) error

	// Search for products
	SearchProducts(ctx context.Context, term string) ([]dto.ProductResponseDTO, error)
}
