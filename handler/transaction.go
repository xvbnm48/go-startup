package handler

import (
	"go-startup/helper"
	"go-startup/transaction"
	"go-startup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// parameter uri
// tangkap parameter, mapping ke input struct
// panggil service, input struct sebagai parameter
// service , berbekal campaign id bisa panggil repository
// repo mencari data transaction suatu campaign

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransaction(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get Campaign Transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionByCampaignID(input)
	if err != nil {
		response := helper.APIResponse("Failed to get Campaign Transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return

	}

	response := helper.APIResponse("Success to get campaign Transaction", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}
