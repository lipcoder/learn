package main

import (
	"fmt"
)

func main() {
	
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "xi"}

	// 输出数组里面的 for循环1
	fmt.Println("len(notes)=", len(notes))
	for i := 0; i < len(notes); i++ {
		v := notes[i]
		fmt.Print(i, v)
		fmt.Print(" ")
	}
	fmt.Println()

	// 本质上就是把上面这个写法糖成下面这样，只不过是将这个控制输出范围交给了编译器
	// range --> 输出右边这个数组/切片的索引和值
	for index, note := range notes {
		fmt.Print(index, note)
		fmt.Print(" ")
	}
	fmt.Println()

	for _, note1 := range notes {
		fmt.Print(note1)
		fmt.Print(" ")
	}
}
