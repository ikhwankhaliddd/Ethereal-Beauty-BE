package users

import "time"

type User struct {
	ID             int    `gorm:"id"`
	Username       string `gorm:"username"`
	Email          string `gorm:"email"`
	Name           string `gorm:"name"`
	PasswordHash   string `gorm:"password_hash"`
	Role           string
	AvatarFileName string    `gorm:"avatar_file_name"`
	CreatedAt      time.Time `gorm:"created_at:CURRENT_TIMESTAMP()"`
	UpdatedAt      time.Time `gorm:"updated_at:CURRENT_TIMESTAMP()"`
}
