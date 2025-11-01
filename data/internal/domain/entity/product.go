package entity

import "time"

// Product представляет доменную модель продукта
type Product struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	ShortDescription string  `json:"short_description"`
	FullDescription  string  `json:"full_description"`
	Composition      string  `json:"composition"`
	Weight           float64 `json:"weight"`
	Price            float64 `json:"price"`
	Photo            string  `json:"photo"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

// ProductFilter представляет фильтры для поиска продуктов
type ProductFilter struct {
	ID   *int
	Name *string
}

