package main

/*
递归：函数自身调用自身,设置好出口条件
*/
import (
	"fmt"
	"time"
)

func test(n int) {
	if n > 10 {
		return
	}
	time.Sleep(time.Second)
	fmt.Println("张夫业")
	test(n + 1)
}

//递归函数求阶乘
func jiecheng(n int) int {
	//time.Sleep(time.Second)
	if n == 1 {
		return 1
	}
	return jiecheng(n-1) * n
}

//递归函数求斐波那契数列
func fibor(n int) int {
	if n <= 1 {
		return 1
	}
	return fibor(n-1) + fibor(n-2)
}
func main() {
	/*	fmt.Println("hello")
		time.Sleep(time.Second)
		test(0)*/
	a := jiecheng(5)
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		b := fibor(i)
		fmt.Println(b)
	}

}
