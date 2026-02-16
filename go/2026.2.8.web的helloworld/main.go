package main

import (
	"fmt"
	"log"
	"net/http"
)

func viewHandler(write http.ResponseWriter, request *http.Request) {
	message := []byte("hello world")
	_, err := write.Write(message)
	if err != nil {
		log.Fatal(err)
	}	
}

func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}

func d(writer http.ResponseWriter, request *http.Request) {
	write(writer, "z")
}
func e(writer http.ResponseWriter, request *http.Request) {
	write(writer, "x")
}
func f(writer http.ResponseWriter, request *http.Request) {
	write(writer, "y")
}

func main() {
	http.HandleFunc("/hello", viewHandler)
	http.HandleFunc("/a", f)
	http.HandleFunc("/b", d)
	http.HandleFunc("/c", e)
	fmt.Println("http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
