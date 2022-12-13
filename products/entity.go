package products

import "time"

type Product struct {
	ID            int
	UserID        int
	Name          string
	Description   string
	Price         int
	UserCount     int
	Benefits      string
	Slug          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ProductImages []ProductImage
}

type ProductImage struct {
	ID        int
	ProductID int
	FileName  string
	IsPrimary int
	CreatedAt time.Time
	UpdatedAt time.Time
}
