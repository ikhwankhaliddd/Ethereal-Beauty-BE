package products

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
