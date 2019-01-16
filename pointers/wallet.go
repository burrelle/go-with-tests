package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin represents a number of Bitcoins
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet : balance
type Wallet struct {
	balance Bitcoin
}

// Deposit : Add Funds
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// ErrInsufficientFunds : No funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw : Remove funds or error
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance : Returns balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
