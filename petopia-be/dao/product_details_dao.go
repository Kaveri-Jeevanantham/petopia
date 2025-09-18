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

// Constants for MongoDB operations
const (
	RegexOperator   = "$regex"
	OptionsOperator = "$options"
	CaseInsensitive = "i"
)

// ProductDetailsDAO handles database operations for ProductDetails
type ProductDetailsDAO struct {
	collection *mongo.Collection
}

// NewProductDetailsDAO creates a new ProductDetailsDAO
func NewProductDetailsDAO(collection *mongo.Collection) *ProductDetailsDAO {
	return &ProductDetailsDAO{
		collection: collection,
	}
}

// CreateProductDetails creates a new product details record
func (dao *ProductDetailsDAO) CreateProductDetails(ctx context.Context, product *mongo_models.ProductDetails) error {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	result, err := dao.collection.InsertOne(ctx, product)
	if err != nil {
		return err
	}

	product.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

// GetProductDetailsByID retrieves a product by its ID
func (dao *ProductDetailsDAO) GetProductDetailsByID(ctx context.Context, id primitive.ObjectID) (*mongo_models.ProductDetails, error) {
	var product mongo_models.ProductDetails
	err := dao.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetProductDetailsByProductID retrieves a product by its product_id
func (dao *ProductDetailsDAO) GetProductDetailsByProductID(ctx context.Context, productID int) (*mongo_models.ProductDetails, error) {
	var product mongo_models.ProductDetails
	err := dao.collection.FindOne(ctx, bson.M{"product_id": productID}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetProductDetailsByCategory retrieves products by category
func (dao *ProductDetailsDAO) GetProductDetailsByCategory(ctx context.Context, category string) ([]*mongo_models.ProductDetails, error) {
	filter := bson.M{"category": category, "availability": true}

	cursor, err := dao.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*mongo_models.ProductDetails
	for cursor.Next(ctx) {
		var product mongo_models.ProductDetails
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, cursor.Err()
}

// GetAvailableProductDetails retrieves all available products
func (dao *ProductDetailsDAO) GetAvailableProductDetails(ctx context.Context) ([]*mongo_models.ProductDetails, error) {
	filter := bson.M{"availability": true}
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := dao.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*mongo_models.ProductDetails
	for cursor.Next(ctx) {
		var product mongo_models.ProductDetails
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, cursor.Err()
}

// UpdateProductDetails updates a product details record
func (dao *ProductDetailsDAO) UpdateProductDetails(ctx context.Context, id primitive.ObjectID, updates bson.M) error {
	updates["updated_at"] = time.Now()

	_, err := dao.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": updates},
	)
	return err
}

// DeleteProductDetails deletes a product details record (soft delete by setting availability to false)
func (dao *ProductDetailsDAO) DeleteProductDetails(ctx context.Context, id primitive.ObjectID) error {
	update := bson.M{
		"$set": bson.M{
			"availability": false,
			"updated_at":   time.Now(),
		},
	}

	_, err := dao.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

// SearchProductDetails searches products by name or description
func (dao *ProductDetailsDAO) SearchProductDetails(ctx context.Context, searchTerm string) ([]*mongo_models.ProductDetails, error) {
	filter := bson.M{
		"$and": []bson.M{
			{"availability": true},
			{
				"$or": []bson.M{
					{"product_name": bson.M{RegexOperator: searchTerm, OptionsOperator: CaseInsensitive}},
					{"description": bson.M{RegexOperator: searchTerm, OptionsOperator: CaseInsensitive}},
					{"category": bson.M{RegexOperator: searchTerm, OptionsOperator: CaseInsensitive}},
					{"brand_name": bson.M{RegexOperator: searchTerm, OptionsOperator: CaseInsensitive}},
				},
			},
		},
	}

	cursor, err := dao.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*mongo_models.ProductDetails
	for cursor.Next(ctx) {
		var product mongo_models.ProductDetails
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, cursor.Err()
}
