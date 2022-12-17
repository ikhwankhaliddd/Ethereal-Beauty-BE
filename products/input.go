package products

import "project_dwi/users"

type GetProductDetailInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateProductInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	Benefits    string `json:"benefits" binding:"required"`
	User        users.User
}

type CreateProductImageInput struct {
	ProductID int  `form:"product_id" binding:"required"`
	IsPrimary bool `form:"is_primary" binding:"required"`
	User      users.User
}
