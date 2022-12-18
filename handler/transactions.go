package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project_dwi/helper"
	"project_dwi/transactions"
	"project_dwi/users"
)

type transactionHandler struct {
	service transactions.Service
}

func NewTransactionsHandler(service transactions.Service) *transactionHandler {
	return &transactionHandler{service}
}
func (h *transactionHandler) GetTransactionsByProductID(c *gin.Context) {
	var input transactions.GetTransactionsByProductIDInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get product's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(users.User)
	input.User = currentUser

	allTransactions, err := h.service.GetTransactionsByProductID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get product's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := transactions.FormatProductTransactions(allTransactions)
	response := helper.APIResponse("Success to get product's transactions", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(users.User)

	userID := currentUser.ID

	allTransactions, err := h.service.GetUserTransactions(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get user's transactions", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := transactions.FormatUserTransactions(allTransactions)
	response := helper.APIResponse("Success to get user's transactions", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) CreateUserTransaction(c *gin.Context) {
	var input transactions.CreateUserTransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"errors": errors,
		}

		response := helper.APIResponse("Failed create transaction", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(users.User)
	input.User = currentUser

	newTransaction, err := h.service.CreateUserTransaction(input)
	if err != nil {
		response := helper.APIResponse("Failed create transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := transactions.FormatTransaction(newTransaction)
	response := helper.APIResponse("Success create transaction", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
