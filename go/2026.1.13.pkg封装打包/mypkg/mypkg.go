package mypkg

import "fmt"

type Myinterface interface {
	MethodWithoutParameters()
	MethodWithParameters(float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MyType 111")
}

func (m MyType) MethodWithParameters(f float64) {
	fmt.Println("MyType 222", f)
}

func (m MyType) MethodWithReturnValue() string {
	return "MyType 333"
}

// 不属于接口的方法：依然允许存在
func (m MyType) MethodNotInterface() {
	fmt.Println("MyType extra method")
}

type OtherType struct {
	name string
}

func NewOtherType(name string) OtherType {
	return OtherType{name: name}
}

// 实现接口的三个方法
func (o OtherType) MethodWithoutParameters() {
	fmt.Println("OtherType 111", o.name)
}

func (o OtherType) MethodWithParameters(f float64) {
	fmt.Println("OtherType 222", o.name, f)
}

func (o OtherType) MethodWithReturnValue() string {
	return "OtherType 333: " + o.name
}
