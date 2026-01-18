package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Print("enter the grade:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(input,err)
}
// 这个主要的部分就是reader := bufio.NewReader(os.Stdin)
// reader是一个结构体指针，具体的流程为创造一个结构体然后返回这个结构体的指针
// 这个结构体类型有ReadString这个方法