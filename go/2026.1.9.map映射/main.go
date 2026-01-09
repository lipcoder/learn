package main

import "fmt"

func makemap1() {
	var mymap map[string]int // map[类型1]类型2 类型1为元素索引格式，类型2为元素格式
	// 此时只是一个“没有指向任何底层哈希表”的引用，还没有真正的 map 存储空间
	// 在写入前必须先创建map，使用make或字面量（也会创建并初始化） 当前为使用make makemap2使用字面量创建
	mymap = make(map[string]int)
	mymap["gold"] = 1
	mymap["father"] = 2
	mymap["mother"] = 3
	fmt.Println(mymap["gold"])
	fmt.Println(mymap["mother"])
}

func makemap2() {
	mymap := map[string]int{"asc": 1, "fsd": 2, "erw": 3} //方便的创建map
	fmt.Println(mymap["asc"])
}

func makemap3() {
	mymap := map[string]int{
		"asc": 1,
		"fsd": 2,
		"erw": 3,
	}
	fmt.Println(mymap["asc"])
}

// map 取不到 key 时会返回“零值”，所以拿它当计数器特别方便
func counts() {
	counts := make(map[string]int)
	counts["a"]++
	fmt.Println(counts["a"], counts["b"], counts["c"])
	counts["a"]++
	fmt.Println(counts["a"], counts["b"], counts["c"])
	counts["c"]++
	fmt.Println(counts["a"], counts["b"], counts["c"])
}

func okstatus() {
	grades := map[string]int{"dfsfsf": 2, "fsdafa": 3}
	var value int
	var ok bool
	var name string
	name = "dfsfsf"
	value, ok = grades[name]
	fmt.Println("name,value,ok")
	fmt.Println(name, value, ok)
	name = "dfsfsdfsf"
	value, ok = grades[name]
	fmt.Println("name,value,ok")
	fmt.Println(name, value, ok)
}

func status1() {
	grades := map[string]float64{"dfsfsf": 2, "fsdafa": 3}
	name := "dfsfsf"
	if grades[name] < 60 {
		fmt.Println(name, grades[name])
	}
	name = "fsdsfs" //当前map没有这个值，所以map内为该索引值为0
	if grades[name] < 60 {
		fmt.Println(name, grades[name])
	}
}

func status2use(grades map[string]float64,name string){
	
	grade,ok := grades[name]
	if !ok {
		fmt.Printf("no grades records for %s.\n",name)
	}else if grade < 60 {
		fmt.Printf("%s is failing!\n",name)
	}else if grade >= 60{
		fmt.Printf("%s is ok\n",name)
	}
}

func status2() {
	grades := map[string]float64{"dfsfsf": 28, "fsdafa": 83}
	var name string
	name = "dsfs"
	status2use(grades,name)
	name = "dfsfsf"
	status2use(grades,name)
	name = "fsdafa"
	status2use(grades,name)
}

func main() {
	makemap1()
	makemap2()
	makemap3()
	counts()
	status1()
	okstatus()
	status2()
}
