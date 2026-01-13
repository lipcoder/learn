package mypkg

import "fmt"

type Myinterface interface {
	MethodWithoutParameters()
	MethodWithParameters(float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("111")
}

func (m MyType) MethodWithParameters(f float64) {
	fmt.Println("222", f)
}

func (m MyType) MethodWithReturnValue() string {
	return "333"
}

func (m MyType) MethodNotInterface() {
	fmt.Println("111")
}