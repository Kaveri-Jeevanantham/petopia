package dto

type CreateUserAccountRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateUserAddressRequest struct {
	UserID  int    `json:"user_id" binding:"required"`
	Address string `json:"address" binding:"required"`
	Pincode string `json:"pincode" binding:"required"`
}

type CreateUserCardDetailsRequest struct {
	UserID     int    `json:"user_id" binding:"required"`
	CardNumber string `json:"card_number" binding:"required"`
}

type CreateSellerInfoRequest struct {
	SellerName    string `json:"seller_name" binding:"required"`
	SellerInfo    string `json:"seller_info"`
	SellerAddress string `json:"seller_address" binding:"required"`
}

type CreateProductRequest struct {
	ProductName string  `json:"product_name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,min=0"`
}

type CreateInventoryRequest struct {
	ProductID         int    `json:"product_id" binding:"required"`
	SellerID          int    `json:"seller_id" binding:"required"`
	InventoryLocation string `json:"inventory_location" binding:"required"`
	Quantity          int    `json:"quantity" binding:"required,min=0"`
}

type CreateCartRequest struct {
	UserID    int `json:"user_id" binding:"required"`
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required,min=1"`
}

type CreateOrderRequest struct {
	UserID     int     `json:"user_id" binding:"required"`
	Address    string  `json:"address" binding:"required"`
	TotalPrice float64 `json:"total_price" binding:"required,min=0"`
}

type CreateProductOrderedRequest struct {
	OrderID        int     `json:"order_id" binding:"required"`
	ProductID      int     `json:"product_id" binding:"required"`
	Quantity       int     `json:"quantity" binding:"required,min=1"`
	Price          float64 `json:"price" binding:"required,min=0"`
	DeliveryStatus string  `json:"delivery_status" binding:"required,oneof=Not Delivered Delivered"`
	ReturnStatus   string  `json:"return_status" binding:"required,oneof=Return Requested Returned None"`
}

type CreateShippingRequest struct {
	UserID                  int         `json:"user_id" binding:"required"`
	ProductID               int         `json:"product_id" binding:"required"`
	Address                 string      `json:"address" binding:"required"`
	LocationTracking        interface{} `json:"location_tracking"`
	AssignedDeliveryPartner string      `json:"assigned_delivery_partner"`
}

type UpdateDeliveryStatusRequest struct {
	DeliveryStatus string `json:"delivery_status" binding:"required,oneof=Not Delivered Delivered"`
}

type UpdateReturnStatusRequest struct {
	ReturnStatus string `json:"return_status" binding:"required,oneof=Return Requested Returned None"`
}

type UpdateInventoryQuantityRequest struct {
	Quantity int `json:"quantity" binding:"required,min=0"`
}

type StartMessageRequest struct {
	UserID    int    `json:"user_id" binding:"required"`
	ProductID int    `json:"product_id" binding:"required"`
	Message   string `json:"message" binding:"required"`
}
