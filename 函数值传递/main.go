package main

import "fmt"

func main() {
	truth := true
	truth = negate(truth)
	fmt.Println(truth)
	lie := false
	negate(lie)
	fmt.Println(lie)
}

func negate(Boolean bool) bool {
	return !Boolean
}
