package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func SaveToFile(f *os.File, s string) {
	f.Write([]byte(s))
}

func SaveToBuffer(b *bytes.Buffer, s string) {
	b.Write([]byte(s))
}

// A：不使用接口。
// 同样是“写入字符串”，但要针对不同具体类型分别写函数。
// 缺点：代码容易重复，扩展到新类型时还要继续加新函数。
func A() {
	var b bytes.Buffer
	SaveToFile(os.Stdout, "hello\n")
	SaveToBuffer(&b, "world")
	fmt.Println(b.String())
}

// Save 只依赖 io.Writer 接口
// 它不关心传进来的是文件、缓冲区还是别的东西
// 只关心这个值有没有 Write 方法
func Save(w io.Writer, s string) {
	w.Write([]byte(s))
}

// B：使用接口。
// 不再把函数参数写死成某个具体类型，而是抽象成“会写”的对象
// 好处：同一份逻辑可以复用给多种类型，减少重复代码
// 注意：底层最终调用的，仍然是具体类型自己的 Write 方法
func B() {
	var b bytes.Buffer
	Save(os.Stdout, "hello\n")
	Save(&b, "world")
	fmt.Println(b.String())
}

func main() {
	A()
	B()
}