package pkg

import "sync"

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

//这里的带一个缓存的空结构channel 正是 sync.Mutex的实现形式,当然实际的Mutex会更加复杂一点
var chanToken = make(chan struct{}, 1)
var mutext sync.Mutex //Mutex是不能应用到重入锁的。即同一个goroutine不能对他加锁两次，第二次会被卡住
//存钱
func Deposit2(amount int) {
	chanToken <- struct{}{}
	defer func() { <-chanToken }()
	balance += amount
}

//查询
func Balance2() int {
	chanToken <- struct{}{}
	defer func() { <-chanToken }()
	return balance
}

//取钱
func Withdraw2(amount int) bool {
	chanToken <- struct{}{}
	defer func() { <-chanToken }()

	if balance >= amount {
		balance -= amount
		return true
	} else {
		return false
	}
}

var rwMutex sync.RWMutex

func Deposit3(amount int) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	balance += amount
}

//查询
func Balance3() int {
	//查询是读操作，，只需要使用读写锁即可
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	return balance
}

//取钱
func Withdraw3(amount int) bool {
	rwMutex.Lock()
	defer rwMutex.Unlock()

	if balance >= amount {
		balance -= amount
		return true
	} else {
		return false
	}
}

func background() {
	for {
		select {
		case chanQuery <- balance:
		case amount := <-chanDeposit:
			balance += amount
		case amount := <-chanWithdraw:
			if balance >= amount {
				balance -= amount
				chanWithdraw2 <- true
			} else {
				chanWithdraw2 <- false
			}
		}
	}
}

func init() {
	var once sync.Once
	once.Do(func() {
		balance = 10
	})

	go background()
}
