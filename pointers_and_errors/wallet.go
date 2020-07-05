package main

// Wallet holds info of a wallet
type Wallet struct {
	balance int
}

// Deposit accepts an int and deposits the amount
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance displays the amount in a wallet
func (w *Wallet) Balance() int {
	return w.balance
}
