package main

import (
	"fmt"
	"lipcoder/magazine"
)

func struct1() {
	address := magazine.Address{
		Street:     "123 Oak St",
		City:       "Omaha",
		State:      "NE",
		PostalCode: "68111",
	}
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.HomeAddress = address
	fmt.Println(subscriber.HomeAddress)
}

func struct2() {
	subscriber := magazine.Subscriber{
		Name: "Aman Singh",
		HomeAddress: magazine.Address{
			Street:     "123 Oak St",
			City:       "Omaha",
			State:      "NE",
			PostalCode: "68111",
		},
	}
	fmt.Println(subscriber)
}

func struct3() {
	subscriber := magazine.Subscriber{
		Name: "Aman Singh",
	}
	fmt.Println(subscriber)
}

func struct4() {
	subscriber := magazine.Employee{
		Name: "Aman Singh",
	}
	subscriber.City = "fdsfs"                    //可以直接少写一层
	subscriber.Address.PostalCode = "dfsfsfsfsd" //也可以写完整
	fmt.Println(subscriber)
}

// 其他写法
func struct5() {
	employee := magazine.Employee{
		Name: "Aman Singh",
		Address: magazine.Address{
			City: "dfsf",
		},
	}
	fmt.Println(employee)
}

func main() {
	struct1()
	struct2()
	struct3()
	struct4()
	struct5()
}
