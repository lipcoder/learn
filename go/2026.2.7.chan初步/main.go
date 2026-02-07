package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func B(mychannel chan int) {
	mychannel <- 12
}

func B1() {
	mychannel := make(chan int)
	go B(mychannel)
	fmt.Println(<-mychannel)
}

func C(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}

func D(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

func E() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go C(channel1)
	go D(channel2)

	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
}

type Page struct {
	URL  string
	Size int
}

func F(url string, channel chan Page) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- Page{URL: url, Size: len(body)}
}

func G() {
	pages := make(chan Page)
	urls := []string{"https://www.lipcoder.top", "https://blog.lipcoder.top"}
	for _, url := range urls {
		go F(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Println(page.URL, page.Size)
	}
}

func main() {
	E()
	G()
}
