package handler

import (
	"fmt"
	"time"
)

// NewBank initializes a Bank with predefined users and balances
func NewBank() *Bank {
	return &Bank{
		users: map[string]*User{
			"Mark": {Name: "Mark", Balance: 100},
			"Jane": {Name: "Jane", Balance: 50},
			"Adam": {Name: "Adam", Balance: 0},
		},
		transactions: []Transaction{},
	}
}

func (b *Bank) TransferMoney(sender, receiver string, amount float64) error {

	b.mu.Lock() // Lock the bank to modify user accounts (write lock)
	defer b.mu.Unlock()

	// Check if both users exist
	senderUser, senderExists := b.users[sender]
	receiverUser, receiverExists := b.users[receiver]

	if !senderExists {
		return fmt.Errorf("sender user '%s' does not exist", sender)
	}
	if !receiverExists {
		return fmt.Errorf("receiver user '%s' does not exist", receiver)
	}

	// Prevent sending money to oneself
	if sender == receiver {
		return fmt.Errorf("cannot transfer money to yourself")
	}

	// Lock sender and receiver to prevent race conditions during the transfer
	senderUser.mu.Lock()
	receiverUser.mu.Lock()
	defer senderUser.mu.Unlock()
	defer receiverUser.mu.Unlock()

	// Check for sufficient funds
	if senderUser.Balance < amount {
		return fmt.Errorf("insufficient funds for user '%s'", sender)
	}
	// Perform the transfer
	senderUser.Balance -= amount
	receiverUser.Balance += amount

	// Log the transaction
	b.transactions = append(b.transactions, Transaction{
		Sender:    sender,
		Receiver:  receiver,
		Amount:    amount,
		Timestamp: time.Now(),
	})

	return nil

}

// AddUser adds a new user to the bank system
func (b *Bank) AddUser(name string, balance float64) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Check if user already exists
	if _, exists := b.users[name]; exists {
		return fmt.Errorf("user %s already exists", name)
	}

	// Add new user
	b.users[name] = &User{
		Name:    name,
		Balance: balance,
	}
	return nil
}
