package dto

import (
	"time"

	mongomodels "petopia-be/models/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MapProductToDTO converts a ProductDetails model to a ProductResponseDTO
func MapProductToDTO(product *mongomodels.ProductDetails) *ProductResponseDTO {
	if product == nil {
		return nil
	}

	return &ProductResponseDTO{
		ID:             product.ID.Hex(),
		ProductID:      product.ProductID,
		ProductName:    product.ProductName,
		Description:    product.Description,
		BrandID:        product.BrandID,
		BrandName:      product.BrandName,
		SellerID:       product.SellerID,
		Category:       product.Category,
		ItemDimensions: product.ItemDimensions,
		Price:          product.Price,
		Discount:       product.Discount,
		Availability:   product.Availability,
		CreatedAt:      product.CreatedAt,
		UpdatedAt:      product.UpdatedAt,
	}
}

// MapProductsToDTOs converts a slice of ProductDetails models to a slice of ProductResponseDTOs
func MapProductsToDTOs(products []mongomodels.ProductDetails) []ProductResponseDTO {
	productDTOs := make([]ProductResponseDTO, len(products))
	for i, product := range products {
		productDTOs[i] = *MapProductToDTO(&product)
	}
	return productDTOs
}

// MapDTOToProduct converts a ProductRequestDTO to a ProductDetails model
func MapDTOToProduct(dto *ProductRequestDTO, id ...string) *mongomodels.ProductDetails {
	if dto == nil {
		return nil
	}

	product := &mongomodels.ProductDetails{
		ProductName:    dto.ProductName,
		Description:    dto.Description,
		BrandID:        dto.BrandID,
		BrandName:      dto.BrandName,
		SellerID:       dto.SellerID,
		Category:       dto.Category,
		ItemDimensions: dto.ItemDimensions,
		Price:          dto.Price,
		Discount:       dto.Discount,
		Availability:   dto.Availability,
	}

	// Set times for new products
	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now

	// Set ID if provided
	if len(id) > 0 && id[0] != "" {
		objectID, err := primitive.ObjectIDFromHex(id[0])
		if err == nil {
			product.ID = objectID
		}
	}

	return product
}

// CreatePaginatedResponse creates a PaginatedResponse from items, total count, page and limit
func CreatePaginatedResponse(items interface{}, total, page, limit int64) *PaginatedResponse {
	totalPages := int64(1)
	if limit > 0 {
		totalPages = (total + limit - 1) / limit // Ceiling division
	}

	return &PaginatedResponse{
		Items:      items,
		Total:      total,
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
	}
}
