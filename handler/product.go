package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"project_dwi/helper"
	"project_dwi/products"
	"project_dwi/users"
	"strconv"
)

type productHandler struct {
	productService products.Service
}

func NewProductHandler(productService products.Service) *productHandler {
	return &productHandler{productService}
}

// GetProducts godoc
// @Summary      Get products
// @Description  get products data by user ID
// @Tags         Products
// @Produce      json
// @Param        user_id   query      int  true  "User ID"
// @Success      200  {object}  helper.Response{data=products.ProductDetailFormatter} "Success to get product"
// @Failure      400  {object}  helper.Response{data=interface{}} "Failed to get product"
// @Router       /products [get]
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

// GetProduct Detail godoc
// @Summary      Get product detail
// @Description  get product detail data by user ID
// @Tags         Products
// @Produce      json
// @Param        user_id   path      int  true  "User ID"
// @Success      200  {object}  helper.Response{data=products.ProductDetailFormatter} "Success to get product"
// @Failure      400  {object}  helper.Response{data=interface{}} "Failed to get product detail"
// @Router       /products/:id [get]
func (h *productHandler) GetProduct(c *gin.Context) {
	var input products.GetProductDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get product detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	productDetail, err := h.productService.GetProductDetail(input)
	if err != nil {
		response := helper.APIResponse("Failed to get product detail", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := products.FormatProductDetail(productDetail)

	response := helper.APIResponse("Success to get product detail", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// CreateProduct godoc
// @Summary      Create product
// @Description  create products
// @Tags         Products
// @Accept		 json
// @Produce      json
// @Param        payload   body      products.CreateProductInput  true  "Request Body"
// @Security	 ApiKeyAuth
// @Success      200  {object}  helper.Response{data=products.ProductDetailFormatter} "Success to create product"
// @Failure      400  {object}  helper.Response{data=interface{}} "Failed to create product"
// @Router       /products [post]
func (h *productHandler) CreateProduct(c *gin.Context) {
	var input products.CreateProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"errors": errors,
		}

		response := helper.APIResponse("Failed create product", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(users.User)

	input.User = currentUser

	newProduct, err := h.productService.CreateProduct(input)
	if err != nil {
		response := helper.APIResponse("Failed create product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := products.FormatProductResponse(newProduct)
	response := helper.APIResponse("Success to create product", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// UpdateProduct godoc
// @Summary      Update product
// @Description  update products
// @Tags         Products
// @Accept		 json
// @Produce      json
// @Param        inputID   path      int  true  "Input ID"
// @Security	 ApiKeyAuth
// @Success      200  {object}  helper.Response{data=products.ProductDetailFormatter} "Success to update product"
// @Failure      400  {object}  helper.Response{data=interface{}} "Failed to update product"
// @Router       /products/:id [put]
func (h *productHandler) UpdateProduct(c *gin.Context) {
	var inputID products.GetProductDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"errors": errors,
		}

		response := helper.APIResponse("Failed to update product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var inputData products.CreateProductInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"errors": errors,
		}

		response := helper.APIResponse("Failed update product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(users.User)
	inputData.User = currentUser

	updatedProduct, err := h.productService.UpdateProduct(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed update product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := products.FormatProductResponse(updatedProduct)
	response := helper.APIResponse("Success to update product", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

// UploadProductImage godoc
// @Summary      Upload product image
// @Description  Upload product image
// @Tags         Products
// @Accept		 json
// @Produce      json
// @Param        payload   formData      products.CreateProductImageInput  true  "Request Data"
// @Security	 ApiKeyAuth
// @Success      200  {object}  helper.Response{data=products.ProductDetailFormatter} "Success to get product"
// @Failure      400  {object}  helper.Response{data=interface{}} "Failed to get product"
// @Router       /product-images [post]
func (h *productHandler) UploadProductImage(c *gin.Context) {
	var input products.CreateProductImageInput

	err := c.ShouldBind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"errors": errors,
		}

		response := helper.APIResponse("Failed to upload product image", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload product image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(users.User)
	input.User = currentUser
	userID := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload product image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.productService.SaveProductImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload product image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Product Image successfully uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}
