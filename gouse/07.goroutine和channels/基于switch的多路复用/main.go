package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("Lift off!")
}

func main() {
	abort := make(chan struct{})

	// 单独启动一个 goroutine 监听用户输入
	// 用户按下回车后，关闭 abort 通道，广播通知主 goroutine 取消发射
	go func() {
		reader := bufio.NewReader(os.Stdin)
		_, _ = reader.ReadBytes('\n')
		close(abort)
	}()

	fmt.Println("Commencing countdown. Press Enter to abort.")

	// 比 time.Tick 更推荐：可以手动停止，避免无用的 ticker 一直存活
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for countdown := 3; countdown > 0; countdown-- {
		fmt.Println(countdown)

		select { //select 负责在这两个 channel 之间做多路复用：谁先到，就执行谁 这个正是本节的重点

		// 如果没有按回车，
		// abort 没关闭，所以 <-abort 一直阻塞
		// 但是每过 1 秒，ticker.C 可读一次

		// 如果按下回车
		// <-abort 立刻就永远可读了
		// 主 goroutine 下一次执行到 select 时，case <-abort: 会马上成立
		case <-ticker.C:
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}

	launch()
}

// 关闭 channel 常常就是用来做“广播通知”的。
// 一旦关闭，所有等着从它接收的人都会立刻被唤醒。
// 不是往里面发一个“终止消息”
// 而是通过 close(abort) 告诉所有接收方：
// “取消吧，结束吧。”

/*
               +----------------------+
               |        select        |
               +----------------------+
                  /                \
                 /                  \
                /                    \
               v                      v
   +------------------+    +------------------+
   |   <-ticker.C     |    |     <-abort      |
   +------------------+    +------------------+
   |   每 1 秒 ready   |    |  close 后立刻就绪 |
   |    周期性触发     |    |   之后一直可读    |
   +------------------+    +------------------+
               \                    /
                \                  /
                 \                /
                  \              /
                   \            /
                    v          v

        +--------------------------------------+
        |   谁先 ready，就先走哪个 case         |
        |   两边都 ready，则随机选一个 case     |
        +--------------------------------------+

              +-------------------+
              |      select       |
              +-------------------+
                 /             \
                /               \
               /                 \
              v                   v
   +------------------+   +----------------------+
   |   <-ticker.C     |   |       <-abort        |
   +------------------+   +----------------------+
   |   还没到 1 秒     |   |   未 close: 阻塞      |
   |   到了才 ready    |   |   已 close: 立刻 ready |
   +------------------+   +----------------------+
                                   |
                                   v
                     +---------------------------+
                     |     Launch aborted!       |
                     +---------------------------+
*/
