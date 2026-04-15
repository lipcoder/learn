/*
// 你平时写 Go 代码，处理值通常靠两种办法：
// 接口：大家都实现同一个方法，比如 String() string
// 类型分支 / 类型断言：switch x := x.(type)，分别处理 string、int、bool 这些已知类型

// 这一节先举了一个类似 fmt.Fprintf / fmt.Sprint 的例子：想写一个 Sprint(x interface{}) string，接收任意值并把它格式化成字符串。最开始可以先判断它有没有实现 String()，再分别处理 string、int、bool 等基础类型。

// 问题马上就来了：
// 基础类型还能一个个写分支。
// 但像 []float64、map[string][]string 这种组合类型会越来越多，几乎写不完。
// 更麻烦的是像 url.Values 这种具名类型。虽然它底层可能就是 map[string][]string，但它和底层类型不是同一个类型，type switch 也不会自动把它当成那个底层类型来匹配。要支持它，你又得专门知道并写出这个类型。
*/
package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {

	fmt.Println("===reflect.TypeOf()")
	a := reflect.TypeOf(3)
	fmt.Println(a.String()) //t.String返回 string，也就是这个类型的字符串表示
	fmt.Println(a)

	// reflect.TypeOf 返回的是一个reflect.Type动态类型的接口值，它总是返回具体的类型
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // 因此，下面的代码将打印 "*os.File" 而不是 "io.Writer"

	fmt.Println("===reflect.ValueOf()")
	b := reflect.ValueOf(3) // a reflect.Value
	fmt.Println(b)          // "3"
	fmt.Printf("%v\n", b)   // "3"
	fmt.Println(b.String()) // NOTE: "<int Value>" 除非这个 Value 持有的是字符串值，否则 v.String() 通常只会给你一个类似类型描述的东西，而不是实际数值

	fmt.Println("===reflect.ValueOf(x).Type()")
	b1 := reflect.ValueOf("fsdfs")
	t1 := b1.Type() //这个就相当于先拿值再用值拿类型,下面那个写法拆开的样子
	fmt.Println(t1)
	fmt.Println(reflect.ValueOf("fsdfs").Type())
	fmt.Println(b1)
	fmt.Println(b1.String())

	fmt.Println("===v.Interface()")
	c := reflect.ValueOf(3)
	x := c.Interface()
	i := x.(int)
	fmt.Println(i) //可以看为reflect.ValueOf 的“逆操作”
	fmt.Println(reflect.ValueOf(3).Interface().(int))

	fmt.Println("===v.Kind()")
	var d io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(d).Kind())
	// Type()：是 int、*os.File、main.Person
	// Kind()：是 Int、Ptr、Struct

	fmt.Println("===v.Elem()")
	e := 2
	v := reflect.ValueOf(&e) // 类型是 *int 的 Value
	f := v.Elem()            // 类型是 int 的 Value，对应 e 本身
	fmt.Println(f)           //Elem() 返回的还是 reflect.Value，不是普通 Go 值
	// 返回“解开一层”后的 reflect.Value
	// 如果 v 是指针，v.Elem() 就是它指向的那个值
	// 如果 v 是接口，v.Elem() 就是接口里装着的那个具体值

}

/*
| 调用                   | 返回类型            | 含义
| -------------------- | --------------- | -----------------------
| `reflect.TypeOf(x)`  | `reflect.Type`  | x 的动态类型
| `reflect.ValueOf(x)` | `reflect.Value` | x 的动态值
| `t.String()`         | `string`        | 类型的字符串表示
| `v.Type()`           | `reflect.Type`  | 这个值的具体类型
| `v.Kind()`           | `reflect.Kind`  | 这个值所属的大类
| `v.Interface()`      | `interface{}`   | 还原成普通接口值
| `v.Elem()`           | `reflect.Value` | 解引用/拆开接口后的一层值
| `v.Bool()`           | `bool`          | 取出布尔值
| `v.Int()`            | `int64`         | 取出有符号整数值
| `v.Uint()`           | `uint64`        | 取出无符号整数值
| `v.String()`         | `string`        | 取字符串值；但非字符串时常不是你想要的实际内容
*/
