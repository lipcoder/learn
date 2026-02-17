package main

import (
	"fmt"
	"os"
)

type MyType string
type Money float64
type Yuan float64
type Dollar float64

func NewMyType() MyType {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = ""
	}
	return MyType(s)
}

// sayHi应该叫关联到MyType这个类型上面
func (m MyType) sayHi() {
	fmt.Println("hi", m)
}

func (g Dollar)ChangeMoney() Yuan {
	return Yuan(g*7)
}

func (g Yuan)ChangeMoney() Dollar{
	return Dollar(g/7)
}

func main() {
	value := NewMyType()
	value.sayHi()

	money := Yuan(100)
	china := money.ChangeMoney()
	fmt.Println(china)
}
