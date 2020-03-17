package main

import (
	"fmt"
	"math/rand"
)

/*
	选择排序：从前往后排，在剩余的元素当中找到最小的元素，
	然后和第一个元素比较大小，如果小于第一个元素。然后交换位置{通过下标进行比较}
*/
func selectionsort(a []int) {
	var min int
	for i := 0; i < len(a)-1; i++ {
		min = i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
		fmt.Println("第", i+1, "次排序", a)
	}
}
func main() {
	var a []int
	fmt.Println("请输入要排序的个数：")
	var b int
	fmt.Scanf("%d", &b)
	a = make([]int, b)
	for i := 0; i < b; i++ {
		a[i] = rand.Intn(100)
	}
	fmt.Println("排序前：", a)
	//var a[]int=[]int{9,8,7,6,5,4,3,2,1,11,12,13,14,15}
	selectionsort(a)
	fmt.Println("排序后：", a)
}
