package main

import "fmt"

/*
编写程序，在终端输出九九乘法表
*/
func main() {
	for i := 1; i <= 9; i++ {
		var s int
		for j := 1; j <= i; j++ {
			s = i * j
			fmt.Printf("%d*%d=%d	", j, i, s)
		}
		fmt.Printf("\n")
	}
}
