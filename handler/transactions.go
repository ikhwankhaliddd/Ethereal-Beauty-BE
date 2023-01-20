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

// GetTransactionsByProductID godoc
// @Summary      Get Transactions By Product ID
// @Description  Get User Transactions By Product ID
// @Tags         Transactions
// @Accept		 json
// @Produce      json
// @Param        payload   body      transactions.GetTransactionsByProductIDInput  true  "Request Body"
// @Param        product_id   path      int  true  "Path"
// @Security	 ApiAuthKey
// @Success      200  {object}  helper.Response{data=transactions.TransactionFormatter} "Success to get transactions data"
// @Failure      400  {object}  helper.Response{data=interface{}} "Failed to get transactions data"
// @Router       /products/:id/transactions [get]
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

// GetUserTransactions godoc
// @Summary      Get User Transactions
// @Description  Get User Transactions
// @Tags         Transactions
// @Produce      json
// @Security	 ApiAuthKey
// @Success      200  {object}  helper.Response{data=transactions.TransactionFormatter} "Success to get user transaction data"
// @Failure      400  {object}  helper.Response{data=interface{}} "Failed to get user transaction"
// @Router       /transactions [get]
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

// CreateUserTransaction godoc
// @Summary      Create User Transaction
// @Description  Create User Transactions
// @Tags         Transactions
// @Accept		 json
// @Produce      json
// @Param		 payload body transactions.CreateUserTransactionInput true "Request Body"
// @Security	 ApiAuthKey
// @Success      200  {object}  helper.Response{data=transactions.TransactionFormatter} "Success to Create Transaction"
// @Failure      400  {object}  helper.Response{data=interface{}} "Failed to create transaction"
// @Router       /transactions [post]
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

// GetNotification godoc
// @Summary      Get Notification
// @Description  Get Notification from Midtrans
// @Tags         Transactions
// @Accept		 json
// @Produce      json
// @Param		 payload body transactions.TransactionNotificationInput true "Request Body"
// @Security	 ApiAuthKey
// @Success      200  {object}  helper.Response{data=transactions.TransactionFormatter} "Success to get notification"
// @Failure      400  {object}  helper.Response{data=interface{}} "Failed to get notification"
// @Router       /transactions/notification [post]
func (h *transactionHandler) GetNotification(c *gin.Context) {
	var input transactions.TransactionNotificationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Failed get notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.ProcessPayment(input)
	if err != nil {
		response := helper.APIResponse("Failed get notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, input)
}
