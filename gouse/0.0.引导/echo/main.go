package main

import (
	"fmt"
	"os"
	"strings"
)

func A(){
	var s , sep string
	for i := 1 ; i < len(os.Args); i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(&os.Args[0])
	fmt.Println(s)
}

// os.Args[0]：程序本身的名字或路径
// os.Args[1:]：才是你真正传进去的参数

func B(args []string){
	s ,sep := "" , ""
	for _,arg := range args{
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// s := ""
// var s string
// var s = ""
// var s string = ""
// 短变量声明只能在函数里面使用，不能在包级使用
// 依赖字符串的默认初始化零值，会默认将s初始化为""
// 这个非常少，一般只有同时声明多个变量的情况下使用
// 显式标注变量的类型

func C(args []string){
	fmt.Println(strings.Join(args," "))
}

func D(args []string){
	var s , sep string
	for i := 0 ; i < len(args); i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func main(){
	args := os.Args[1:]
	fmt.Println(&os.Args[0],args)
	fmt.Println("A=======")
	A()
	fmt.Println("B=======")
	B(args)
	fmt.Println("C=======")
	C(args)
	fmt.Println("D=======")
	D(args)
	// os.Args是一个[]string，保存了程序启动时从命令行接收到的所有参数,
	// os.Args 是 os 包里的一个全局变量
}