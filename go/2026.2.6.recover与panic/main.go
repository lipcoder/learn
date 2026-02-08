// panic发生 → 开始退栈 → 执行defer ←【recover 只能在这里生效】→ 继续退栈/或被拦住

package main

import (
	"errors"
	"fmt"
)

func A() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover接住panic,panic内容为:", r)
		}
		fmt.Println("此段无论是否panic都会执行")
	}()

	func() {
		fmt.Println("panic准备")
		panic("too deep")
	}()

	fmt.Println("此段不会执行")
}

func B() {
	defer func() {
		p := recover()
		fmt.Printf("recover得到的是: %#v ,静态类型: %T\n", p, p)

		// 先做类型断言，再用 error 的方法
		if err, ok := p.(error); ok {
			fmt.Println(err.Error())
			return
		}else {
			fmt.Println(p)
		}
		// 可以像正常if加东西，判断是哪个就走哪个，常见可以是字符串
	}()

	panic(errors.New("there is an error"))
}

func main(){
	A()
	B()
}