package main

import (
	"errors"
	"fmt"
)

var InsufficientFundsError = errors.New("cannot widdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("address of balance in Desposit is", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}
