package main

import "fmt"

/*
	注意事项：函数首字母需要大写
*/
func main() {
	var a int = 4
	var b int = 3
	Switch(a, b)
	Switch1(&a, &b) //指针类型的参数，所以要传递地址值，故用&
	println(a, b)
}
func Switch(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
	fmt.Println(a, b) //期待结果3 4 实际结果：3 4
}
func Switch1(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
	return //期待结果3 4 实际结果：3 4
}
