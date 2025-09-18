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

// MongoBrandRepository implements BrandRepository using MongoDB
type MongoBrandRepository struct {
	collection *mongo.Collection
}

// NewMongoBrandRepository creates a new MongoDB brand repository
func NewMongoBrandRepository() BrandRepository {
	collection := db.GetMongoCollection("brands")
	return &MongoBrandRepository{collection: collection}
}

// GetCollection returns the MongoDB collection
func (r *MongoBrandRepository) GetCollection() *mongo.Collection {
	return r.collection
}

// Create inserts a new brand
func (r *MongoBrandRepository) Create(ctx context.Context, brand *mongomodels.ProductBrand) (*mongomodels.ProductBrand, error) {
	brand.CreatedAt = time.Now()
	brand.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, brand)
	if err != nil {
		return nil, err
	}

	brand.ID = result.InsertedID.(primitive.ObjectID)
	return brand, nil
}

// FindAll returns all brands matching the filters
func (r *MongoBrandRepository) FindAll(ctx context.Context, filters bson.M, page, limit int64) ([]mongomodels.ProductBrand, error) {
	skip := (page - 1) * limit
	options := options.Find().SetSkip(skip).SetLimit(limit)

	cursor, err := r.collection.Find(ctx, filters, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var brands []mongomodels.ProductBrand
	if err := cursor.All(ctx, &brands); err != nil {
		return nil, err
	}

	return brands, nil
}

// FindByID returns a brand by its ID
func (r *MongoBrandRepository) FindByID(ctx context.Context, id string) (*mongomodels.ProductBrand, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid brand ID")
	}

	var brand mongomodels.ProductBrand
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&brand)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("brand not found")
		}
		return nil, err
	}

	return &brand, nil
}

// Update modifies an existing brand
func (r *MongoBrandRepository) Update(ctx context.Context, id string, brand *mongomodels.ProductBrand) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid brand ID")
	}

	brand.UpdatedAt = time.Now()
	update := bson.M{
		"$set": bson.M{
			"brand_name":       brand.BrandName,
			"description":      brand.Description,
			"logo_url":         brand.LogoURL,
			"website":          brand.Website,
			"country":          brand.Country,
			"established_year": brand.EstablishedYear,
			"categories":       brand.Categories,
			"is_active":        brand.IsActive,
			"updated_at":       brand.UpdatedAt,
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("brand not found")
	}

	return nil
}

// Delete removes a brand
func (r *MongoBrandRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid brand ID")
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("brand not found")
	}

	return nil
}
