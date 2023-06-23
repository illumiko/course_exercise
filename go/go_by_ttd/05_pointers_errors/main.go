package main

import (
	"errors"
	"fmt"
)

type Bitcoin float64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%#v BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposite(deposite Bitcoin) {
	w.balance += deposite
}

func (w *Wallet) Withdraw(withdraw Bitcoin) error {
	if w.balance > withdraw {
		w.balance -= withdraw
		return nil
	}
	return errors.New("Insufficient funds")

}
