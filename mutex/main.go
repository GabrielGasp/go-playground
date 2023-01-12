package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	Owner       string
	Balance     int
	accountLock sync.Mutex // only accessible by BankAccount methods
}

func (account *BankAccount) Withdraw(amount int) {
	account.accountLock.Lock()
	account.Balance -= amount
	fmt.Println("Balance: ", account.Balance)
	account.accountLock.Unlock()
}

func main() {
	account := BankAccount{Owner: "John", Balance: 100}

	for i := 0; i < 12; i++ {
		go account.Withdraw(10)
	}

	time.Sleep(1 * time.Second)
}
