package main

import (
	"math/rand"
	"time"
)

type TransferRequest struct {
	ToAccountID int64 `json:"toAccountID"`
	amount      int64 `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Default value for int64 is 0 in go
type Account struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	Number        int64     `json:"number"`
	Balance       int64     `json:"balance"`
	CreatedAt     time.Time `json:"created_at"`
	DeleteAccount bool      `json:"delete_account"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
	}
}
