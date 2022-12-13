package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project_dwi/helper"
	"project_dwi/products"
	"strconv"
)

type productHandler struct {
	productService products.Service
}

func NewProductHandler(productService products.Service) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) GetProducts(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	userProducts, err := h.productService.GetProducts(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := products.FormatProductsResponse(userProducts)
	response := helper.APIResponse("Success to get products", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
