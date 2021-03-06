package bank

import "sync"

type Account struct {
	balance int
	l sync.Mutex
}

func (a *Account) Deposit(amount int) {
	a.l.Lock()
	defer a.l.Unlock()
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

func (a *Account) Transfer(b *Account, amount int) bool {
	a.l.Lock()
	if a.balance < amount {
		a.l.Unlock()
		return false
	}
	a.balance -= amount
	a.l.Unlock()

	b.l.Lock()
	b.balance += amount
	b.l.Unlock()
	return true
}