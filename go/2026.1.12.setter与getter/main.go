package main

import (
	"fmt"
	"lipcoder/calendar"
	"log"
)

func main() {
	data := calendar.Date{}
	err := data.SetYear(23)
	if err != nil {
		log.Fatal(err)
	}
	err = data.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = data.SetDay(12)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

	fmt.Println(data.Day())
	// data.Day() = 13
}

/*
	这个程序可以防止直接在main里面直接去定义这个day，就是靠将这几个变量调整为首字母小写
	这样可以防止绕过数据的核验机制
	这样叫setter方法

	data.Day()这个叫getter方法
*/
