package main

import "sync"

var (
	sema     = make(chan struct{}, 1)
	balance1 int
)

func Deposit1(amount int) {
	sema <- struct{}{} //传一个东西(当token了)，就是占用了
	balance1 += amount
	<-sema //把那个东西取出来(释放了)
}

func Balance1() int {
	sema <- struct{}{} //上面这个routine如果在用的话，现在就被堵住了
	b := balance1
	<-sema
	return b
}

func Balance2() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

var (
	mu      sync.Mutex
	balance int
)

func Deposit3(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Balance3() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func Withdraw3(amount int) bool {
	Deposit3(-amount)
	if Balance3() < 0 {
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

func Withdraw4(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

// This function requires that the lock be held.
func deposit(amount int) { balance += amount }

/*
// 每次一个goroutine访问bank变量时（这里只有balance余额变量），它都会调用mutex的Lock方法来获取一个互斥锁
// 如果其它的goroutine已经获得了这个锁的话，这个操作会被阻塞直到其它goroutine调用了Unlock使该锁变回可用状态
// mutex会保护共享变量。惯例来说，被mutex所保护的变量是在mutex变量声明之后立刻声明的
// 如果你的做法和惯例不符，确保在文档里对你的做法进行说明
*/
