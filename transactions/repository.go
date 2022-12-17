package transactions

import "gorm.io/gorm"

type Repository interface {
	GetByProductID(productID int) ([]Transactions, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByProductID(productID int) ([]Transactions, error) {
	var transactions []Transactions

	err := r.db.Preload("User").Where("product_id = ?", productID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
