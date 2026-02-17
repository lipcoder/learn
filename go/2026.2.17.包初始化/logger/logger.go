package logger

import (
	"fmt"

	"lipcoder/config"
)

var Prefix = func() string {
	fmt.Println("logger: init Prefix (package var initializer)")
	return fmt.Sprintf("[port=%d] ", config.Port)
}()

func init() {
	fmt.Println("logger: init() (can use config.Port =", config.Port, ")")
}
// 包级变量config.Port

func Println(msg string) {
	fmt.Println(Prefix + msg)
}

