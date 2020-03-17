package main

import (
	"errors"
	"fmt"
)

/*
内置函数：无需导入包，就可以使用的函数
		1.close，主要用于channal
		2.len ，用来计算数组和字符串的长度
		3.new 为变量分配内存，分配值类型传给变量的是一个地址(变量)
		4.make 为变量分配内存，分配的类型是引用类型，比如chan，map，slice
		5.append 用于追加元素到元素数组和slice
		6.panic&&recover 用来进行错误处理
*/
func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	v := 0
	a := 100 / v
	println(a)
}
func initconfig() (err error) {
	err = errors.New("init config failed")
	return err
}
func main() {
	//test()
	/*	var s=new(int)
		println(s)
		fmt.Printf("%v\n",*s)
		chanints := make(chan int)
		println(chanints)
		var arr []int
		arr=append(arr,'a')
		fmt.Println(arr)
		arr=append(arr,10,20,30)
		fmt.Println(arr)
		arr=append(arr,arr...)
		fmt.Println(arr)*/
	err := initconfig()
	if err != nil {
		fmt.Println(err)
	}
}
