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
	mystruct.name = "lip" //直接赋值
	fmt.Println("name:", mystruct.name)
	p := part{
		count:       342,
		description: "fasdfadsf",
	} //字面量赋值
	fmt.Println(p.count, p.description)
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
	p.count = 12
	fmt.Println("p1:", p.count, p.description)
	return p
}

// 使用指针修改struct
func struct3() {
	var s part
	sc := struct3use1(&s) //var sc *part
	fmt.Println(sc.count) //直接sc.count的原因和下面这个use1里面提到原因的一样
	fmt.Println(s.count)
	struct3use2(sc) //也是那个原因,sc还是指针，所以输入时直接sc就行
	fmt.Println(sc.count)
}
func struct3use1(s *part) *part {
	s.count = 32
	// s 是 *subscriber，但你写 s.rate = ... 不需要 (*s).rate
	// 因为 go 对 指针访问字段 做了语法糖：s.rate 会自动解引用
	return s
	// 但是实际上s还是指针，如果返回的话直接写s
}
func struct3use2(s *part) {
	s.count = 42
	// 还是那个原因，所以直接s.count就行
}

func main() {
	struct1()
	struct2()
	struct3()
	// 类型(part)和变量(某个 part 实例)是不一样的，在3里面修改的这是s(part 类型)的count值，不是结构体的

}
