package main

import (
	"fmt"
	"lipcoder/mypkg"
)

func main() {
	var value mypkg.Myinterface

	// 1) 用 MyType
	value = mypkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameters(23)
	fmt.Println(value.MethodWithReturnValue())

	fmt.Println("----")

	// 2) 换成 OtherType（main 的调用逻辑完全不变）
	value = mypkg.NewOtherType("abc")
	value.MethodWithoutParameters()
	value.MethodWithParameters(23)
	fmt.Println(value.MethodWithReturnValue())
}
