package transactions

import (
	"project_dwi/users"
	"time"
)

type Transactions struct {
	ID        int
	ProductID int
	UserID    int
	Amount    int
	Status    string
	Code      string
	User      users.User
	CreatedAt time.Time
	UpdatedAt time.Time
}
