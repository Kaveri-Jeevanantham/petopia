package service

import (
	"context"
	"petopia-be/dao"
	"petopia-be/db"
	"petopia-be/dto"
	"petopia-be/repository"
)

// ServiceContainer holds all services
type ServiceContainer struct {
	ProductService ProductService
}

// NewServiceContainer creates a new service container with all services
func NewServiceContainer() *ServiceContainer {
	// Initialize repositories
	productRepo := repository.NewMongoProductRepository()

	// Initialize DAOs
	brandCollection := db.GetMongoCollection("brands")
	reviewCollection := db.GetMongoCollection("reviews")

	brandDAO := dao.NewProductBrandDAO(brandCollection)
	reviewDAO := dao.NewCustomerReviewDAO(reviewCollection)

	// Create services
	productService := NewProductServiceV2(productRepo, brandDAO, reviewDAO)

	return &ServiceContainer{
		ProductService: productService,
	}
}

// ProductServiceV2 is the new implementation of ProductService using the repository pattern
type ProductServiceV2 struct {
	productRepo repository.ProductRepository
	brandDAO    *dao.ProductBrandDAO
	reviewDAO   *dao.CustomerReviewDAO
}

// NewProductServiceV2 creates a new ProductServiceV2
func NewProductServiceV2(
	productRepo repository.ProductRepository,
	brandDAO *dao.ProductBrandDAO,
	reviewDAO *dao.CustomerReviewDAO,
) ProductService {
	return &ProductServiceV2{
		productRepo: productRepo,
		brandDAO:    brandDAO,
		reviewDAO:   reviewDAO,
	}
}

// Implementation of ProductService interface using the repository pattern
// CreateProduct implements ProductService
func (s *ProductServiceV2) CreateProduct(ctx context.Context, requestDTO dto.ProductRequestDTO) (*dto.ProductResponseDTO, error) {
	// Convert DTO to model
	product := dto.MapDTOToProduct(&requestDTO)

	// Check if brand exists and set brand name
	if product.BrandID > 0 && product.BrandName == "" {
		brand, err := s.brandDAO.GetProductBrandByBrandID(ctx, product.BrandID)
		if err == nil {
			product.BrandName = brand.BrandName
		}
	}

	// Create product
	createdProduct, err := s.productRepo.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	// Convert to response DTO
	return dto.MapProductToDTO(createdProduct), nil
}

// GetProductByID implements ProductService
func (s *ProductServiceV2) GetProductByID(ctx context.Context, id string) (*dto.ProductResponseDTO, error) {
	product, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return dto.MapProductToDTO(product), nil
}

// GetProductByProductID implements ProductService
func (s *ProductServiceV2) GetProductByProductID(ctx context.Context, productID int) (*dto.ProductResponseDTO, error) {
	product, err := s.productRepo.FindByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	return dto.MapProductToDTO(product), nil
}

// GetAllProducts implements ProductService
func (s *ProductServiceV2) GetAllProducts(ctx context.Context, page, limit int64) (*dto.PaginatedResponse, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	products, err := s.productRepo.FindAll(ctx, nil, page, limit)
	if err != nil {
		return nil, err
	}

	total, err := s.productRepo.Count(ctx, nil)
	if err != nil {
		return nil, err
	}

	productDTOs := dto.MapProductsToDTOs(products)
	return dto.CreatePaginatedResponse(productDTOs, total, page, limit), nil
}

// GetProductsByCategory implements ProductService
func (s *ProductServiceV2) GetProductsByCategory(ctx context.Context, category string) ([]dto.ProductResponseDTO, error) {
	products, err := s.productRepo.FindByCategory(ctx, category)
	if err != nil {
		return nil, err
	}
	return dto.MapProductsToDTOs(products), nil
}

// GetAvailableProducts implements ProductService
func (s *ProductServiceV2) GetAvailableProducts(ctx context.Context, page, limit int64) (*dto.PaginatedResponse, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	products, total, err := s.productRepo.FindAvailable(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	productDTOs := dto.MapProductsToDTOs(products)
	return dto.CreatePaginatedResponse(productDTOs, total, page, limit), nil
}

// UpdateProduct implements ProductService
func (s *ProductServiceV2) UpdateProduct(ctx context.Context, id string, requestDTO dto.ProductRequestDTO) error {
	// Check if product exists
	_, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// Check if brand exists and set brand name
	if requestDTO.BrandID > 0 && requestDTO.BrandName == "" {
		brand, err := s.brandDAO.GetProductBrandByBrandID(ctx, requestDTO.BrandID)
		if err == nil {
			requestDTO.BrandName = brand.BrandName
		}
	}

	// Update product fields
	product := dto.MapDTOToProduct(&requestDTO, id)

	return s.productRepo.Update(ctx, id, product)
}

// DeleteProduct implements ProductService
func (s *ProductServiceV2) DeleteProduct(ctx context.Context, id string) error {
	return s.productRepo.Delete(ctx, id)
}

// SearchProducts implements ProductService
func (s *ProductServiceV2) SearchProducts(ctx context.Context, term string) ([]dto.ProductResponseDTO, error) {
	products, err := s.productRepo.Search(ctx, term)
	if err != nil {
		return nil, err
	}
	return dto.MapProductsToDTOs(products), nil
}
