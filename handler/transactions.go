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
