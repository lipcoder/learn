package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func A(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(body))
}

func main() {
	go A("https://example.com/")
	go A("https://www.lipcoder.top")
	go A("https://blog.lipcoder.top")
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
}
