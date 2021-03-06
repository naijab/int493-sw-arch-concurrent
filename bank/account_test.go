package bank

import (
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