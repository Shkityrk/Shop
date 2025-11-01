package dto

// CreateCartItemRequest DTO для создания элемента корзины
type CreateCartItemRequest struct {
	UserID    int `json:"user_id" binding:"required"`
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required,gt=0"`
}

// UpdateCartItemRequest DTO для обновления элемента корзины
type UpdateCartItemRequest struct {
	Quantity *int `json:"quantity,omitempty" binding:"omitempty,gt=0"`
}

// CartItemResponse DTO для ответа с данными элемента корзины
type CartItemResponse struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

