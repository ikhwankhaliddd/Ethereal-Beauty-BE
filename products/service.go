package products

import (
	"errors"
	"fmt"
	"github.com/gosimple/slug"
)

type Service interface {
	GetProducts(userID int) ([]Product, error)
	GetProductDetail(input GetProductDetailInput) (Product, error)
	CreateProduct(input CreateProductInput) (Product, error)
	UpdateProduct(inputID GetProductDetailInput, inputData CreateProductInput) (Product, error)
	SaveProductImage(input CreateProductImageInput, fileLocation string) (ProductImage, error)
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

func (s *service) UpdateProduct(inputID GetProductDetailInput, inputData CreateProductInput) (Product, error) {
	product, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return product, nil
	}
	if product.UserID != inputData.User.ID {
		return product, errors.New("not an owner of this product")
	}

	product.Name = inputData.Name
	product.Description = inputData.Description
	product.Price = inputData.Price
	product.Benefits = inputData.Benefits

	updatedProduct, err := s.repository.Update(product)
	if err != nil {
		return updatedProduct, err
	}
	return updatedProduct, nil
}

func (s *service) SaveProductImage(input CreateProductImageInput, fileLocation string) (ProductImage, error) {
	product, err := s.repository.FindByID(input.ProductID)
	if err != nil {
		return ProductImage{}, err
	}

	if product.UserID != input.User.ID {
		return ProductImage{}, errors.New("not an owner of this product")
	}
	isPrimary := 0

	if input.IsPrimary == true {
		isPrimary = 1
		_, err := s.repository.MarkAllImagesAsNonPrimary(input.ProductID)
		if err != nil {
			return ProductImage{}, err
		}
	}

	productImage := ProductImage{
		ID:       input.ProductID,
		FileName: fileLocation,
	}
	productImage.IsPrimary = isPrimary

	newProductImage, err := s.repository.CreateImage(productImage)
	if err != nil {
		return newProductImage, err
	}

	return newProductImage, nil
}
