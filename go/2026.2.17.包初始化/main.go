package main

import (
	"fmt"

	"lipcoder/logger"
)

func init() {
	fmt.Println("main: init()")
}

func main() {
	fmt.Println("main: main()")
	logger.Println("hello")
}
