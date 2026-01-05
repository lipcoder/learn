package main

import "fmt"

func main(){
	notes := [7]string{"do","re","mi","fa","so","la","xi"}

	// 输出数组里面的 for循环1
	fmt.Println("len(notes)=",len(notes))
	for i:=0 ;i<len(notes);i++{
		v := notes[i]
		fmt.Println(i,v)
	}

	// 本质上就是把上面这个写法糖成下面这样，只不过是将这个控制输出范围交给了编译器
	// 输出数组里面的 for range
	for index, note := range notes{
		fmt.Println(index,note)
	}

	for _,note1 := range notes{
		fmt.Println(note1)
	}
}