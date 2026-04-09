package main

import (
	"log"
	"os"
	"sync"

	"gopl.io/ch8/thumbnail"
)

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f) // NOTE: ignoring errors
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f) // NOTE: ignoring errors
			ch <- struct{}{}
		}(f) //这个的意思是使用当前的f值
		// 回忆一下之前在5.6.1节中，匿名函数中的循环变量快照问题。
		// 上面这个单独的变量f是被所有的匿名函数值所共享，且会被连续的循环迭代所更新的。
		// 当新的goroutine开始执行字面函数时，for循环可能已经更新了f并且开始了另一轮的迭代或者（更有可能的）已经结束了整个循环，
		// 所以当这些goroutine开始读取f的值时，它们所看到的值已经是slice的最后一个元素了
		// 显式地添加这个参数，我们能够确保使用的f是当go语句执行时的“当前”那个f
	}
	// Wait for goroutines to complete.
	for range filenames {
		<-ch
	}
}

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err // NOTE: incorrect: goroutine leak!
		}
		// 当它遇到第一个非nil的error时会直接将error返回到调用方，使得没有一个goroutine去排空errors channel。
		// 这样剩下的worker goroutine在向这个channel中发送值时，都会永远地阻塞下去，并且永远都不会退出。
		// 这种情况叫做goroutine泄露（§8.4.4），可能会导致整个程序卡住或者跑出out of memory的错误。
	}

	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))

	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

// 多个 goroutine 并发干活，然后把“错误”汇总回主 goroutine。

// 也就是：
// fan-out：把任务分给多个 goroutine
// fan-in：把每个 goroutine 的 error 收回来
// 主 goroutine 等待 len(filenames) 次接收

// 但它有个问题：
// 一旦某个 err != nil
// 主 goroutine 直接 return
// 剩下的 goroutine 可能还在执行
// 它们执行完后还会 errors <- err
// 但这时已经没人接收了
// 如果 channel 无缓冲，就会卡死在发送上

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1) //计数器加一，代表启动了一个待完成的任务

		go func(f string) {
			defer wg.Done() //当任务完成的时候，把这个待完成数减回去
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait() //等待这个计数器归零
		close(sizes) //当计数器归零的时候关闭sizes这个channel
	}()
	var total int64
	//只有 sizes 被关闭时它才会停
	for size := range sizes { //原来range读取channel传来的信息是不需要<-，range已经完成了这个事情
		total += size
	}
	return total
}

//
