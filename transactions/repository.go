package transactions

import "gorm.io/gorm"

type Repository interface {
	GetTransactionsByProductID(productID int) ([]Transactions, error)
	GetUserTransactions(userID int) ([]Transactions, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetTransactionsByProductID(productID int) ([]Transactions, error) {
	var transactions []Transactions

	err := r.db.Preload("User").Where("product_id = ?", productID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) GetUserTransactions(userID int) ([]Transactions, error) {
	var transactions []Transactions

	err := r.db.Preload("Product.ProductImages", "product_images.is_primary = 1").Where("user_id = ?", userID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return []Transactions{}, err
	}
	return transactions, nil
}
