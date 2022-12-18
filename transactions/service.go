package transactions

import (
	"errors"
	"fmt"
	"math/rand"
	"project_dwi/payment"
	"project_dwi/products"
	"strconv"
)

type Service interface {
	GetTransactionsByProductID(input GetTransactionsByProductIDInput) ([]Transactions, error)
	GetUserTransactions(userID int) ([]Transactions, error)
	CreateUserTransaction(input CreateUserTransactionInput) (Transactions, error)
}

type service struct {
	repository        Repository
	productRepository products.Repository
	paymentService    payment.Service
}

func NewService(repository Repository, productRepository products.Repository, paymentService payment.Service) *service {
	return &service{repository, productRepository, paymentService}
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

func (s *service) CreateUserTransaction(input CreateUserTransactionInput) (Transactions, error) {
	transaction := Transactions{
		Amount:    input.Amount,
		ProductID: input.ProductID,
		UserID:    input.User.ID,
		Status:    "pending",
		Code:      generateOrderCode(),
	}
	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		ID:     newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)

	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL

	newTransaction, err = s.repository.Update(newTransaction)

	return newTransaction, nil
}

func generateOrderCode() string {
	randomNumber := rand.Intn(999999999)
	convertedNumber := strconv.Itoa(randomNumber)
	code := fmt.Sprintf("ORDER-%s", convertedNumber)
	return code
}
