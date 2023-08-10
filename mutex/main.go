package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	Owner       string
	Balance     int
	accountLock sync.Mutex // only accessible by BankAccount methods
}

func (account *BankAccount) Withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer account.accountLock.Unlock()

	account.accountLock.Lock()
	account.Balance -= amount
	fmt.Println("Balance: ", account.Balance)
}

func main() {
	var wg sync.WaitGroup

	account := BankAccount{Owner: "John", Balance: 100}

	for i := 0; i < 12; i++ {
		wg.Add(1)
		go account.Withdraw(10, &wg)
	}

	wg.Wait()
}
