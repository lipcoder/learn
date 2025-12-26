package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	target := rand.Intn(100) + 1 //赋值一个随机数，来自于math下的rand包

	/*
		在较老的 Go 版本里（例如 Go 1.19 及更早），math/rand 的包级别随机数生成器如果你不手动 rand.Seed(...)，它的默认种子是固定的（等价于一直用同一个 seed，经典是 1）。
		所以程序每次启动都会从同一条伪随机序列的开头开始，第一次 rand.Intn(100)+1 永远是同一个值（书里示例是 82）。

		从较新的 Go 版本开始（Go 1.20 起），math/rand 的包级别函数会在程序启动时自动用随机种子初始化（接近“每次启动都不同”的效果），所以你不写 rand.Seed 也会得到不同的结果——这正是你截图里每次 go run 输出都不同的原因。
	*/

	fmt.Println("I’ve chosen a random number between 1 and 100")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)

	sucess := false

	for x := 10 ; x>0 ;x-- {

		fmt.Print("make a guess")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil || guess > 100 || guess < 1 {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("low")
			continue
		} else if guess > target {
			fmt.Println("hign")
			continue
		} else {
			sucess = true
			fmt.Println("good")
			break
		}
	}

	if !sucess{
		fmt.Print("sorry")
		fmt.Println(target)
	}
}
