package repository

import (
	"context"
	mongomodels "petopia-be/models/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProductRepository defines the interface for product data operations
type ProductRepository interface {
	Create(ctx context.Context, product *mongomodels.ProductDetails) (*mongomodels.ProductDetails, error)
	FindAll(ctx context.Context, filters bson.M, page, limit int64) ([]mongomodels.ProductDetails, error)
	FindByID(ctx context.Context, id string) (*mongomodels.ProductDetails, error)
	FindByProductID(ctx context.Context, productID int) (*mongomodels.ProductDetails, error)
	FindByCategory(ctx context.Context, category string) ([]mongomodels.ProductDetails, error)
	FindAvailable(ctx context.Context, page, limit int64) ([]mongomodels.ProductDetails, int64, error)
	Update(ctx context.Context, id string, product *mongomodels.ProductDetails) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, term string) ([]mongomodels.ProductDetails, error)
	Count(ctx context.Context, filter bson.M) (int64, error)
	GetCollection() *mongo.Collection
}

// BrandRepository defines the interface for brand data operations
type BrandRepository interface {
	Create(ctx context.Context, brand *mongomodels.ProductBrand) (*mongomodels.ProductBrand, error)
	FindAll(ctx context.Context, filters bson.M, page, limit int64) ([]mongomodels.ProductBrand, error)
	FindByID(ctx context.Context, id string) (*mongomodels.ProductBrand, error)
	Update(ctx context.Context, id string, brand *mongomodels.ProductBrand) error
	Delete(ctx context.Context, id string) error
	GetCollection() *mongo.Collection
}
