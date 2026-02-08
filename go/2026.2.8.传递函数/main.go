package main

import (
	"fmt"
)

func sayHi() {
	fmt.Println("hi")
}

func sayBye() {
	fmt.Println("Bye")
}

func twice(theFunction func()) {
	theFunction()
	theFunction()
}

func A() {
	twice(sayHi)
	twice(sayBye)
}

// =========================================

func sayH() {
	fmt.Println("h")
}

func suan(a int, b int) int {
	return a / b
}

func B() {
	var g func()
	var m func(int, int) int
	g = sayH
	m = suan
	g()
	fmt.Println(m(4, 2))
}

func main() {
	A()
	B()
}

// 具有一级函数的编程语言还允许将函数作为参数传递给其他函数