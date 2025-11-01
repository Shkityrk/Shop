package service

import (
	"context"
	"data-service/internal/application/dto"
	"data-service/internal/domain/entity"
	"data-service/internal/domain/repository"
	"errors"
	"time"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

// ProductService сервис для работы с продуктами
type ProductService struct {
	repo repository.ProductRepository
}

// NewProductService создает новый экземпляр ProductService
func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// CreateProduct создает новый продукт
func (s *ProductService) CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*dto.ProductResponse, error) {
	product := &entity.Product{
		Name:             req.Name,
		ShortDescription: req.ShortDescription,
		FullDescription:  req.FullDescription,
		Composition:      req.Composition,
		Weight:           req.Weight,
		Price:            req.Price,
		Photo:            req.Photo,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	createdProduct, err := s.repo.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return s.toProductResponse(createdProduct), nil
}

// GetProductByID получает продукт по ID
func (s *ProductService) GetProductByID(ctx context.Context, id int) (*dto.ProductResponse, error) {
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, ErrProductNotFound
	}

	return s.toProductResponse(product), nil
}

// GetProductByName получает продукт по названию
func (s *ProductService) GetProductByName(ctx context.Context, name string) (*dto.ProductResponse, error) {
	product, err := s.repo.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, ErrProductNotFound
	}

	return s.toProductResponse(product), nil
}

// ListProducts получает список всех продуктов
func (s *ProductService) ListProducts(ctx context.Context, filter *entity.ProductFilter) ([]*dto.ProductResponse, error) {
	products, err := s.repo.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.ProductResponse, 0, len(products))
	for _, product := range products {
		responses = append(responses, s.toProductResponse(product))
	}

	return responses, nil
}

// UpdateProduct обновляет данные продукта
func (s *ProductService) UpdateProduct(ctx context.Context, id int, req *dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	product, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, ErrProductNotFound
	}

	// Обновляем только переданные поля
	if req.Name != nil {
		product.Name = *req.Name
	}
	if req.ShortDescription != nil {
		product.ShortDescription = *req.ShortDescription
	}
	if req.FullDescription != nil {
		product.FullDescription = *req.FullDescription
	}
	if req.Composition != nil {
		product.Composition = *req.Composition
	}
	if req.Weight != nil {
		product.Weight = *req.Weight
	}
	if req.Price != nil {
		product.Price = *req.Price
	}
	if req.Photo != nil {
		product.Photo = *req.Photo
	}
	product.UpdatedAt = time.Now()

	updatedProduct, err := s.repo.Update(ctx, product)
	if err != nil {
		return nil, err
	}

	return s.toProductResponse(updatedProduct), nil
}

// DeleteProduct удаляет продукт по ID
func (s *ProductService) DeleteProduct(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}

// CheckProductExists проверяет существование продукта
func (s *ProductService) CheckProductExists(ctx context.Context, id int) (bool, error) {
	return s.repo.Exists(ctx, id)
}

// toProductResponse преобразует entity.Product в dto.ProductResponse
func (s *ProductService) toProductResponse(product *entity.Product) *dto.ProductResponse {
	return &dto.ProductResponse{
		ID:               product.ID,
		Name:             product.Name,
		ShortDescription: product.ShortDescription,
		FullDescription:  product.FullDescription,
		Composition:      product.Composition,
		Weight:           product.Weight,
		Price:            product.Price,
		Photo:            product.Photo,
		CreatedAt:        product.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        product.UpdatedAt.Format(time.RFC3339),
	}
}

