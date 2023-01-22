package users

import "time"

type User struct {
	ID             int
	Username       string
	Email          string
	Name           string
	PasswordHash   string
	Role           string
	AvatarFileName string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
