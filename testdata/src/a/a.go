package a

import (
	"fmt"
)

func f() {
	fmt.Println("hello world")
	fmt := 1 // want `variable "fmt" shadows imported package`
	_ = fmt
}
