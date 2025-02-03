package repository

import (
	"bank/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type IRepository interface {
	TransferMoney(ctx context.Context, transferMoney *models.TransferMoney) (string, error)
	Replenishment(ctx context.Context, replenishment *models.Replenishment) (string, error)
	GetLatestTransactions(ctx context.Context, id int) ([]models.Transaction, error)
}

type Repository struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) IRepository {
	return &Repository{conn: conn}
}

func (r *Repository) TransferMoney(ctx context.Context, transferMoney *models.TransferMoney) (string, error) {
	var result string

	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)
	query := `SELECT transfer_money($1, $2, $3);`
	err = tx.QueryRow(ctx, query, transferMoney.From_user_id, transferMoney.To_user_id, transferMoney.Amount).Scan(&result)
	if err != nil {
		return "", fmt.Errorf("transfer money failed: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return "", fmt.Errorf("transaction commit failed: %w", err)
	}

	return result, nil
}

func (r *Repository) Replenishment(ctx context.Context, replenishment *models.Replenishment) (string, error) {
	var result string

	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)
	query := `SELECT replenishment($1, $2);`
	err = tx.QueryRow(ctx, query, replenishment.UserID, replenishment.Amount).Scan(&result)
	if err != nil {
		return "", fmt.Errorf("replenishment failed: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return "", fmt.Errorf("transaction commit failed: %w", err)
	}

	return result, nil
}

func (r *Repository) GetLatestTransactions(ctx context.Context, id int) ([]models.Transaction, error) {
	query := `
		SELECT u.name, u.surname, t.amount, t.transaction_date FROM Transactions as t
		LEFT JOIN Users as u ON u.id = t.to_user_id OR u.id = t.from_user_id
		WHERE u.id = $1
		ORDER BY t.transaction_date DESC LIMIT 10
	`
	rows, err := r.conn.Query(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("error while transactions: %w", err)
	}
	defer rows.Close()

	return pgx.CollectRows(rows, func(row pgx.CollectableRow) (models.Transaction, error) {
		var t models.Transaction
		err := row.Scan(&t.Name, &t.Surname, &t.Amount, &t.TransactionDate)
		return t, err
	})
}
