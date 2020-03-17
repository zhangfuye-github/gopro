package main

import (
	"fmt"
	"math/rand"
)

/*
插入排序：插入排序，就是找到合适的位置，依次和前面的元素比较，如果大于前面的元素，则交换位置
*/
func insertionsort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] > a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
		fmt.Println("第", i, "次排序", a)
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
	insertionsort(a)
	fmt.Println("排序后：", a)
}
