package main

import (
	"fmt"
	_ "go_dev/day02/demo3/add" //只初始化，不引用，涉及关键字_
)

func main() {
	fmt.Println("main")

}
func init() {
	fmt.Println("init")
}
