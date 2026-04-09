package main

import (
	"flag"
	"fmt"
	"time"
)

var verbose = flag.Bool("v", false, "显示进度")

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	flag.Parse()

	limit := 30000000

	var tick <-chan time.Time
	var ticker *time.Ticker
	if *verbose {
		ticker = time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		tick = ticker.C
	}

	count := 0

	for i := 2; i <= limit; i++ {
		select {
		case <-tick:
			fmt.Printf("当前检查到 %d\n", i)
		default:
		}

		if isPrime(i) {
			count++
		}
	}

	fmt.Printf("1 到 %d 之间共有 %d 个质数\n", limit, count)
}