package pkg

var balance = 0
var chanQuery = make(chan int)
var chanDeposit = make(chan int)
var chanWithdraw = make(chan int)
var chanWithdraw2 = make(chan bool)

//银行演示并行goroutine访问共享变量的方法。

//存钱
func Deposit(amount int) {
	chanDeposit <- amount
}

//查询
func Balance() int {
	return <-chanQuery
}

//取钱
func Withdraw(amount int) bool {
	chanWithdraw <- amount
	return <-chanWithdraw2
}

func background() {
	for {
		select {
		case chanQuery <- balance:
		case amount := <-chanDeposit:
			balance += amount
		case amount := <-chanWithdraw:
			if balance > amount {
				balance -= amount
				chanWithdraw2 <- true
			} else {
				chanWithdraw2 <- false
			}
		}
	}
}

func init() {
	go background()
}
