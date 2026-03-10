package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] { //for1
		go fetch(url, ch)
	}

	for range os.Args[1:] { //for2
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

// 关于这篇代码里面讲述的这个goroutine知识点，
// 我非常感谢我的好友https://github.com/RetiredGuitar64
// 他的讲解成功板正了我的想法，我现在来讲讲我当前的看法

// main里面for1循环创建了对应数量的goroutine来拉取链接内容，但是当前main的goroutine是没有堵塞的，依旧在正常运行
// 所以这些goreoutine就打包成为了GMP模型里面的g，for2创建一个channel需要接收，
// 这个channel导致main的goroutine发生堵塞，然后一个g开始运行，运行时产生channel，传递给main的channel，
// main的一个channel被消费，所以for2会再产生一个channel等待接收，此时堵塞发生，这个循环会等之前所有的g被消耗完

// 上面这个是我粗浅的第一步认知，下面这段话是我询问完chatgpt后更正的结果

// 这段代码里，只创建了一个channel  （ch := make(chan string)）
// 第二个 for 只是连续从这个同一个 ch 上接收 N 次消息
// 每个 fetch goroutine 执行完后，向这个 ch 发送一个字符串结果
// 每次执行 go fetch(...)，runtime 会创建一个新的 goroutine（GMP调度模型里的一个 G）
// 这些 G 会被 Go 调度器安排到某个 M/P 上运行
// main 正在做一次无法立即完成的接收操作
// 被消费的是channel 里的一个值，不是 channel 本身
// main 会一直阻塞/继续接收，直到收齐所有 goroutine 的结果

// 下面是准确的讲述
// main 里第一个 for 循环通过 go fetch(url, ch) 启动了多个 goroutine
// 启动之后，main goroutine 不会因为 go 语句而阻塞，它会继续往下执行
// 然后 main 进入第二个 for 循环，这个循环会对同一个 channel ch 连续执行接收操作
// 因为 ch 是无缓冲 channel，如果此时还没有 goroutine 发送结果，main goroutine 就会阻塞在 <-ch 这里等待
// 每个 fetch goroutine 执行完成后，会向 ch 发送一个字符串结果
// 一旦某个 goroutine 发送成功，main 就会接收到这条结果，打印出来，然后继续下一次 <-ch
// 由于第二个 for 循环的接收次数和启动的 goroutine 数量相同，
// 所以 main 会一直等待，直到所有 fetch goroutine 都发送过一次结果，
// 这样就实现了“等待所有并发任务完成”的效果
