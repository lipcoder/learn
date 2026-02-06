package main

import (
	"fmt"
	"log"
)

func Socialize() error {
	defer fmt.Println("goodbye")
	fmt.Println("hello")
	if 2>1 {
		return fmt.Errorf("i don't want to talk")
	}
	fmt.Println("nice ")
	return nil
}

func main() {
	defer fmt.Println("good")
	err := Socialize()
	if err != nil {
		log.Fatal(err)
		// 因为当前的log.Fatal(err)实际上是执行了一个print函数将报错打印出来，
		// 然后就执行了os.Exit(0),相当于强制关机之前注册的defer也不会执行
	}
}
