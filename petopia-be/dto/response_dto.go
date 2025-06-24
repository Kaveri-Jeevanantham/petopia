package dto

import "time"

type BaseResponse struct {
	CreatedAt  time.Time  `json:"created_at"`
	CreatedBy  *int       `json:"created_by,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	UpdatedBy  *int       `json:"updated_by,omitempty"`
	IsDeleted  bool       `json:"is_deleted"`
}

type UserAccountResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	BaseResponse
}

type UserAddressResponse struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Address string `json:"address"`
	Pincode string `json:"pincode"`
	BaseResponse
}

type UserCardDetailsResponse struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	CardNumber string `json:"card_number"`
	BaseResponse
}

type SellerInfoResponse struct {
	ID            int    `json:"id"`
	SellerName    string `json:"seller_name"`
	SellerInfo    string `json:"seller_info"`
	SellerAddress string `json:"seller_address"`
	BaseResponse
}

type ProductResponse struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	BaseResponse
}

type InventoryResponse struct {
	ID                int    `json:"id"`
	ProductID         int    `json:"product_id"`
	SellerID         int    `json:"seller_id"`
	InventoryLocation string `json:"inventory_location"`
	Quantity         int    `json:"quantity"`
	Product          *ProductResponse    `json:"product,omitempty"`
	Seller           *SellerInfoResponse `json:"seller,omitempty"`
	BaseResponse
}

type CartResponse struct {
	ID        int              `json:"id"`
	UserID    int              `json:"user_id"`
	ProductID int              `json:"product_id"`
	Quantity  int              `json:"quantity"`
	Product   *ProductResponse `json:"product,omitempty"`
	BaseResponse
}

type OrderResponse struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	Address    string  `json:"address"`
	TotalPrice float64 `json:"total_price"`
	Products   []ProductOrderedResponse `json:"products,omitempty"`
	BaseResponse
}

type ProductOrderedResponse struct {
	ID             int              `json:"id"`
	OrderID        int              `json:"order_id"`
	ProductID      int              `json:"product_id"`
	Quantity       int              `json:"quantity"`
	Price          float64          `json:"price"`
	DeliveryStatus string           `json:"delivery_status"`
	ReturnStatus   string           `json:"return_status"`
	Product        *ProductResponse `json:"product,omitempty"`
	BaseResponse
}

type ShippingResponse struct {
	ID                     int         `json:"id"`
	UserID                 int         `json:"user_id"`
	ProductID              int         `json:"product_id"`
	Address                string      `json:"address"`
	LocationTracking       interface{} `json:"location_tracking"`
	AssignedDeliveryPartner string    `json:"assigned_delivery_partner"`
	Product                *ProductResponse `json:"product,omitempty"`
	BaseResponse
}

type ListResponse struct {
	Total int         `json:"total"`
	Items interface{} `json:"items"`
}
