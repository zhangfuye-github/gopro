package main

import (
	"fmt"
	"go_dev/day01/package_demo/cacl2"
	"time"
)

func main() {
	sum := cacl2.ADD(1, 2)
	fmt.Println(sum)
	time.Sleep(10 * time.Second)
}
