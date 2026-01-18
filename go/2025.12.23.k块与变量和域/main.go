package main

import (
	"fmt"
)

var packageVar = "package"

func main() {
	var functionVar = "function"
	if true {
		var conditionalVar = "conditional"
		fmt.Println(packageVar)
		fmt.Println(functionVar)
		fmt.Println(conditionalVar)
	}
	fmt.Println(packageVar)
	fmt.Println(functionVar)
	// fmt.Println(conditionalVar)
}

// var packageVar = "package"
// 定义在 main() 外面，属于包级变量
// 在同一个包里（这里就是 package main）的任何函数都能访问它
// 所以你在if里面，if外面都能 fmt.Println(packageVar)

// var functionVar = "function"
// 定义在 main() 里面，属于 函数作用域

// var conditionalVar = "conditional"
// 定义在 if { ... } 这个花括号块里，属于 块级作用域。