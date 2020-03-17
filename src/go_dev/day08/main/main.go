package main

import (
	"fmt"
	"runtime"
)

/*
如何设置gorutine运行的核数
*/
func main() {
	numCPU := runtime.NumCPU() //记住runtime包，该包与协成程有关系
	fmt.Println(numCPU)
}
