package transactions

import "gorm.io/gorm"

type Repository interface {
	GetTransactionsByProductID(productID int) ([]Transactions, error)
	GetUserTransactions(userID int) ([]Transactions, error)
	GetByID(ID int) (Transactions, error)
	Save(transactions Transactions) (Transactions, error)
	Update(transaction Transactions) (Transactions, error)
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

func (r *repository) Save(transactions Transactions) (Transactions, error) {
	err := r.db.Create(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) Update(transactions Transactions) (Transactions, error) {
	err := r.db.Save(&transactions).Error

	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) GetByID(ID int) (Transactions, error) {
	var transaction Transactions

	err := r.db.Where("id = ?", ID).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}