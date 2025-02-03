package handler

import (
	"bank/models"
	"bank/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TransactionHandler struct {
	service service.ITransactionService
}

func NewTransactionHandler(s service.ITransactionService) *TransactionHandler {
	return &TransactionHandler{service: s}
}

func (h *TransactionHandler) TransferMoney(c *gin.Context) {
	var TransferMoney models.TransferMoney
	err := c.BindJSON(&TransferMoney)
	if err != nil {
		logrus.WithError(err).Error("error binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	transaction, err := h.service.TransferMoney(ctx, &TransferMoney)
	if err != nil {
		logrus.WithError(err).Error("error while transfering money")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"transaction": transaction})
}

func (h *TransactionHandler) Replenishment(c *gin.Context) {
	var replenishment models.Replenishment
	err := c.BindJSON(&replenishment)
	if err != nil {
		logrus.WithError(err).Error("error binding JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	transaction, err := h.service.Replenishment(ctx, &replenishment)
	if err != nil {
		logrus.WithError(err).Error("error replenishment")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"transaction": transaction})
}

func (h *TransactionHandler) GetLatestTransactions(c *gin.Context) {
	id := c.Param("userID")
	userID, err := strconv.Atoi(id)
	if err != nil {
		logrus.WithError(err).Error("error parsing user id")
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}
	ctx := c.Request.Context()
	transactions, err := h.service.GetLatestTransactions(ctx, userID)
	if err != nil {
		logrus.WithError(err).Error("error getting transactions")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}
