package models

import (
	"time"
)

type BaseModel struct {
	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  *int       `json:"created_by"`
	UpdatedAt  *time.Time `json:"updated_at"`
	UpdatedBy  *int       `json:"updated_by"`
	DeletedAt  *time.Time `json:"deleted_at"`
	DeletedBy  *int       `json:"deleted_by"`
	IsDeleted  bool       `json:"is_deleted"`
}

type UserAccount struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	BaseModel
}

type UserAddress struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Address string `json:"address"`
	Pincode string `json:"pincode"`
	BaseModel
}

type UserCardDetails struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	CardNumber string `json:"card_number"`
	BaseModel
}

type SellerInfo struct {
	ID            int    `json:"id"`
	SellerName    string `json:"seller_name"`
	SellerInfo    string `json:"seller_info"`
	SellerAddress string `json:"seller_address"`
	BaseModel
}

type Product struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	BaseModel
}

type InventoryInformation struct {
	ID                int    `json:"id"`
	ProductID         int    `json:"product_id"`
	SellerID         int    `json:"seller_id"`
	InventoryLocation string `json:"inventory_location"`
	Quantity         int    `json:"quantity"`
	BaseModel
}

type Cart struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
	BaseModel
}

type Order struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	Address    string  `json:"address"`
	TotalPrice float64 `json:"total_price"`
	BaseModel
}

type ProductOrdered struct {
	ID             int     `json:"id"`
	OrderID        int     `json:"order_id"`
	ProductID      int     `json:"product_id"`
	Quantity       int     `json:"quantity"`
	Price          float64 `json:"price"`
	DeliveryStatus string  `json:"delivery_status"`
	ReturnStatus   string  `json:"return_status"`
	BaseModel
}

type Shipping struct {
	ID                     int         `json:"id"`
	UserID                 int         `json:"user_id"`
	ProductID              int         `json:"product_id"`
	Address                string      `json:"address"`
	LocationTracking       interface{} `json:"location_tracking"`
	AssignedDeliveryPartner string    `json:"assigned_delivery_partner"`
	BaseModel
}
