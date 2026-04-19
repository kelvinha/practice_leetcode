package interview_challenge

import (
	"fmt"
	"sync"
	"testing"
)

type BankAccount struct {
	id     int
	amount int
}

// setup account customer
var (
	customerA = BankAccount{id: 1, amount: 500000}
	mutex     sync.Mutex
	once      sync.Once
)

func TestRaceConditionSaldo(t *testing.T) {
	totalReq := 100
	wg := sync.WaitGroup{}

	for i := 0; i < totalReq; i++ {
		//once.Do(func() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			RaceConditionSaldo(&customerA, 100000)
		}()
		//})
	}

	wg.Wait()
}

func RaceConditionSaldo(customer *BankAccount, withDraw int) {
	mutex.Lock()
	if customer.amount < withDraw {
		fmt.Println("Saldo tidak cukup")
		return
	}
	customer.amount = customer.amount - withDraw
	fmt.Println("saldo sekarang", customer.amount)
	mutex.Unlock()
}
