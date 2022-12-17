package transactions

import (
	"errors"
	"project_dwi/products"
)

type Service interface {
	GetTransactionsByProductID(input GetTransactionsByProductIDInput) ([]Transactions, error)
	GetUserTransactions(userID int) ([]Transactions, error)
}

type service struct {
	repository        Repository
	productRepository products.Repository
}

func NewService(repository Repository, productRepository products.Repository) *service {
	return &service{repository, productRepository}
}

func (s *service) GetTransactionsByProductID(input GetTransactionsByProductIDInput) ([]Transactions, error) {
	product, err := s.productRepository.FindByID(input.ID)

	if err != nil {
		return []Transactions{}, err
	}

	if product.UserID != input.User.ID {
		return []Transactions{}, errors.New("not the owner of this product")
	}
	transactions, err := s.repository.GetTransactionsByProductID(input.ID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) GetUserTransactions(userID int) ([]Transactions, error) {
	transactions, err := s.repository.GetUserTransactions(userID)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
