package main

import "fmt"

// 使用make创建map
func makemap1() {
	fmt.Println("===创建map的两种方法")
	var mymap map[string]int
	// map[类型1]类型2 类型1为元素索引格式，类型2为元素格式
	// 此时只是一个“没有指向任何底层哈希表”的引用，还没有真正的 map 存储空间
	mymap = make(map[string]int, 10)
	mymap["gold"] = 1
	mymap["father"] = 2
	mymap["mother"] = 3
	fmt.Print(mymap["gold"], " ")
	fmt.Println(mymap["mother"], " ")
}

// 方便的创建map的两种写法（使用字面量创建）
func makemap2() {
	var mymap = make(map[string]int)
	mymap["gold"] = 1
	mymap["father"] = 2
	mymap["mother"] = 3
	mymap1 := map[string]int{"asc": 1, "fsd": 2, "erw": 3}
	mymap2 := map[string]int{
		"asc": 1,
		"fsd": 2,
		"erw": 3,
	}
	fmt.Print(mymap["asc"], " ")
	fmt.Print(mymap1["asc"], " ")
	fmt.Println(mymap2["asc"])
}

// map 取不到 key 时会返回“零值”，所以拿它当计数器特别方便
func counts() {
	fmt.Println("===以map空值为0的来当作计数器")
	counts := make(map[string]int)
	counts["a"]++
	fmt.Println(counts["a"], counts["b"], counts["c"])
	counts["a"]++
	fmt.Println(counts["a"], counts["b"], counts["c"])
	counts["c"]++
	fmt.Println(counts["a"], counts["b"], counts["c"])
}

func okstatus() {
	fmt.Println("===社区约定俗成的ok判断法")
	grades := map[string]int{"abcd": 2, "aria": 3}
	var value int
	var ok bool
	var name string
	name = "abcd"
	value, ok = grades[name]
	fmt.Println("name:", name, "value:", value, "ok:", ok)
	value, ok = grades["abcd"]
	fmt.Println("name:", name, "value:", value, "ok:", ok)
	value, ok = grades["abce"]
	fmt.Println("name:", name, "value:", value, "ok:", ok)
}

func status1() {
	fmt.Println("===判断器实践")
	grades := map[string]float64{"adcd": 25, "abcd": 39}
	name := "adcd"
	if grades[name] < 60 {
		fmt.Println(name, grades[name])
	}
	name = "aecd" //当前map没有这个值，所以map内为该索引值为0
	if grades[name] < 60 {
		fmt.Println(name, grades[name])
	}
}

func status2use(grades map[string]float64, name string) {
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("no grades records for %s.\n", name)
	} else if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	} else if grade >= 60 {
		fmt.Printf("%s is ok\n", name)
	}
}

func status2() {
	fmt.Println("map的传递")
	grades := map[string]float64{"adcd": 28, "acdd": 83}
	var name string
	name = "addd"
	status2use(grades, name)
	name = "adcd"
	status2use(grades, name)
	name = "acdd"
	status2use(grades, name)
}

func main() {
	makemap1()
	makemap2()
	counts()
	okstatus()
	status1()
	status2()
}
