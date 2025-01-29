package pointerserrors

import (
	"errors"
	"fmt"
)

// var creates global variables to the package
var ErrInsufficientFunds= errors.New("Insufficient funds cannot withdraw")
type Bitcoin int

type Stringer interface {
	string() string
}

type Wallet struct {
	balance Bitcoin

}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}