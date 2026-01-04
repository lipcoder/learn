package main

import (
	"fmt"
	"reflect"
)

func sayhi() {
	fmt.Print("hi")
}

func main() {
	amm := 10
	fmt.Println(&amm)
	var myint int
	var myintpointer *int
	myintpointer = &myint
	fmt.Println(reflect.TypeOf(&myint))
	fmt.Println(myintpointer)
	var myfloat float64 
	fmt.Println(reflect.TypeOf(&myfloat))

	sayhi()
	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 15)
	fmt.Printf("A string: %s\n", "hello")
	fmt.Printf("A boolean: %t\n", false)
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign: %%\n")
	
}
