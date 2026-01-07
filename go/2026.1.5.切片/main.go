package main

import "fmt"

func f(s []int) {
	s = append(s, 100)
}

func main() {

	a := make([]int, 0, 1)
	// name := make ([]类型 , len ，cap)
	// len 表示当前能用的元素个数
	// cap 表示底层数组还能装多少
	fmt.Println("a初始", a)
	f(a)
	fmt.Println("f函数为未返回,现在还是旧a", a)

	b := make([]int, 1, 2)
	// 上面是 0,1，无数据，如果去定义a[0]=10,相当于“往一个长度为 0 的容器的第 0 格塞东西”，必然越界，
	// 运行时会 panic（常见报错：index out of range）
	// 但是现在是 1,2 ,所以可以向第一个位置写，之前的 0,1 需要去进行 append
	b[0] = 10
	fmt.Println(b)

	c := make([]int, 6)
	fmt.Println(len(c))

	d := make([]int, 10)
	d[0] = 1
	d[1] = 2
	d[2] = 3
	d[3] = 4
	d[4] = 12
	for _, note := range d {
		fmt.Print(note, " ")
	}

}

// cap:
// 从 slice 起点到其底层数组结尾之间的元素个数

// len:
// slice 的长度（len）是它当前包含的元素个数
