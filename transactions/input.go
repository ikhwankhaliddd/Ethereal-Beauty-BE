package transactions

import "project_dwi/users"

type GetTransactionsByProductIDInput struct {
	ID   int `uri:"id" binding:"required"`
	User users.User
}
