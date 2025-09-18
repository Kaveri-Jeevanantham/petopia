package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProductDetails represents product information in MongoDB
type ProductDetails struct {
	ID             primitive.ObjectID     `bson:"_id,omitempty" json:"id"`
	ProductID      int                    `bson:"product_id" json:"product_id"`
	ProductName    string                 `bson:"product_name" json:"product_name"`
	Description    string                 `bson:"description" json:"description"`
	BrandID        int                    `bson:"brand_id" json:"brand_id"`
	BrandName      string                 `bson:"brand_name" json:"brand_name"`
	SellerID       int                    `bson:"seller_id" json:"seller_id"`
	Category       string                 `bson:"category" json:"category"`
	ItemDimensions map[string]interface{} `bson:"item_dimensions" json:"item_dimensions"`
	Price          float64                `bson:"price" json:"price"`
	Discount       float64                `bson:"discount" json:"discount"`
	Availability   bool                   `bson:"availability" json:"availability"`
	CreatedAt      time.Time              `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time              `bson:"updated_at" json:"updated_at"`
}
