package products

import (
	"fmt"
	"github.com/gosimple/slug"
)

type Service interface {
	GetProducts(userID int) ([]Product, error)
	GetProductDetail(input GetProductDetailInput) (Product, error)
	CreateProduct(input CreateProductInput) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetProducts(userID int) ([]Product, error) {
	if userID != 0 {
		product, err := s.repository.FindByUserID(userID)
		if err != nil {
			return product, err
		}
		return product, nil
	}
	products, err := s.repository.FindAll()
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s *service) GetProductDetail(input GetProductDetailInput) (Product, error) {
	product, err := s.repository.FindByID(input.ID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *service) CreateProduct(input CreateProductInput) (Product, error) {
	product := Product{
		Name:        input.Name,
		UserID:      input.User.ID,
		Description: input.Description,
		Price:       input.Price,
		Benefits:    input.Benefits,
	}

	stringSlug := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	product.Slug = slug.Make(stringSlug)

	newProduct, err := s.repository.Save(product)
	if err != nil {
		return newProduct, err
	}
	return newProduct, nil
}
