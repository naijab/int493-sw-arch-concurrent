package bank

import (
	"math/rand"
	"sync"
	"testing"
)

func TestAccountWithRaceCondition(t *testing.T) {
	acc := Account{}
	acc.Deposit(1000000)

	wg := sync.WaitGroup{}
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				acc.Withdraw(1)
			}
		}()
	}
	wg.Wait() // lock and wait util 1000 goroutine finish

	balance := acc.Balance()
	if balance != 0 {
		t.Errorf("actual=%d", balance)
	}
}

func TestAccountTransfer(t *testing.T) {
	acc1 := &Account{}
	acc1.Deposit(1000000)
	acc2 := &Account{}
	acc2.Deposit(1000000)
	acc3 := &Account{}
	acc3.Deposit(1000000)
	accounts := []*Account{acc1, acc2, acc3}

	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		from := accounts[rand.Int()%3]
		to := accounts[rand.Int()%3]

		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				from.Transfer(to, 1)
			}
		}()
	}
	wg.Wait()

	balance := acc1.Balance()+acc2.Balance()+acc3.Balance()
	if balance != 3000000 {
		t.Errorf("actual=%d", balance)
	}
}