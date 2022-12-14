package products

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Product, error)
	FindByUserID(userID int) ([]Product, error)
	FindByID(ID int) (Product, error)
	Save(product Product) (Product, error)
	Update(product Product) (Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product

	err := r.db.Preload("ProductImages", "product_images.is_primary = 1").Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *repository) FindByUserID(userID int) ([]Product, error) {
	var product []Product

	err := r.db.Where("user_id", userID).Preload("ProductImages", "product_images.is_primary = 1").Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindByID(ID int) (Product, error) {
	var product Product
	err := r.db.Preload("User").Preload("ProductImages").Where("id = ?", ID).Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) Save(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Update(product Product) (Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}
