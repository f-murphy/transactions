package service

import (
	"bank/mocks"
	"bank/models"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTransferMoney(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	service := NewTransactionService(mockRepo)

	transferMoney := &models.TransferMoney{
		From_user_id: 1,
		To_user_id:   2,
		Amount:       100,
	}

	mockRepo.On("TransferMoney", mock.Anything, transferMoney).Return("success", nil)

	result, err := service.TransferMoney(context.Background(), transferMoney)

	assert.NoError(t, err)
	assert.Equal(t, "success", result)
	mockRepo.AssertExpectations(t)
}

func TestReplenishment(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	service := NewTransactionService(mockRepo)

	replenishment := &models.Replenishment{
		UserID: 1,
		Amount: 100,
	}

	mockRepo.On("Replenishment", mock.Anything, replenishment).Return("success", nil)

	result, err := service.Replenishment(context.Background(), replenishment)

	assert.NoError(t, err)
	assert.Equal(t, "success", result)
	mockRepo.AssertExpectations(t)
}

func TestReplenishment_NegativeAmount(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	service := NewTransactionService(mockRepo)

	replenishment := &models.Replenishment{
		UserID: 1,
		Amount: -100,
	}

	mockRepo.On("Replenishment", mock.Anything, replenishment).Return("", errors.New("negative amount"))

	result, err := service.Replenishment(context.Background(), replenishment)

	assert.Error(t, err)
	assert.Equal(t, "", result)
	assert.Equal(t, "negative amount", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestGetLatestTransactions(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	service := NewTransactionService(mockRepo)

	userID := 1
	date1, _ := time.Parse("2006-01-02", "2023-10-01")
	date2, _ := time.Parse("2006-01-02", "2023-10-02")

	transactions := []models.Transaction{
		{Name: "Максим", Surname: "Костромов", Amount: 300, TransactionDate: date1},
		{Name: "Максим", Surname: "Костромов", Amount: 200, TransactionDate: date2},
	}
	mockRepo.On("GetLatestTransactions", mock.Anything, userID).Return(transactions, nil)

	result, err := service.GetLatestTransactions(context.Background(), userID)

	assert.NoError(t, err)
	assert.Equal(t, transactions, result)
	mockRepo.AssertExpectations(t)
}

func TestGetLatestTransactions_EmptyList(t *testing.T) {
	mockRepo := new(mocks.MockRepository)
	service := NewTransactionService(mockRepo)

	userID := 1

	mockRepo.On("GetLatestTransactions", mock.Anything, userID).Return([]models.Transaction{}, nil)

	result, err := service.GetLatestTransactions(context.Background(), userID)

	assert.NoError(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}