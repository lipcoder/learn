package main

import (
	"fmt"
	"time"
)

func main() {
	C()
	time.Sleep(200 * time.Millisecond)
}

func C() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("C: recover 接住了：", r)
		} else {
			fmt.Println("C: 没有 panic 发生在 C 这条调用栈上，所以 recover 是 nil")
		}
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("C1: recover 接住了：", r)
			}
		}()

		fmt.Println("C1: goroutine: 我准备 panic 了")
		panic("panic inside goroutine")
	}()

	fmt.Println("C: goroutine 已启动")
}
