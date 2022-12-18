package transactions

import (
	"project_dwi/users"
)

type GetTransactionsByProductIDInput struct {
	ID   int `uri:"id" binding:"required"`
	User users.User
}

type CreateUserTransactionInput struct {
	Amount    int `json:"amount" binding:"required"`
	ProductID int `json:"product_id" binding:"required"`
	User      users.User
}
