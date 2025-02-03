package service

import (
	"bank/models"
	"bank/repository"
	"context"
)

type ITransactionService interface {
	TransferMoney(ctx context.Context, transferMoney *models.TransferMoney) (string, error)
    Replenishment(ctx context.Context, replenishment *models.Replenishment) (string, error)
	GetLatestTransactions(ctx context.Context, id int) ([]models.Transaction, error)
}

type TransactionService struct {
	repo repository.IRepository
}

func NewTransactionService(repo repository.IRepository) ITransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) TransferMoney(ctx context.Context, transferMoney *models.TransferMoney) (string, error) {
	return s.repo.TransferMoney(ctx, transferMoney)
}

func (s *TransactionService) Replenishment(ctx context.Context, replenishment *models.Replenishment) (string, error) {
	return s.repo.Replenishment(ctx, replenishment)
}

func (s *TransactionService) GetLatestTransactions(ctx context.Context, id int) ([]models.Transaction, error) {
	return s.repo.GetLatestTransactions(ctx, id)
}
