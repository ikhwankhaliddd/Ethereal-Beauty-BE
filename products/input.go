package products

type GetProductDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
