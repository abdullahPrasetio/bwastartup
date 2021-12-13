package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"fmt"
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
	fmt.Println(transactions)

	response := helper.APIResponseSuccess("Success get transaction campaign", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}
