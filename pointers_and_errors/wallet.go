package bitcoin

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds stores error message
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin is an int type
type Bitcoin int

// Stringer interface
type Stringer interface {
	String() string
}

// Wallet holds info of a wallet
type Wallet struct {
	balance Bitcoin
}

// Deposit accepts an Bitcoin type and deposits the amount
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance displays the amount in a wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw Bitcoin amount
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// String returns the amount of Bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
