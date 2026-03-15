package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }
	// 上面这个是将非/的请求打回去
	w.Write([]byte("这是/，你可能写错了"))
}

func snippet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/snippet/" {
		w.Write([]byte("现在是/snippet/"))
		return
	}
	http.NotFound(w, r)
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dispaly a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { //r.Method表示当前请求方法,比如GET,POST,PUT,DELETE
		w.Header().Set("Allow", "POST")       //在调用w.WriteHeader()或w.Write()后再修改响应头映射就不会有改变了
		w.WriteHeader(405)                    //给响应写状态码,响应只能有效写一次，
		w.Write([]byte("Method Not Allowed")) //给响应写具体内容，要先写响应头再写body，否则头是默认的200
		return
	}
	w.Write([]byte("creat a new snippet..."))
}
func snippetCreate1(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("creat a new snippet..."))
}
func snippetCreate2(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("creat a new snippet..."))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/", snippet)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/snippet/create1", snippetCreate1)
	mux.HandleFunc("/snippet/create2", snippetCreate2)

	log.Print("serve on 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

// 路由的本质是路径到处理函数的映射

// 固定路径与子树路径
// 固定路径(fixed path) 比如/snippet/view ，特点就是必须要完全匹配
// 子树路径(subtree path) 当访问/snippet/view/ ，会去返回去找最长的或者说上一级，这个会返回/snippet/，
// 当这个都没有的时候会返回/
