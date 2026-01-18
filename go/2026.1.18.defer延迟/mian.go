package main

import (
	"fmt"
	"log"
)

func Socialize()error{
	defer fmt.Println("goodbye")
	fmt.Println("hello")
	return fmt.Errorf("i don't want to talk")
	fmt.Println("nice ")
	return nil
}

func main(){
	err := Socialize()
	if err != nil {
		log.Fatal(err)
	}
}