package main

import (
	"fmt"
	"pkg"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		pkg.Deposit(100)
		fmt.Println("存100", pkg.Balance())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		pkg.Withdraw(100)
		fmt.Println("取100", pkg.Balance())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		pkg.Withdraw(300)
		fmt.Println("取100", pkg.Balance())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		pkg.Withdraw(200)
		fmt.Println("存100", pkg.Balance())
		wg.Done()
	}()

	wg.Wait()
}
