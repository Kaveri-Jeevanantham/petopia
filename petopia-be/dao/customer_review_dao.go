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

// CustomerReviewDAO handles database operations for CustomerReview
type CustomerReviewDAO struct {
	collection *mongo.Collection
}

// NewCustomerReviewDAO creates a new CustomerReviewDAO
func NewCustomerReviewDAO(collection *mongo.Collection) *CustomerReviewDAO {
	return &CustomerReviewDAO{
		collection: collection,
	}
}

// CreateCustomerReview creates a new customer review record
func (dao *CustomerReviewDAO) CreateCustomerReview(ctx context.Context, review *mongo_models.CustomerReview) error {
	review.CreatedAt = time.Now()
	review.UpdatedAt = time.Now()

	result, err := dao.collection.InsertOne(ctx, review)
	if err != nil {
		return err
	}

	review.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

// GetCustomerReviewByID retrieves a review by its ID
func (dao *CustomerReviewDAO) GetCustomerReviewByID(ctx context.Context, id primitive.ObjectID) (*mongo_models.CustomerReview, error) {
	var review mongo_models.CustomerReview
	err := dao.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&review)
	if err != nil {
		return nil, err
	}
	return &review, nil
}

// GetReviewsByProductID retrieves all reviews for a specific product
func (dao *CustomerReviewDAO) GetReviewsByProductID(ctx context.Context, productID int) ([]*mongo_models.CustomerReview, error) {
	filter := bson.M{"product_id": productID}
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := dao.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reviews []*mongo_models.CustomerReview
	for cursor.Next(ctx) {
		var review mongo_models.CustomerReview
		if err := cursor.Decode(&review); err != nil {
			return nil, err
		}
		reviews = append(reviews, &review)
	}

	return reviews, cursor.Err()
}

// GetReviewsByCustomerID retrieves all reviews by a specific customer
func (dao *CustomerReviewDAO) GetReviewsByCustomerID(ctx context.Context, customerID int) ([]*mongo_models.CustomerReview, error) {
	filter := bson.M{"customer_id": customerID}
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := dao.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reviews []*mongo_models.CustomerReview
	for cursor.Next(ctx) {
		var review mongo_models.CustomerReview
		if err := cursor.Decode(&review); err != nil {
			return nil, err
		}
		reviews = append(reviews, &review)
	}

	return reviews, cursor.Err()
}

// GetReviewsByRating retrieves reviews by minimum rating
func (dao *CustomerReviewDAO) GetReviewsByRating(ctx context.Context, productID int, minRating float64) ([]*mongo_models.CustomerReview, error) {
	filter := bson.M{
		"product_id": productID,
		"rating":     bson.M{"$gte": minRating},
	}
	opts := options.Find().SetSort(bson.D{{Key: "helpful_votes", Value: -1}})

	cursor, err := dao.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reviews []*mongo_models.CustomerReview
	for cursor.Next(ctx) {
		var review mongo_models.CustomerReview
		if err := cursor.Decode(&review); err != nil {
			return nil, err
		}
		reviews = append(reviews, &review)
	}

	return reviews, cursor.Err()
}

// GetVerifiedReviews retrieves only verified purchase reviews for a product
func (dao *CustomerReviewDAO) GetVerifiedReviews(ctx context.Context, productID int) ([]*mongo_models.CustomerReview, error) {
	filter := bson.M{
		"product_id":        productID,
		"verified_purchase": true,
	}
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := dao.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reviews []*mongo_models.CustomerReview
	for cursor.Next(ctx) {
		var review mongo_models.CustomerReview
		if err := cursor.Decode(&review); err != nil {
			return nil, err
		}
		reviews = append(reviews, &review)
	}

	return reviews, cursor.Err()
}

// UpdateCustomerReview updates a customer review record
func (dao *CustomerReviewDAO) UpdateCustomerReview(ctx context.Context, id primitive.ObjectID, updates bson.M) error {
	updates["updated_at"] = time.Now()

	_, err := dao.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": updates},
	)
	return err
}

// IncrementHelpfulVotes increments the helpful votes count for a review
func (dao *CustomerReviewDAO) IncrementHelpfulVotes(ctx context.Context, id primitive.ObjectID) error {
	update := bson.M{
		"$inc": bson.M{"helpful_votes": 1},
		"$set": bson.M{"updated_at": time.Now()},
	}

	_, err := dao.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

// DeleteCustomerReview deletes a customer review record
func (dao *CustomerReviewDAO) DeleteCustomerReview(ctx context.Context, id primitive.ObjectID) error {
	_, err := dao.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

// GetAverageRatingForProduct calculates average rating for a product
func (dao *CustomerReviewDAO) GetAverageRatingForProduct(ctx context.Context, productID int) (float64, error) {
	pipeline := []bson.M{
		{"$match": bson.M{"product_id": productID}},
		{"$group": bson.M{
			"_id":           nil,
			"averageRating": bson.M{"$avg": "$rating"},
		}},
	}

	cursor, err := dao.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	var result struct {
		AverageRating float64 `bson:"averageRating"`
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return 0, err
		}
		return result.AverageRating, nil
	}

	return 0, nil
}
