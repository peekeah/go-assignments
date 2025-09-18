package main

import "errors"

type Account struct {
	Id      string
	Name    string
	City    string
	Age     int
	Balance float64
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func (a *Account) validateInsufficient(amount float64) error {
	if a.Balance < amount {
		return errors.New("insufficient balance")
	}

	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if err := a.validateInsufficient(amount); err != nil {
		return err
	}
	a.Balance -= amount
	return nil
}

func (a *Account) Transfer(r *Account, amount float64) error {
	if err := a.validateInsufficient(amount); err != nil {
		return err
	}
	a.Balance -= amount
	r.Balance += amount
	return nil
}
