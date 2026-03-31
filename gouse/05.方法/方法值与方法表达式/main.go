package main

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) Add(f Point) Point {
	return Point{p.X + f.X, p.Y + f.Y}
}

func main() {

	abc := Point{1, 2}
	def := Point{3, 4}
	ghj := Point{4, 7}

	f1 := abc.Add
	fmt.Println("方法值:")
	fmt.Println(f1(def))

	// 没绑定任何对象,
	// 相当于func(p Point, q Point) Point {return p.Add(q)}
	f2 := Point.Add
	fmt.Println("方法表达式:")
	fmt.Println(f2(abc, ghj))
	fmt.Println(f2(ghj, def))
	fmt.Println(f2(def, abc))

	// 方法值会复制
	klm := Point{10, 20}
	f3 := klm.Add

	klm.X = 100

	fmt.Println("方法值是否捕获旧值:")
	fmt.Println(f3(Point{1, 1})) // 仍然使用旧的 klm（10,20）
}
