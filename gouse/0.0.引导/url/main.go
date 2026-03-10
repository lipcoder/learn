package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if err := fetch(url); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
	}
}

func fetch(url string) error {
	url = makeURL(url)
	fmt.Println("makeURL返回的url",url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("status:", resp.Status)
	fmt.Println("statusCode:", resp.StatusCode)

	_, err = io.Copy(os.Stdout, resp.Body) //resp.Body 是正文
	if err != nil {
		return fmt.Errorf("copying %s: %w", url, err)
	}

	return nil
}

func makeURL(url string) string {
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	return url
}

// 相较于原书的写法，这这个写法叫职责分离

// resource, err := openSomething()
// if err != nil {
//     return err
// }
// defer resource.Close() 
// 后面就放心写业务逻辑，这几乎是 Go 里最经典的结构之一，谁调用资源就谁释放

//  ╭─ 18:11  aria  ~/project/learn/temp  ⭠ (bc2b695) main *%|u+1 
//  ╰──➤ $ go run main.go https://www.baidu.com
// https://https://www.baidu.com
// fetch: Get "https://https//www.baidu.com": EOF
//  ╭─ 18:14  aria  ~/project/learn/temp  ⭠ (bc2b695) main *%|u+1 
//  ╰──➤ $ go run main.go https://www.baidu.com
// http://https://www.baidu.com
// 上面这两个结果不一样的主要原因是http://https://www.baidu.com 勉强可以当一个链接去解析，
// 但是https://https://www.baidu.com不行