package main

import (
	"fmt"
	"lipcoder/mypkg"
)

func main(){
	var value mypkg.Myinterface
	value = mypkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameters(23)
	fmt.Println(value.MethodWithReturnValue())
}