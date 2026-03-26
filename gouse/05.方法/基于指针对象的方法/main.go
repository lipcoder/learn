package main

import "fmt"

type Point struct {
	X, Y float64
}

func (p *Point) scale(factor float64) {
	p.X = factor
	p.Y = factor
}

func (p Point) mm() {
	fmt.Println(1)
}

func main() {
	p := Point{1, 2}
	p.scale(2)

	r := &Point{1, 2}
	fmt.Println(r)
	r.scale(2)

	p1 := Point{1, 2}
	pptr := &p1
	pptr.scale(2)

	p2 := Point{1, 2}
	(&p2).scale(2)

	fmt.Println(r, p, p1, p2)

	p.mm()
}

// 怎么调用指针接收器方法
// 书里给了三种合法写法来调用 (*Point).ScaleBy：

// r := &Point{1, 2}
// r.ScaleBy(2)

// p := Point{1, 2}
// pptr := &p
// pptr.ScaleBy(2)

// p := Point{1, 2}
// (&p).ScaleBy(2)

// 这三种都行，因为本质上你最终都给了它一个 *Point

// 但 Go 进一步做了语法糖：如果接收器需要的是 *Point，而你手上有一个可取地址的 Point 变量，那你可以直接写：p.ScaleBy(2)
// 编译器会自动帮你变成：(&p).ScaleBy(2),也就是隐式取地址。这一点是本章最重要的调用规则之一
// 不管你的method的receiver是指针类型还是非指针类型，都是可以通过指针/非指针类型进行调用的，编译器会帮你做类型转换
