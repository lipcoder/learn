package main

import "sync"

// 1.
var (
	sema     = make(chan struct{}, 1)
	balance0 int
)

func Deposit1(amount int) {
	sema <- struct{}{} //传一个东西(当token了)，就是占用了
	balance0 += amount
	<-sema //把那个东西取出来(释放了)
}

func Balance1() int {
	sema <- struct{}{} //上面这个routine如果在用的话，现在就被堵住了
	b := balance0
	<-sema
	return b
}

// 2.
var (
	mu1      sync.Mutex
	balance1 int
)

func Deposit2(amount int) {
	mu1.Lock()
	balance1 += amount
	mu1.Unlock()
}

// 读取存款的函数
func Balance2() int {
	mu1.Lock()
	defer mu1.Unlock()
	return balance1
}

// 进行存取操作的函数
func Withdraw2(amount int) bool {
	Deposit2(-amount)
	if Balance2() < 0 {
		return false
	}
	return true
}

// 原子化
/*
// 这段代码的问题不是“没加锁”，而是：

// Deposit(-amount) 自己加锁一次
// Balance() 自己加锁一次
// Deposit(amount) 又自己加锁一次

// 也就是说，整个 Withdraw 过程被拆成了三个独立的临界区。
// 那中间就可能被别的 goroutine 插进来。

// 所以问题本质是：单个小操作是安全的，但整个业务动作不是原子的。

// 而取款这个动作，逻辑上应该是一个整体：
// 先扣钱
// 检查余额够不够
// 不够就回滚

// 这三步应该在同一次持锁期间完成，不能让别人看到中间状态。
*/

// 3.
var (
	mu2      sync.Mutex
	mu3      sync.RWMutex
	balance2 int
)

func Withdraw3(amount int) bool {
	mu1.Lock()
	defer mu1.Unlock()
	deposit(-amount)
	if balance2 < 0 {
		deposit(amount)
		return false
	}
	return true
}

// This function requires that the lock be held.
func deposit(amount int) { balance2 += amount }

/*
RWMutex 的规则到底是什么

第一条：多个读者可以同时拿读锁
Balance() 这种只读函数可以拿读锁，而不是普通写锁。这样多个 Balance() 可以并发执行

第二条：写者必须独占
像 Deposit()、Withdraw() 这种修改共享变量的函数，仍然要用普通互斥锁的那套。
写操作执行时，不允许别的读者或写者同时进入。

第三条：读写不能同时发生
RLock() 只有在没有写操作进行时才能成功共享；一旦有 goroutine 持有写锁，新的读者也得等。

所以它的本质是：
读-读可以并发，写-写不行，读-写也不行。
*/

// 读取存款的函数
func Balance3() int {
	mu3.RLock()
	defer mu3.RUnlock()
	return balance2
}

/*
// 每次一个goroutine访问bank变量时（这里只有balance余额变量），它都会调用mutex的Lock方法来获取一个互斥锁
// 如果其它的goroutine已经获得了这个锁的话，这个操作会被阻塞直到其它goroutine调用了Unlock使该锁变回可用状态
// mutex会保护共享变量。惯例来说，被mutex所保护的变量是在mutex变量声明之后立刻声明的
// 如果你的做法和惯例不符，确保在文档里对你的做法进行说明
*/
