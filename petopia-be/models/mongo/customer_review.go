package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CustomerReview represents customer reviews in MongoDB
type CustomerReview struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProductID        int                `bson:"product_id" json:"product_id"`
	CustomerID       int                `bson:"customer_id" json:"customer_id"`
	CustomerName     string             `bson:"customer_name" json:"customer_name"`
	Rating           float64            `bson:"rating" json:"rating"`
	Title            string             `bson:"title" json:"title"`
	Comment          string             `bson:"comment" json:"comment"`
	Images           []string           `bson:"images" json:"images"`
	VerifiedPurchase bool               `bson:"verified_purchase" json:"verified_purchase"`
	HelpfulVotes     int                `bson:"helpful_votes" json:"helpful_votes"`
	Filters          map[string]string  `bson:"filters" json:"filters"`
	PetInfo          PetInfo            `bson:"pet_info" json:"pet_info"`
	CreatedAt        time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at" json:"updated_at"`
}

// PetInfo represents pet information in reviews
type PetInfo struct {
	PetType  string `bson:"pet_type" json:"pet_type"`
	PetBreed string `bson:"pet_breed" json:"pet_breed"`
	PetAge   int    `bson:"pet_age" json:"pet_age"`
	PetSize  string `bson:"pet_size" json:"pet_size"`
}
