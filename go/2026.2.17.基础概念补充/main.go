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

// 头一次看到这个赋值可以直接这样，之前看见这个多重赋值以为只有函数返回值用
// 当然这个在函数返回值主要是返回那个比尔值
func D() {
	a := 10
	b := 12
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}

// 代码块的范围很有意思
func E() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
}

// 位移运算
func F() {
	a := 52
	b := 4
	c := a << b
	fmt.Println(c)
}

func G() {
	s := "hello world "
	fmt.Println(s, len(s))
	fmt.Println(s[:5])
	fmt.Println(s[:7])
	fmt.Println(s[2:])

}

// 复数
func H() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x * y)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
}

func main() {
	H()
}
