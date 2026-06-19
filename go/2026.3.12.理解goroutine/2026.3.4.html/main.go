package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type PageData struct {
	Title string
	Name  string
}

func main() {
	// 1) 定义模板内容（注意：这就是一个字符串模板）
	const tpl = 
`<h1>{{.Title}}</h1>
<p>Hello, {{.Name}}!</p>
`

	// 2) 解析模板：把字符串编译成一个 *template.Template
	t, err := template.New("hello").Parse(tpl)
	// template.New("hello") 返回一个 *Template
	// 这个类型有 Parse 方法
	// Parse(tpl) 解析模板文本，返回（解析后的模板对象，解析错误 err）
	fmt.Println(t.Name())
	if err != nil {
		log.Fatal(err)
	}

	// 3) 准备要填充的数据
	data := PageData{
		Title: "Welcome",
		Name:  "Gopher",
	}

	// 4) 执行模板：把 data 填进去，输出到 os.Stdout
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
