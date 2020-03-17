package main

/*
指针类型
*/
import "fmt"

func modify(a *int) {
	*a = 10
}
func main() {
	var a string
	var b int
	fmt.Println(&a) //输出地址值
	fmt.Println(b)  //调用函数之前
	modify(&b)      // 调用函数
	fmt.Println(b)  //调用函数之后
}
