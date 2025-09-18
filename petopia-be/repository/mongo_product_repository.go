package repository

import (
	"context"
	"errors"
	"petopia-be/db"
	mongomodels "petopia-be/models/mongo"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Error constants to avoid duplication
const (
	ErrProductNotFound  = "product not found"
	ErrInvalidProductID = "invalid product ID"
)

// MongoDB regex constants
const (
	RegexKey     = "$regex"
	OptionsKey   = "$options"
	CaseOption   = "i" // Case insensitive
)

// MongoProductRepository implements ProductRepository using MongoDB
type MongoProductRepository struct {
	collection *mongo.Collection
}

// NewMongoProductRepository creates a new MongoDB product repository
func NewMongoProductRepository() ProductRepository {
	collection := db.GetMongoCollection("products")
	return &MongoProductRepository{collection: collection}
}

// GetCollection returns the MongoDB collection
func (r *MongoProductRepository) GetCollection() *mongo.Collection {
	return r.collection
}

// Create inserts a new product
func (r *MongoProductRepository) Create(ctx context.Context, product *mongomodels.ProductDetails) (*mongomodels.ProductDetails, error) {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}

	product.ID = result.InsertedID.(primitive.ObjectID)
	return product, nil
}

// FindAll returns all products matching the filters
func (r *MongoProductRepository) FindAll(ctx context.Context, filters bson.M, page, limit int64) ([]mongomodels.ProductDetails, error) {
	skip := (page - 1) * limit
	options := options.Find().SetSkip(skip).SetLimit(limit)

	cursor, err := r.collection.Find(ctx, filters, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []mongomodels.ProductDetails
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}

// FindByID returns a product by its ID
func (r *MongoProductRepository) FindByID(ctx context.Context, id string) (*mongomodels.ProductDetails, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid product ID")
	}

	var product mongomodels.ProductDetails
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(ErrProductNotFound)
		}
		return nil, err
	}

	return &product, nil
}

// Update modifies an existing product
func (r *MongoProductRepository) Update(ctx context.Context, id string, product *mongomodels.ProductDetails) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New(ErrInvalidProductID)
	}

	product.UpdatedAt = time.Now()
	update := bson.M{
		"$set": bson.M{
			"product_name":    product.ProductName,
			"description":     product.Description,
			"brand_id":        product.BrandID,
			"brand_name":      product.BrandName,
			"seller_id":       product.SellerID,
			"category":        product.Category,
			"item_dimensions": product.ItemDimensions,
			"price":           product.Price,
			"discount":        product.Discount,
			"availability":    product.Availability,
			"updated_at":      product.UpdatedAt,
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New(ErrProductNotFound)
	}

	return nil
}

// Delete removes a product
func (r *MongoProductRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New(ErrInvalidProductID)
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New(ErrProductNotFound)
	}

	return nil
}

// FindByProductID retrieves a product by its product_id
func (r *MongoProductRepository) FindByProductID(ctx context.Context, productID int) (*mongomodels.ProductDetails, error) {
	var product mongomodels.ProductDetails
	err := r.collection.FindOne(ctx, bson.M{"product_id": productID}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &product, nil
}

// FindByCategory retrieves products by category
func (r *MongoProductRepository) FindByCategory(ctx context.Context, category string) ([]mongomodels.ProductDetails, error) {
	filter := bson.M{"category": category, "availability": true}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []mongomodels.ProductDetails
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}

// FindAvailable retrieves all available products with pagination
func (r *MongoProductRepository) FindAvailable(ctx context.Context, page, limit int64) ([]mongomodels.ProductDetails, int64, error) {
	filter := bson.M{"availability": true}

	// Get total count for pagination
	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	// Configure pagination
	skip := (page - 1) * limit
	opts := options.Find().
		SetSort(bson.D{{Key: "created_at", Value: -1}}).
		SetSkip(skip).
		SetLimit(limit)

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var products []mongomodels.ProductDetails
	if err := cursor.All(ctx, &products); err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

// Search finds products by name, description, etc.
func (r *MongoProductRepository) Search(ctx context.Context, term string) ([]mongomodels.ProductDetails, error) {
	filter := bson.M{
		"$and": []bson.M{
			{"availability": true},
			{
				"$or": []bson.M{
					{"product_name": bson.M{"$regex": term, "$options": "i"}},
					{"description": bson.M{"$regex": term, "$options": "i"}},
					{"category": bson.M{"$regex": term, "$options": "i"}},
					{"brand_name": bson.M{"$regex": term, "$options": "i"}},
				},
			},
		},
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []mongomodels.ProductDetails
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}

// Count returns the number of documents matching a filter
func (r *MongoProductRepository) Count(ctx context.Context, filter bson.M) (int64, error) {
	return r.collection.CountDocuments(ctx, filter)
}
