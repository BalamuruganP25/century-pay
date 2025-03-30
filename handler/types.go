package handler

import (
	"sync"
	"time"
)

// User struct holds information about the user's name and balance
type User struct {
	Name    string
	Balance float64
	mu      sync.Mutex
}

// Transaction struct holds information about a completed transaction
type Transaction struct {
	Sender    string    `json:"sender"`
	Receiver  string    `json:"receiver"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

// Bank struct holds all user accounts and transaction history
type Bank struct {
	users        map[string]*User
	transactions []Transaction
	mu           sync.RWMutex
}

// api request
type TransferRequest struct {
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Amount   float64 `json:"amount"`
}
