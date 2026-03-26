## 基于指针对象的方法

------

## 方法的接收器，可以是值，也可以是指针

```go
func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}

func main(){
    // 以下这三种都行，因为本质上你最终都给了它一个 *Point
	r := &Point{1, 2}
	fmt.Println(r)
	r.scale(2)

	p1 := Point{1, 2}
	pptr := &p1
	pptr.scale(2)

	p2 := Point{1, 2}
	(&p2).scale(2)
    // 但 Go 进一步做了语法糖：如果接收器需要的是 *Point，
    // 而你手上有一个可取地址的 Point 变量，那你可以直接写：p.ScaleBy(2)
	// 编译器会自动帮你补变成：(&p).ScaleBy(2),也就是隐式取地址，这一点是本章最重要的调用规则之一
    p := Point{1, 2}
	p.scale(2)

	fmt.Println(r, p, p1, p2)
}
```

但临时值不行。比如：

```go
Point{1, 2}.scale(2)
```

这会报错，因为字面量 `Point{1,2}` 不是一个可取地址的变量，编译器没法帮你自动补 `&`，网页给的报错就是 “can't take address of Point literal”

这个地方你可以记成一句话：

**Go 只会帮你偷偷取地址，但前提是这个东西本来就有地址**

------

## nil可以作为接收器

```go
package url

type Values map[string][]string
func (v Values) Get(key string) string {
    if vs := v[key]; len(vs) > 0 {
        return vs[0]
    }
    return ""
}

func (v Values) Add(key, value string) {
    v[key] = append(v[key], value)
}
```

#### Q ：这里的这个nil应该说是当前的这个接收器参数的类型是*IntList，但是这个接收器的值是nil是吗

### A：

你说的这个理解：

> 当前接收器参数的类型是 `*IntList`，但是这个接收器的值是 `nil`

**对，完全对**

接收器本质上也是一个参数 这个参数类型是 `*IntList` 它当前传进来的值刚好是 `nil`

如果方法一进来就先判断：,if list == nil,那就没问题。

但如果你一上来就：return list.Value,那就会因为解引用 nil 指针而 panic。

所以不是“nil 接收器天然安全”，而是：

nil 作为接收器值是允许的；是否安全取决于方法内部是否正确处理 nil。

------

## 方法调用本质上就是把接收器当第一个参数传进去

```go
package url

type Values map[string][]string
func (v Values) Get(key string) string {
    if vs := v[key]; len(vs) > 0 {
        return vs[0]
    }
    return ""
}

func (v Values) Add(key, value string) {
    v[key] = append(v[key], value)
}
```

```go
m := url.Values{"lang": {"en"}} // direct construction
m.Add("item", "1")
m.Add("item", "2")
```

#### Q ：那为什么 `Add` 能改到外面的 `m`？

这是这个地方最容易误解的点。

因为 `Values` 的底层是：

```go
type Values map[string][]string
```

而 **map 是引用语义的类型**。

所以当你写：

```go
m.Add("item", "3")
```

本质上虽然是把 `m` “按值传给了 v”，但是这个“值”本身是一个 map 描述符，它内部指向底层哈希表。

所以：

- `v` 和 `m` 是两份 map 变量
- 但它们都指向**同一份底层数据**

于是你在 `Add` 里写：

```go
v[key] = append(v[key], value)
```

改的是底层 map 数据，所以外面的 `m` 能看到变化。

这和 struct 值接收器不一样，比如：

```go
type Point struct {
    X, Y int
}

func (p Point) SetX(x int) {
    p.X = x
}
```

这里 `p` 是整个结构体的副本。
 你改 `p.X`，外面那个原始 `Point` 不变。

但是 map 不一样：

```go
type Values map[string][]string
```

map 本身不是“整张表的数据实体”，它更像一个“指向底层结构的句柄”。

所以即使 `v` 是副本，副本里装的仍然是“指向同一个底层表”的信息。