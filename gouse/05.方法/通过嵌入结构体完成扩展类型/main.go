package main

import "fmt"

type Point struct {
	X, Y int
}

func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

type ColoredPoint struct {
	Point
	Color string
}

func main() {
	cp := ColoredPoint{
		Point: Point{1, 2},
		Color: "red",
	}

	cp.Move(10, 20) // 直接调用 Point 的方法
	fmt.Println(cp.X, cp.Y, cp.Color)
}
