package main

import (
	"fmt"
	"os"
)

func test(args ...interface{}) {
    fmt.Println(args)
}

func main(){
	fmt.Fprintf(os.Stdout,"heello")
	test(1, "abc", true)
}