package main

import "fmt"

func f(flag bool)*int{
	a:=1
	b:=2
	if flag{
		return &a
	}
	return &b
}

func main(){
	a := f(true)
	fmt.Println(*a)
	_ = *a
}

/*
 ╰──➤ $ go build -gcflags="-m" temp/1/main.go
# command-line-arguments
temp/1/main.go:5:6: can inline f

temp/1/main.go:15:8: inlining call to f
内联不等于不逃逸，这个很有意思，哪怕最后编译的时候将函数都插到main里面了还是要将这个要返回的值放到堆上

temp/1/main.go:16:13: inlining call to fmt.Println

temp/1/main.go:6:2: moved to heap: a
temp/1/main.go:7:2: moved to heap: b
编译器做的是静态分析：
它只需要确认一件事：
a 的地址 可能被返回
b 的地址 可能被返回

temp/1/main.go:16:13: ... argument does not escape
这是在说 fmt.Println(...) 的某些参数在内部分析后，不需要“长期存活到堆上”，可以在更短生命周期内处理

temp/1/main.go:16:14: *a escapes to heap
*/