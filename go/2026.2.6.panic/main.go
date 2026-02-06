package main

import "fmt"

func main(){
	one()
}

func one(){
	defer fmt.Println("one")
	two()
}

func two(){
	defer fmt.Println("two")
	three()
}

func three(){
	defer fmt.Println("three")
	panic("too deep")
}