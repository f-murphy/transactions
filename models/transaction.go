package models

import "time"

type Transaction struct {
	Name            string
	Surname         string
	Amount          float32
	TransactionDate time.Time
}

type TransferMoney struct {
	From_user_id int
	To_user_id   int
	Amount       float32
}

type Replenishment struct {
	UserID int
	Amount float32
}