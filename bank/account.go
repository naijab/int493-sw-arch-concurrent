package bank

import "sync"

type Account struct {
	balance int
	l sync.Mutex
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) Withdraw(amount int) bool {
	a.l.Lock()
	defer a.l.Unlock() // if complete and return will call Unlock
	if a.balance >= amount {
		a.balance -= amount
		return true
	}
	return false
}

func (a *Account) Balance() int {
	return a.balance
}

