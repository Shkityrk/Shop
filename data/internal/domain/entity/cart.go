package entity

import "time"

// CartItem представляет доменную модель элемента корзины
type CartItem struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CartFilter представляет фильтры для поиска элементов корзины
type CartFilter struct {
	ID        *int
	UserID    *int
	ProductID *int
}

