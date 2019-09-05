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
		fmt.Println("存100")
		pkg.Deposit(100)
		fmt.Println("存100,剩余", pkg.Balance())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("取100")
		ok := pkg.Withdraw(100)
		fmt.Println("取100", ok, ",剩余", pkg.Balance())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("取300")
		ok := pkg.Withdraw(300)
		fmt.Println("取300", ok, "剩余", pkg.Balance())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("存120")
		pkg.Deposit(120)
		fmt.Println("存120,剩余", pkg.Balance())
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("存180")
		pkg.Deposit(180)
		fmt.Println("存180,剩余", pkg.Balance())
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("finish")
}
