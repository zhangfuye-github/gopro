package main

import "fmt"

/*
	make和new的区别:new只是声明一个变量，make是给变量赋予空间,通过案例细品make和new的区别
*/
func test() {
	i := new([]int) //这是slice，值类型的变量
	fmt.Println(i)
	ints := make([]int, 10) //这是管道，引用类型的变量
	fmt.Println(ints)
	*i = make([]int, 5)
	(*i)[0] = 100
	ints[0] = 100
	fmt.Println(i)
	return
}
func main() {
	test()
}
