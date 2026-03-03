package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{
			Title:  "Casablanca",
			Year:   1942,
			Color:  false, // 因为 omitempty，这个字段会被省略
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
		},
		{
			Title:  "Cool Hand Luke",
			Year:   1967,
			Color:  true, // true 会输出
			Actors: []string{"Paul Newman"},
		},
		{
			Title:  "Bullitt",
			Year:   1968,
			Color:  true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"},
		},
	}

	// Marshal：Go 值 -> JSON
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %v", err)
	}
	fmt.Println("=== json.Marshal 输出===")
	fmt.Printf("%s\n\n", data)
	// fmt.Println(data) //不知道输出的是什么东西

	// 更好看：MarshalIndent（带缩进）
	data2, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling (indent) failed: %v", err)
	}
	fmt.Println("=== json.MarshalIndent 输出===")
	fmt.Println(string(data2))

	// （可选）再演示一下 Unmarshal：JSON -> Go 值
	var decoded []Movie
	if err := json.Unmarshal(data2, &decoded); err != nil {
		log.Fatalf("JSON unmarshaling failed: %v", err)
	}
	fmt.Println("\n=== 反序列化回 Go 后（decoded）===")
	for _, m := range decoded {
		fmt.Printf(" Title=%q Year=%d Color=%v Actors=%v\n", m.Title, m.Year, m.Color, m.Actors)
	}
}
