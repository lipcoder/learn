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
	}
}
