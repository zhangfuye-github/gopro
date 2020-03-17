package main

import (
	"fmt"
	"os"
)

func main() {
	var arr []int //在go中数组是值类型，所以不会存在异常
	fmt.Println(arr)
	file, e := os.OpenFile("D:\\log\\test.txt", os.O_CREATE|os.O_WRONLY, 0664)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Fprintf(file, "1234")

}
