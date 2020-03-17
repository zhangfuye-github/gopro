package main

import (
	"fmt"
	"math/rand"
)

/*
冒泡排序：就是一次比较相邻的元素，把大的或者小的元素往后移动
	需要用到for循环，第一层for循环取要排序的元素，第二层for循环进行比较排序。
*/
func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
		fmt.Println(a)
	}
}
func main() {
	//var a[]int=[]int{9,8,7,6,5,4,3,2,1,11,12,13,14,15}
	var a []int
	fmt.Println("请输入要排序的个数：")
	var b int
	fmt.Scanf("%d", &b)
	a = make([]int, b)
	for i := 0; i < b; i++ {
		a[i] = rand.Intn(100)
	}
	bubbleSort(a)
	fmt.Println(a)
}
