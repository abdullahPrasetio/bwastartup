package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransaction(c *gin.Context) {
	var input transaction.GetCampaignTransactionInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponseError("Failed to get campaign's transactions", http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Ambil current user
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionByCampaignID(input)
	if err != nil {
		response := helper.APIResponseError("Failed to get campaign's transactions", http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponseSuccess("Success get transaction campaign", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransaction(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	transactions, err := h.service.GetTransactionByUserID(userID)
	if err != nil {
		response := helper.APIResponseError("Failed to get users's transactions", http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponseSuccess("Success get transaction campaign", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponseError("Failed to create transaction", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newTransaction, err := h.service.CreateTransaction(input)
	if err != nil {
		response := helper.APIResponseError("Failed to create transaction", http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponseSuccess("Success to create transaction", transaction.FormatTransaction(newTransaction))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetNotification(c *gin.Context) {
	var input transaction.TransactionNotificationInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.APIResponseError("Failed to process notification", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.PaymentProcess(input)
	if err != nil {
		response := helper.APIResponseError("Failed to process notification", http.StatusBadRequest, nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	c.JSON(http.StatusOK, input)
}
