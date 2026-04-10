package main

import "fmt"

func main() {
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}

// 手动控制有多少个操作系统的线程同时执行Go的代码
// 默认的值是运行机器上的CPU的核心数
// GOMAXPROCS=2 go run main.go 
// GOMAXPROCS=1 go run main.go 