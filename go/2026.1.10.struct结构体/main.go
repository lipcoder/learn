package main

import "fmt"

// 定义一‘种’结构体
type part struct {
	description string
	count       int
}

// 定义一’个‘结构体
func struct1() {
	var mystruct struct {
		name   string
		rate   float64
		active bool
	}
	mystruct.name = "lip"
	fmt.Println("name:", mystruct.name)
}

// 使用结构体并与structinfo传递返回结构体
func struct2() {
	var bolts part
	bolts.description = "hex"
	bolts.count = 24
	fmt.Println("b1:", bolts.count, bolts.description)

	boltscopy := structinfo(bolts)
	fmt.Println("b2:", bolts.count, bolts.description)
	fmt.Println("bc1", boltscopy.count, boltscopy.description)
}
func structinfo(p part) part {
	p.count=12
	fmt.Println("p1:", p.count, p.description)
	return p
}

// 使用指针修改struct
func struct3() {
	var s part
	struct3use(&s)
	fmt.Println(s.count)
}
func struct3use(s *part) {
	s.count = 32
	// s 是 *subscriber，但你写 s.rate = ... 不需要 (*s).rate
	// 因为 Go 对 指针访问字段 做了语法糖：s.rate 会自动解引用
}

func main() {
	struct1()
	struct2()

	struct3()
	// 类型(part)和变量(某个 part 实例)是不一样的，在3里面修改的这是s(part 类型)的count值，不是结构体的

}
