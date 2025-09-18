package dao

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	mongo_models "petopia-be/models/mongo"
)

// ProductBrandDAO handles database operations for ProductBrand
type ProductBrandDAO struct {
	collection *mongo.Collection
}

// NewProductBrandDAO creates a new ProductBrandDAO
func NewProductBrandDAO(collection *mongo.Collection) *ProductBrandDAO {
	return &ProductBrandDAO{
		collection: collection,
	}
}

// CreateProductBrand creates a new product brand record
func (dao *ProductBrandDAO) CreateProductBrand(ctx context.Context, brand *mongo_models.ProductBrand) error {
	brand.CreatedAt = time.Now()
	brand.UpdatedAt = time.Now()

	result, err := dao.collection.InsertOne(ctx, brand)
	if err != nil {
		return err
	}

	brand.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

// GetProductBrandByID retrieves a brand by its ID
func (dao *ProductBrandDAO) GetProductBrandByID(ctx context.Context, id primitive.ObjectID) (*mongo_models.ProductBrand, error) {
	var brand mongo_models.ProductBrand
	err := dao.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&brand)
	if err != nil {
		return nil, err
	}
	return &brand, nil
}

// GetProductBrandByBrandID retrieves a brand by its brand_id
func (dao *ProductBrandDAO) GetProductBrandByBrandID(ctx context.Context, brandID int) (*mongo_models.ProductBrand, error) {
	var brand mongo_models.ProductBrand
	err := dao.collection.FindOne(ctx, bson.M{"brand_id": brandID}).Decode(&brand)
	if err != nil {
		return nil, err
	}
	return &brand, nil
}

// GetActiveBrands retrieves all active brands
func (dao *ProductBrandDAO) GetActiveBrands(ctx context.Context) ([]*mongo_models.ProductBrand, error) {
	filter := bson.M{"is_active": true}
	opts := options.Find().SetSort(bson.D{{Key: "brand_name", Value: 1}})

	cursor, err := dao.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var brands []*mongo_models.ProductBrand
	for cursor.Next(ctx) {
		var brand mongo_models.ProductBrand
		if err := cursor.Decode(&brand); err != nil {
			return nil, err
		}
		brands = append(brands, &brand)
	}

	return brands, cursor.Err()
}

// GetBrandsByCategory retrieves brands by category
func (dao *ProductBrandDAO) GetBrandsByCategory(ctx context.Context, category string) ([]*mongo_models.ProductBrand, error) {
	filter := bson.M{
		"is_active":  true,
		"categories": bson.M{"$in": []string{category}},
	}

	cursor, err := dao.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var brands []*mongo_models.ProductBrand
	for cursor.Next(ctx) {
		var brand mongo_models.ProductBrand
		if err := cursor.Decode(&brand); err != nil {
			return nil, err
		}
		brands = append(brands, &brand)
	}

	return brands, cursor.Err()
}

// UpdateProductBrand updates a product brand record
func (dao *ProductBrandDAO) UpdateProductBrand(ctx context.Context, id primitive.ObjectID, updates bson.M) error {
	updates["updated_at"] = time.Now()

	_, err := dao.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": updates},
	)
	return err
}

// DeactivateProductBrand deactivates a product brand
func (dao *ProductBrandDAO) DeactivateProductBrand(ctx context.Context, id primitive.ObjectID) error {
	update := bson.M{
		"$set": bson.M{
			"is_active":  false,
			"updated_at": time.Now(),
		},
	}

	_, err := dao.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

// SearchProductBrands searches brands by name
func (dao *ProductBrandDAO) SearchProductBrands(ctx context.Context, searchTerm string) ([]*mongo_models.ProductBrand, error) {
	filter := bson.M{
		"$and": []bson.M{
			{"is_active": true},
			{
				"$or": []bson.M{
					{"brand_name": bson.M{RegexOperator: searchTerm, OptionsOperator: CaseInsensitive}},
					{"description": bson.M{RegexOperator: searchTerm, OptionsOperator: CaseInsensitive}},
				},
			},
		},
	}

	cursor, err := dao.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var brands []*mongo_models.ProductBrand
	for cursor.Next(ctx) {
		var brand mongo_models.ProductBrand
		if err := cursor.Decode(&brand); err != nil {
			return nil, err
		}
		brands = append(brands, &brand)
	}

	return brands, cursor.Err()
}
