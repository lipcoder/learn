// 声明当前程序所属的包
// main 包是 Go 程序的入口包
package main

// 导入程序中需要用到的标准库
import (
	"bufio" // 用于带缓冲地读取输入（如从终端读取一整行）
	"fmt"
	"log"       // 用于日志输出（出现错误时直接终止程序）
	"math/rand" // 用于生成随机数
	"os"        // 提供与操作系统交互的功能（如标准输入输出）
	"strconv"   // 用于字符串和基本数据类型之间的转换
	"strings"   // 用于字符串处理（如去除空格、换行符）
)

// main 函数是程序的入口，程序从这里开始执行
func main() {

	// 生成一个 1~100 之间的随机整数
	// rand.Intn(100) 会生成 0~99 的随机数，因此需要 +1
	target := rand.Intn(100) + 1

	// 提示用户游戏开始
	fmt.Println("I’ve chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	// 创建一个 Reader，用于从标准输入（键盘）读取用户输入
	// os.Stdin 表示程序的标准输入
	reader := bufio.NewReader(os.Stdin)

	// 提示用户输入猜测的数字
	fmt.Print("Make a guess: ")

	// 从输入中读取一行字符串，直到遇到换行符 '\n'
	input, err := reader.ReadString('\n')
	if err != nil {
		// 如果读取输入时发生错误，输出错误信息并终止程序
		log.Fatal(err)
	}

	// 去除输入字符串两端的空白字符（包括换行符）
	input = strings.TrimSpace(input)

	// 将用户输入的字符串转换为整数
	guess, err := strconv.Atoi(input)
	if err != nil || guess > 100 || guess < 1 {
		// 如果字符串无法转换为整数（如输入了字母），直接终止程序
		log.Fatal(err)
	}

	if guess < target {
		fmt.Println("Oops. Your guess was LOW.")
	} else if guess > target {
		fmt.Println("Oops. Your guess was HIGH.")
	} else {
		fmt.Println("Good job! You guessed it!")
	}
}
