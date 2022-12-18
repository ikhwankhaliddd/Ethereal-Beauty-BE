package products

import "strings"

type ProductFormatResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"product_name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Price       int    `json:"price"`
	UserID      int    `json:"user_id"`
	Slug        string `json:"slug"`
}

func FormatProductResponse(product Product) ProductFormatResponse {
	formatter := ProductFormatResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		ImageURL:    "",
		Price:       product.Price,
		Slug:        product.Slug,
		UserID:      product.UserID,
	}

	if len(product.ProductImages) > 0 {
		formatter.ImageURL = product.ProductImages[0].FileName
	}

	return formatter
}

func FormatProductsResponse(products []Product) []ProductFormatResponse {

	productsFormatter := []ProductFormatResponse{}

	for _, product := range products {
		productFormatter := FormatProductResponse(product)
		productsFormatter = append(productsFormatter, productFormatter)
	}

	return productsFormatter
}

type ProductDetailFormatter struct {
	ID          int                     `json:"id"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	ImageUrl    string                  `json:"image_url"`
	Price       int                     `json:"price"`
	UserCount   int                     `json:"user_count"`
	UserID      int                     `json:"user_id"`
	Slug        string                  `json:"slug"`
	Benefits    []string                `json:"benefits"`
	User        ProductUserFormatter    `json:"user"`
	Images      []ProductImageFormatter `json:"images"`
}

type ProductUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type ProductImageFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatProductDetail(product Product) ProductDetailFormatter {
	productDetailFormatter := ProductDetailFormatter{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		UserID:      product.UserID,
		UserCount:   product.UserCount,
		Slug:        product.Slug,
		ImageUrl:    "",
	}

	if len(product.ProductImages) > 0 {
		productDetailFormatter.ImageUrl = product.ProductImages[0].FileName
	}

	var benefits []string

	for _, benefit := range strings.Split(product.Benefits, ",") {
		benefits = append(benefits, strings.TrimSpace(benefit))
	}
	productDetailFormatter.Benefits = benefits

	user := product.User
	productUserFormatter := ProductUserFormatter{
		Name:     user.Name,
		ImageUrl: user.AvatarFileName,
	}
	productDetailFormatter.User = productUserFormatter

	images := []ProductImageFormatter{}
	for _, image := range product.ProductImages {
		productImageFormatter := ProductImageFormatter{
			ImageUrl: image.FileName,
		}
		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		productImageFormatter.IsPrimary = isPrimary

		images = append(images, productImageFormatter)
	}

	productDetailFormatter.Images = images

	return productDetailFormatter
}
