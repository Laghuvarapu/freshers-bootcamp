package main

import (
	"fmt"
	"sync"
)

var balance int = 500
var mutex = &sync.Mutex{}

func deposit(amountDepositing int, wg *sync.WaitGroup) {
	mutex.Lock()
	balance += amountDepositing
	fmt.Printf("Deposited %d amount and new current balance %d\n", amountDepositing, balance)
	mutex.Unlock()
	wg.Done()
}

func withdrawl(amountWithdrawing int, wg *sync.WaitGroup) {
	mutex.Lock()
	if balance < amountWithdrawing {
		fmt.Println("There is no enough money to withdrawl")
	} else {
		balance -= amountWithdrawing
		fmt.Printf("Withdrawing %d amount and new current balance %d\n", amountWithdrawing, balance)
	}
	mutex.Unlock()

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go withdrawl(700, &wg)

	go deposit(200, &wg)

	wg.Wait()
}
