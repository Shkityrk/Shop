package models

// UserCreate represents a user registration request
type UserCreate struct {
	FirstName string `json:"first_name" binding:"required" example:"John"`
	LastName  string `json:"last_name" binding:"required" example:"Doe"`
	Username  string `json:"username" binding:"required" example:"johndoe"`
	Email     string `json:"email" binding:"required,email" example:"john@example.com"`
	Password  string `json:"password" binding:"required" example:"password123"`
	UserRole  string `json:"user_role" example:"client"`
}

// UserLogin represents a user login request
type UserLogin struct {
	Username string `json:"username" binding:"required" example:"johndoe"`
	Password string `json:"password" binding:"required" example:"password123"`
}

// UserOut represents a user response
type UserOut struct {
	ID        int    `json:"id" example:"1"`
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
	Username  string `json:"username" example:"johndoe"`
	Email     string `json:"email" example:"john@example.com"`
	UserRole  string `json:"user_role" example:"client"`
}

// Staff represents a staff member (user with role != client)
type Staff struct {
	ID        int    `json:"id" example:"1"`
	FirstName string `json:"first_name" example:"John"`
	LastName  string `json:"last_name" example:"Doe"`
	UserRole  string `json:"user_role" example:"staff"`
}

// ProductCreate represents a product creation request
type ProductCreate struct {
	Name             string  `json:"name" binding:"required" example:"Chocolate Cake"`
	ShortDescription string  `json:"short_description" binding:"required" example:"Delicious chocolate cake"`
	FullDescription  string  `json:"full_description" binding:"required" example:"A rich and moist chocolate cake with chocolate frosting"`
	Composition      string  `json:"composition" binding:"required" example:"Flour, Sugar, Cocoa, Eggs, Butter"`
	Weight           float64 `json:"weight" binding:"required" example:"500"`
	Price            float64 `json:"price" binding:"required" example:"25.99"`
	Photo            string  `json:"photo" binding:"required" example:"https://example.com/cake.jpg"`
}

// Product represents a product response
type Product struct {
	ID               int     `json:"id" example:"1"`
	Name             string  `json:"name" example:"Chocolate Cake"`
	ShortDescription string  `json:"short_description" example:"Delicious chocolate cake"`
	FullDescription  string  `json:"full_description" example:"A rich and moist chocolate cake with chocolate frosting"`
	Composition      string  `json:"composition" example:"Flour, Sugar, Cocoa, Eggs, Butter"`
	Weight           float64 `json:"weight" example:"500"`
	Price            float64 `json:"price" example:"25.99"`
	Photo            string  `json:"photo" example:"https://example.com/cake.jpg"`
}

// CartItemCreate represents a cart item creation request
type CartItemCreate struct {
	ProductID int `json:"product_id" binding:"required" example:"1"`
	Quantity  int `json:"quantity" binding:"required" example:"2"`
}

// CartItem represents a cart item response
type CartItem struct {
	ID        int `json:"id" example:"1"`
	UserID    int `json:"user_id" example:"1"`
	ProductID int `json:"product_id" example:"1"`
	Quantity  int `json:"quantity" example:"2"`
}

// WmsItem represents a line to check/commit in WMS
type WmsItem struct {
	ProductID int `json:"product_id" binding:"required" example:"1"`
	Quantity  int `json:"quantity" binding:"required" example:"2"`
}

// WmsRequest holds items for WMS check/commit
type WmsRequest struct {
	Items []WmsItem `json:"items" binding:"required,dive"`
}

// ShipmentItemCreate represents an item in shipment creation
type ShipmentItemCreate struct {
	ProductID int `json:"product_id" binding:"required" example:"1"`
	Quantity  int `json:"quantity" binding:"required" example:"2"`
}

// ShipmentCreate represents create shipment payload
type ShipmentCreate struct {
	OrderID   int                  `json:"order_id" binding:"required" example:"1001"`
	UserID    int                  `json:"user_id" binding:"required" example:"5"`
	Address   string               `json:"address" binding:"required" example:"Some street 1"`
	CourierID *int                 `json:"courier_id,omitempty" example:"10"`
	Items     []ShipmentItemCreate `json:"items" binding:"required,dive"`
}

// ShipmentItem represents shipment item response
type ShipmentItem struct {
	ID        int `json:"id" example:"1"`
	ProductID int `json:"product_id" example:"1"`
	Quantity  int `json:"quantity" example:"2"`
}

// Shipment represents shipment response
type Shipment struct {
	ID           int            `json:"id" example:"1"`
	OrderID      int            `json:"order_id" example:"1001"`
	UserID       int            `json:"user_id" example:"5"`
	Address      string         `json:"address" example:"Some street 1"`
	TrackingCode string         `json:"tracking_code" example:"abc123"`
	Status       string         `json:"status" example:"created"`
	CourierID    *int           `json:"courier_id,omitempty" example:"10"`
	CreatedAt    string         `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt    string         `json:"updated_at" example:"2024-01-01T00:00:00Z"`
	Items        []ShipmentItem `json:"items"`
}

// ShipmentStatusUpdate represents status update payload
type ShipmentStatusUpdate struct {
	Status    string `json:"status" binding:"required" example:"assigned"`
	CourierID *int   `json:"courier_id,omitempty" example:"10"`
}

// MessageResponse represents a simple message response
type MessageResponse struct {
	Message string `json:"message" example:"Success"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error" example:"Error message"`
}

