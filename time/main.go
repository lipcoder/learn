package main

import (
	"fmt"
	"time"
)

func main(){
	var now time.Time = time.Now()
	var year int = now.Year()
	var day int = now.Day()
	fmt.Print(year,day)
}