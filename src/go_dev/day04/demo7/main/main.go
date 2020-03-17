package main

import "fmt"

/*
切片是一个长度可变的数组，是数组的引用的，是引用类型
用数组初始化切片时，切片的容量等于引用数组的长度,切片的初始化方法有两种：
1.make分配空间，这种情况下len=cap，
2.通过引用数组，给切片追加元素要用到append函数
*/
func main() {
	var arr [6]int = [6]int{1, 2, 3, 4, 5, 6}
	var slice []int
	/*slice=arr[:4]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice=slice[0:len(slice)-1]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice=make([]int,100)
	for i:=0;i<10 ;i++  {
		slice[i]=i
	}
	fmt.Println(len(slice))//已经分配空间，并且有默认值
	fmt.Println(cap(slice))*/
	slice = arr[0:4]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice, 1, 2, 4) //切片时按数组的长度扩容
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	//arr=append(arr,1,2,3,4,5,6,7,8,9)//数组不支持扩容
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
}
