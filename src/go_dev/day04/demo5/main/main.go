package main

import "fmt"

/*
	数组和切片：
		1.数组：数据类型相同，长度固定的一个序列
				var  a []int  var b [9]int
			注意：长度是数组类型的一部分
		2.切片：
*/
func test(arr *[5]int) {
	arr[0] = 100 //要想改变原来的数组，需要传地址
}
func main() {
	var arr [5]int
	fmt.Println(arr)
	test(&arr)
	fmt.Println(arr)
}
