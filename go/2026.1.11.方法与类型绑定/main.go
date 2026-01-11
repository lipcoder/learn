package main

import "fmt"

type Money float64

func (m Money) ToCNY() float64 {
	return float64(m) * 7
}

func (m Money) ToJPY() float64 {
	return float64(m) * 150
}

func main() {
	var m Money = 10
	fmt.Println(m.ToCNY())
	fmt.Println(m.ToJPY())
}
