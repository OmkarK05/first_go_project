// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"time"
)

type Account struct {
	ID        int32
	Owner     string
	Balance   int64
	Currency  string
	CreatedAt time.Time
}

type Entry struct {
	ID        int32
	AccountID int64
	// It can be negative or positive
	Amount    int64
	CreatedAt time.Time
}

type Transfer struct {
	ID            int32
	FromAccountID int64
	ToAccountID   int64
	// It should be positive
	Amount    int64
	CreatedAt time.Time
}
