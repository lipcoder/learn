package config

import (
	"fmt"
	"os"
)

var Port = func() int {
	fmt.Println("config: init Port (package var initializer)")
	if os.Getenv("APP_PORT") == "12" {
		return 8080
	}
	return 9090
}()

func init() {
	fmt.Println("config: init()")
}

// 包内顺序是先变量，后init

/*
	 ╭─ 16:27  aria  ~/project/learn/go/2026.2.17.包初始化  ⭠ (c83b263) main %|u=
	 ╰──➤ $ go run main.go
	config: init Port (package var initializer)
	config: init()
	logger: init Prefix (package var initializer)
	logger: init() (can use config.Port = 9090 )
	main: init()
	main: main()
	[port=9090] hello
*/
