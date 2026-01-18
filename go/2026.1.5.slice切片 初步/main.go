package main

import "fmt"

func f(s []int, a int) []int {
	s = append(s, a)
	return s
}

func f1() {
	s := make([]int, 0, 1) // slice := make ([]类型 , len ，cap)
	fmt.Println("s初始", len(s), cap(s), &s[:1][0], s)
	f(s, 100)
	fmt.Println("s现在", len(s), cap(s), &s[:1][0], s) // f函数为未返回,所以没有变化
	s = f(s, 100)
	fmt.Println("s更新", len(s), cap(s), &s[:1][0], s) // f函数为返回,现在有变化
	s = f(s, 100)
	fmt.Println("s更新", len(s), cap(s), &s[:1][0], s)
	// 为什么不能直接 &s[0]？
	// 当 len(s) == 0 时，s[0] 越界，立刻 panic，即使 cap(s) > 0 也不行
	// 因为索引 s[i] 的合法性只看 len，不看 cap
	// s[:1] 是一次 重新切片,能安全取到“底层数组第 0 个槽位”的那个元素
}

func f2() {
	s := make([]int, 1, 2)
	fmt.Println("新的s", len(s), cap(s), &s[:1][0], s)
	s[0] = 10
	fmt.Println("s更新", len(s), cap(s), &s[:1][0], s)
	s = f(s, 11)
	fmt.Println("s更新", len(s), cap(s), &s[:1][0], s)
	// f1的切片是 0,1，无数据，如果去定义s[0]=10,相当于“往长度为0的容器的第0格塞东西”，
	// 必然越界，运行时会panic报错index out of range）
	// 但是现在是 1,2 ,所以可以向第一个位置写，之前的 0,1 可以去进行 append
	// 不直接自动append，所以哪怕底层数组是有一个位置的，但是现在的切片长度是0,不能放东西
	// len 表示当前能用的元素个数(slice 的长度（len）是它当前包含的元素个数)
	// cap 表示底层数组还能装多少(从 slice 起点到其底层数组结尾之间的元素个数)
}

func f3() {
	d := make([]int, 10)
	fmt.Println(d)
	d[0] = 3
	d[1] = 2
	d[2] = 5
	d[3] = 2
	d[4] = 2
	fmt.Println(d)
	for _, note := range d {
		fmt.Print(note, " ")
	}
	fmt.Println()
}

// 先创建数组，然后创建切片，更容易理解切片的本质
func f4() {
	arr := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println("array:", arr)

	e := arr[1:4]
	fmt.Println("切片e:", e)
	fmt.Println("len(e):", len(e))
	fmt.Println("cap(e):", cap(e))
}

func main() {
	f1()
	f2()
	f3()
	f4()
}

// 可以将切片看成这个
// type slice struct {
// ptr *T   // 指向底层数组
// len int  // 当前长度
// cap int  // 容量
// }

// arr := [5]int{1,2,3,4,5}
// s := arr[1:4]
// arr: [1][2][3][4][5]
//         ↑
//         s.ptr
// 		   s.len = 3
// 		   s.cap = 4
