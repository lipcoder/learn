package main

import "fmt"

func f(s []int) {
	s = append(s, 100)
}

func fr(s []int) []int {
	s = append(s, 100)
	return s
}

func f1() {
	s := make([]int, 0, 1) // slice := make ([]类型 , len ，cap)
	fmt.Println("a初始", s)
	f(s)
	fmt.Println("a现在", s) // f函数为未返回
}

func f2() {
	s := make([]int, 1, 2)
	s[0] = 10
	fmt.Println(s)
	// f1的切片是 0,1，无数据，如果去定义a[0]=10,相当于“往长度为0的容器的第0格塞东西”，必然越界，运行时会panic报错index out of range）
	// 但是现在是 1,2 ,所以可以向第一个位置写，之前的 0,1 可以去进行 append
	// 不直接自动append，所以哪怕底层数组是有一个位置的，但是现在的切片长度是0,不能放东西
	// len 表示当前能用的元素个数
	// cap 表示底层数组还能装多少
}

func f3() {
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
	fmt.Println(" ")
}

func f4() {
	arr := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println("array:", arr)

	e := arr[1:4]
	fmt.Println("e:", e)
	fmt.Println("len(s1):", len(e))
	fmt.Println("cap(s1):", cap(e))
}

func main() {

	// 1
	// 直接创建切片
	f1()

	// 2
	// 直接创建切片
	f2()

	// 3
	f3()

	// 4
	// 先创建数组，然后创建切片，更容易理解切片的本质
	f4()

}

// cap:
// 从 slice 起点到其底层数组结尾之间的元素个数

// len:
// slice 的长度（len）是它当前包含的元素个数

/*
	可以将切片看成这个
	type slice struct {
    ptr *T   // 指向底层数组
    len int  // 当前长度
    cap int  // 容量
	}

   	arr := [5]int{1,2,3,4,5}
   	s := arr[1:4]

   	arr: [1][2][3][4][5]
            ↑
            s.ptr
    s.len = 3
    s.cap = 4

*/
