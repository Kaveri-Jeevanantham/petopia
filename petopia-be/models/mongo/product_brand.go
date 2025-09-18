package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProductBrand represents brand information in MongoDB
type ProductBrand struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	BrandID         int                `bson:"brand_id" json:"brand_id"`
	BrandName       string             `bson:"brand_name" json:"brand_name"`
	Description     string             `bson:"description" json:"description"`
	LogoURL         string             `bson:"logo_url" json:"logo_url"`
	Website         string             `bson:"website" json:"website"`
	Country         string             `bson:"country" json:"country"`
	EstablishedYear int                `bson:"established_year" json:"established_year"`
	Categories      []string           `bson:"categories" json:"categories"`
	IsActive        bool               `bson:"is_active" json:"is_active"`
	CreatedAt       time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt       time.Time          `bson:"updated_at" json:"updated_at"`
}
