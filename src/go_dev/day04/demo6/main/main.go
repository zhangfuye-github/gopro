package main

import "fmt"

/*
练习：使用非递归的方式实现斐波那契数列，打印前100个数。
数组要求必须长度固定，因此用切片，切片是长度可变的数组，是引用类型的传递
*/
func main() {
	for {
		fmt.Println("请输入一个整数：")
		var a int
		fmt.Scanf("%d", &a)
		var arr []int
		arr = make([]int, a)
		arr[0] = 1
		arr[1] = 1
		for i := 2; i < a; i++ {
			arr[i] = arr[i-1] + arr[i-2]
		}
		for i := 0; i < a; i++ {
			fmt.Println(arr[i])
		}
	}

}
