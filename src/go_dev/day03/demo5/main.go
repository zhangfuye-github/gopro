package main

/*
	switch 分支：
		条件分支语句和其他语言的区别：不需要break，如果需要贯穿的话可以使用fallthrough语句，
其次多个条件值可以用逗号分隔开，case 1,2,3:
*/
import "fmt"

func main() {
	var i int = 0
	switch i {
	case 0:
		fallthrough
	case 1:
		fmt.Println("i的值是1")
	case 3:
		fmt.Println("i的值是3")
	default:
		fmt.Println("i是默认值")
	}
}
