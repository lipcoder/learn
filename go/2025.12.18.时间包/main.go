package main

import (
	"fmt"
	"time"
)

func main(){
	var now time.Time 
	now = time.Now()
	var year int = now.Year()
	var day int = now.Day()
	fmt.Print(year,day)
}

// time.Time 就是 time 包里定义的那个时间类型
// 这个类型下有Now这个方法将这个结构体now下的各个值初始化