// VerifyResponse represents a verification response
type VerifyResponse struct {
	Exists bool `json:"exists" example:"true"`
}

// WarehouseCreate represents a warehouse creation request
type WarehouseCreate struct {
	Name         string  `json:"name" binding:"required" example:"Main Warehouse"`
	Address      string  `json:"address" binding:"required" example:"123 Storage St"`
	WorkingHours *string `json:"working_hours,omitempty" example:"9:00-18:00"`
}

// Warehouse represents a warehouse response
type Warehouse struct {
	ID           int     `json:"id" example:"1"`
	Name         string  `json:"name" example:"Main Warehouse"`
	Address      string  `json:"address" example:"123 Storage St"`
	WorkingHours *string `json:"working_hours,omitempty" example:"9:00-18:00"`
}

// StorageRuleCreate represents a storage rule creation request
type StorageRuleCreate struct {
	Name        string   `json:"name" binding:"required" example:"Hazardous Materials"`
	Description *string  `json:"description,omitempty" example:"Rules for hazardous materials storage"`
	IsHazardous bool     `json:"is_hazardous" example:"true"`
	IsOversized bool     `json:"is_oversized" example:"false"`
	TempMin     *float64 `json:"temp_min,omitempty" example:"10.0"`
	TempMax     *float64 `json:"temp_max,omitempty" example:"25.0"`
}

// StorageRule represents a storage rule response
type StorageRule struct {
	ID          int      `json:"id" example:"1"`
	Name        string   `json:"name" example:"Hazardous Materials"`
	Description *string  `json:"description,omitempty" example:"Rules for hazardous materials storage"`
	IsHazardous bool     `json:"is_hazardous" example:"true"`
	IsOversized bool     `json:"is_oversized" example:"false"`
	TempMin     *float64 `json:"temp_min,omitempty" example:"10.0"`
	TempMax     *float64 `json:"temp_max,omitempty" example:"25.0"`
}

// BinLocationCreate represents a bin location creation request
type BinLocationCreate struct {
	WarehouseID   int    `json:"warehouse_id" binding:"required" example:"1"`
	Zone          string `json:"zone" binding:"required" example:"A"`
	Aisle         string `json:"aisle" binding:"required" example:"1"`
	Rack          string `json:"rack" binding:"required" example:"1"`
	BinCode       string `json:"bin_code" binding:"required" example:"A-1-1-01"`
	StorageRuleID *int   `json:"storage_rule_id,omitempty" example:"1"`
	ProductID     *int   `json:"product_id,omitempty" example:"1"`
	Quantity      *int   `json:"quantity,omitempty" example:"100"`
}

// BinLocation represents a bin location response
type BinLocation struct {
	ID            int    `json:"id" example:"1"`
	WarehouseID   int    `json:"warehouse_id" example:"1"`
	Zone          string `json:"zone" example:"A"`
	Aisle         string `json:"aisle" example:"1"`
	Rack          string `json:"rack" example:"1"`
	BinCode       string `json:"bin_code" example:"A-1-1-01"`
	StorageRuleID *int   `json:"storage_rule_id,omitempty" example:"1"`
	ProductID     *int   `json:"product_id,omitempty" example:"1"`
	Quantity      *int   `json:"quantity,omitempty" example:"100"`
}

// InventoryItemCreate represents an inventory item creation request
type InventoryItemCreate struct {
	ProductID   int `json:"product_id" binding:"required" example:"1"`
	WarehouseID int `json:"warehouse_id" binding:"required" example:"1"`
	BinID       int `json:"bin_id" binding:"required" example:"1"`
	Quantity    int `json:"quantity" binding:"required" example:"100"`
}

// InventoryItem represents an inventory item response
type InventoryItem struct {
	ID          int `json:"id" example:"1"`
	ProductID   int `json:"product_id" example:"1"`
	WarehouseID int `json:"warehouse_id" example:"1"`
	BinID       int `json:"bin_id" example:"1"`
	Quantity    int `json:"quantity" example:"100"`
}

// InventoryMoveRequest represents a request to move inventory between bins
type InventoryMoveRequest struct {
	ProductID       int `json:"product_id" binding:"required" example:"1"`
	FromWarehouseID int `json:"from_warehouse_id" binding:"required" example:"1"`
	FromBinID       int `json:"from_bin_id" binding:"required" example:"1"`
	ToWarehouseID   int `json:"to_warehouse_id" binding:"required" example:"2"`
	ToBinID         int `json:"to_bin_id" binding:"required" example:"2"`
	Quantity        int `json:"quantity" binding:"required" example:"10"`
}

// ProductTotalQuantity represents total quantity for a product
type ProductTotalQuantity struct {
	ProductID     int `json:"product_id" example:"1"`
	TotalQuantity int `json:"total_quantity" example:"500"`
}

// StatusResponse represents a simple status response
type StatusResponse struct {
	Status string `json:"status" example:"deleted"`
}

