package models

// UserCreate represents a user registration request
type UserCreate struct {
	FirstName string `json:"first_name" binding:"required" example:"John"`
	LastName  string `json:"last_name" binding:"required" example:"Doe"`
	Username  string `json:"username" binding:"required" example:"johndoe"`
	Email     string `json:"email" binding:"required,email" example:"john@example.com"`
	Password  string `json:"password" binding:"required" example:"password123"`
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

