package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 1) 先把 os.Stdout 放进一个空接口里
	var x interface{} = os.Stdout

	// x 的静态类型是 interface{}
	// x 的动态类型其实是 *os.File
	fmt.Printf("x 的静态类型: interface{}\n")
	fmt.Printf("x 的动态类型: %T\n\n", x)

	// x 的静态类型是interface{}，而interface{}没有Write方法，无法通过编译
	// w.Write([]byte("hello from io.Writer\n\n"))

	// 2) 把 x 断言为 io.Writer
	w := x.(io.Writer)
	fmt.Printf("w 的静态类型: io.Writer\n")
	fmt.Printf("w 的动态类型: %T\n", w)

	// 通过 io.Writer 接口写数据
	w.Write([]byte("hello from io.Writer\n\n"))

	// 3) 再把 x 断言为具体类型 *os.File
	f := x.(*os.File)
	fmt.Printf("f 的静态类型: *os.File\n")
	fmt.Printf("f 的动态类型: %T\n", f)

	// 通过具体类型 *os.File 写数据
	f.Write([]byte("hello from *os.File\n\n"))

	// 4) 再断言成另一个接口类型，比如 io.ReadWriter
	// *os.File 同时实现了 Read 和 Write，所以这个断言也能成功
	rw := x.(io.ReadWriter)
	fmt.Printf("rw 的静态类型: io.ReadWriter\n")
	fmt.Printf("rw 的动态类型: %T\n", rw)

	// 用 rw 的 Write 方法
	rw.Write([]byte("hello from io.ReadWriter\n\n"))

}
