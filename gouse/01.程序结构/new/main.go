package main

import (
	"fmt"
)

// new创建的过程是申请一块内存初始化给这个变量
func A() {
	p := new(int)
	fmt.Println(*p)
	fmt.Printf("type: %T\n", p)
	fmt.Printf("type: %T\n", *p)
	*p = 2
	fmt.Println(*p)
}

// 只是因为我看到这个*p = true有点奇怪
func B() {
	p := new(bool)
	*p = true
	fmt.Printf("type: %T,%T\n", p, *p)
}

// 命名返回值
func C2() (r int) {
	r = 10
	return
	// 不需要写返回值或者变量都行
}

func C1() (r int) {
	r = 10
	return 12
}

func C() {
	a := C2()
	b := C1()
	fmt.Println(a, b)
}

func main() {
	A()
	B()
	C()
}

// make 与 new 的区别
// slice/map/chan 这种“引用类型”有内部指针/结构，仅仅“给它一个零值”并不能让它可用，必须 make 初始化内部结构
