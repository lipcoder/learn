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

	if v, ok := value.(mypkg.MyType); ok {
		v.MethodNotInterface() // 这个方法不在接口里
	}

	fmt.Println("----")

	// 2) 换成 OtherType（main 的调用逻辑完全不变）
	value = mypkg.NewOtherType("abc")
	value.MethodWithoutParameters()
	value.MethodWithParameters(23)
	fmt.Println(value.MethodWithReturnValue())

	if v, ok := value.(mypkg.MyType); ok {
		v.MethodNotInterface()
	} else {
		fmt.Println("value is not MyType, cannot call MethodNotInterface")
	}
}

// 接口是为了把“使用方”与“实现方”隔离开（解耦）

// 不同类型但函数名一致,则运行时根据 value 里面装的真实类型，自动调用对应实现。
// 上面的叫动态分派（dynamic dispatch）