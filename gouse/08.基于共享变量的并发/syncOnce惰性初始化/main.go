package main

import (
	"image"
	"sync"
)

var icons map[string]image.Image

// 从磁盘读文件 / 解码图片 / 返回 image.Image
func loadIcon(_ string) image.Image {

	var a image.Image
	return a
}
func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

var mu1 sync.Mutex

func Icon1(name string) image.Image {
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

/*
// 单线程下这没问题，但并发下不安全。因为 Icon 不是一个原子操作，它包含“检查 nil → 初始化 → 赋值完成”这几个步骤。
// 多个 goroutine 同时进来时，可能都看到 icons == nil，然后重复执行 loadIcons()。

// 更关键的是，这一节强调了一个很容易忽略的点：危险不只是“初始化多次”，还可能是“另一个 goroutine 看到了一个已经非 nil、但还没初始化完的 icons”。
// 因为没有同步时，编译器和 CPU 可能对内存操作重排。书里举的例子是：先 make(map)，让 icons 变成非 nil，然后再一项项填进去。
// 这样别的 goroutine 看到 icons != nil 时，map 可能还是半成品。
*/

var mu2 sync.Mutex

func Icon2(name string) image.Image {
	mu2.Lock()
	defer mu2.Unlock()

	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

// 这样就安全了，因为同一时间只有一个 goroutine 能进入这段逻辑。
// 缺点是：初始化完成后，每次读 icons 还是得抢互斥锁，并发读性能不够好。

var mu3 sync.RWMutex

func Icon3(name string) image.Image {
	mu3.RLock()
	if icons != nil {
		icon := icons[name]
		mu3.RUnlock()
		return icon
	}
	mu3.RUnlock()

	mu3.Lock()
	if icons == nil {
		loadIcons()
	}
	icon := icons[name]
	mu3.Unlock()
	return icon
}

// 为了提高读性能，书里又给了一个 RWMutex 的版本：先拿读锁，如果已经初始化，就直接读；
// 如果还没初始化，再释放读锁、加写锁、重新检查一次，然后初始化。
// 这里为什么要“重新检查一次”？因为从你释放读锁到拿到写锁这段时间，可能别的 goroutine 已经完成初始化了

var loadIconsOnce sync.Once

func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

/*
// sync.Once 的意思就是：无论多少个 goroutine 同时调用，Do(loadIcons) 里的 loadIcons 最多只会执行一次。
// 后续调用不会再执行初始化逻辑。更重要的是，它还保证了初始化对共享内存的影响对其他 goroutine 可见，所以不会出现“别人看到半初始化状态”的问题。

// 你可以把 sync.Once 理解成：

// 它内部有一个“是否已经执行过”的标记
// 也有一个锁来保护这个过程
// 第一次调用时，执行初始化函数并把标记设为完成
// 后续调用直接跳过，但仍然保留正确的并发可见性保证
*/
