package main

import (
	"fmt"
	"lipcoder/mypkg"
)

func main() {
	var value mypkg.Myinterface

	// 1) 用 MyType
	value = mypkg.MyType(5) //把一个具体类型为MyType、具体值为5的值，赋给了接口变量value
	// value 的静态类型是 Myinterface
	// value 的动态类型是 MyType
	// value 的动态值是 5
	value.MethodWithoutParameters()
	value.MethodWithParameters(23)
	fmt.Println(value.MethodWithReturnValue())

	fmt.Println("----")

	if v, ok := value.(mypkg.MyType); ok {
		v.MethodNotInterface() // 这个方法不在接口里,要做类型断言
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

	fmt.Println("----")
	
	var u mypkg.OtherType
	u.Age = 10
	var value1 mypkg.Myinterface = u
	value1.MethodWithoutParameters()
}

// 接口是为了把“使用方”与“实现方”隔离开（解耦）

// 不同类型但函数名一致,则运行时根据 value 里面装的真实类型，自动调用对应实现。
// 上面的叫动态分派（dynamic dispatch）
