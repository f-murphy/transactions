package mocks

import (
	"bank/models"
	"context"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) TransferMoney(ctx context.Context, transferMoney *models.TransferMoney) (string, error) {
	args := m.Called(ctx, transferMoney)
	return args.String(0), args.Error(1)
}

func (m *MockRepository) Replenishment(ctx context.Context, replenishment *models.Replenishment) (string, error) {
	args := m.Called(ctx, replenishment)
	return args.String(0), args.Error(1)
}

func (m *MockRepository) GetLatestTransactions(ctx context.Context, id int) ([]models.Transaction, error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]models.Transaction), args.Error(1)
}