package main

import "fmt"

func main() {
	truth := true
	truth = negate(truth)
	fmt.Println(truth)
	lie := false
	negate(lie)
	fmt.Println(lie)
}

func negate(Boolean bool) bool {
	return !Boolean
}

// 第一个的结果为false第二个也是，因为go的函数值传递均为形参也就是复制的方式
// 第二个结果是因为没有重新赋值