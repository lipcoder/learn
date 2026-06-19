package main

import (
	"encoding/json"
	"fmt"
)

type Resp struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	body := []byte(`{
		"code": 0,
		"message": "ok",
		"data": {
			"id": 123,
			"name": "Tom"
		}
	}`)

	var resp Resp
	if err := json.Unmarshal(body, &resp); err != nil {
		panic(err)
	}

	fmt.Println(resp.Code)
	fmt.Println(resp.Message)
	fmt.Println(string(resp.Data))

	var user User
	if err := json.Unmarshal(resp.Data, &user); err != nil {
		panic(err)
	}

	fmt.Println(user.ID)
	fmt.Println(user.Name)
}