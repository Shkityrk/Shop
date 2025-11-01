package dto

// CreateProductRequest DTO для создания продукта
type CreateProductRequest struct {
	Name             string  `json:"name" binding:"required"`
	ShortDescription string  `json:"short_description" binding:"required"`
	FullDescription  string  `json:"full_description" binding:"required"`
	Composition      string  `json:"composition" binding:"required"`
	Weight           float64 `json:"weight" binding:"required,gt=0"`
	Price            float64 `json:"price" binding:"required,gt=0"`
	Photo            string  `json:"photo" binding:"required"`
}

// UpdateProductRequest DTO для обновления продукта
type UpdateProductRequest struct {
	Name             *string  `json:"name,omitempty"`
	ShortDescription *string  `json:"short_description,omitempty"`
	FullDescription  *string  `json:"full_description,omitempty"`
	Composition      *string  `json:"composition,omitempty"`
	Weight           *float64 `json:"weight,omitempty"`
	Price            *float64 `json:"price,omitempty"`
	Photo            *string  `json:"photo,omitempty"`
}

// ProductResponse DTO для ответа с данными продукта
type ProductResponse struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	ShortDescription string  `json:"short_description"`
	FullDescription  string  `json:"full_description"`
	Composition      string  `json:"composition"`
	Weight           float64 `json:"weight"`
	Price            float64 `json:"price"`
	Photo            string  `json:"photo"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}

// CheckProductExistsResponse DTO для ответа проверки существования продукта
type CheckProductExistsResponse struct {
	Exists bool `json:"exists"`
}

