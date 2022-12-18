package transactions

import (
	"project_dwi/products"
	"project_dwi/users"
	"time"
)

type Transactions struct {
	ID         int
	ProductID  int
	UserID     int
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	User       users.User
	Product    products.Product
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
