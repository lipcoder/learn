package main

import (
	"fmt"
	"time"
)

func A(a string) int {
	for i := 0; i < 50; i++ {
		fmt.Print(a)
	}
	return 1
}

func main(){
	go A("a")
	go A("b")
	// a := go A() 这个是无效的，go语句不能使用返回值
	// 因为在尝试使用它之前，不能保证返回值已经准备好
	time.Sleep(time.Second)
}