package main

import (
	"fmt"
	"time"
)

/*
	三个作业：
		1.求素数，判断 101-200 之间有多少个素数，并输出所有素数。
		2.求水仙花数，打印出100-999中所有的“水仙花数”，所谓“水仙花数”是指一
			个三位数，其各位数字立方和等于该数本身。
		3.求阶乘，对于一个数n，求n的阶乘之和，即： 1！ + 2！ + 3！+…n!
*/
//判断 101-200 之间有多少个素数，并输出所有素数。
func prime(a int, n int) {
	var sum int = 0
	for i := a; i <= n; i++ {
		flag := false
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = true
				break
			}
		}
		if flag == false {
			fmt.Printf("%5d", i)
			sum++
			if sum%5 == 0 {
				fmt.Printf("\n")
			}
		}
	}
	time.Sleep(time.Second)
	fmt.Printf("\n")
	println("素数个数为:", sum)
}

//求水仙花数，打印出100-999中所有的“水仙花数”，
func Narcissus(m int, n int) {
	var a int //个位
	var b int //十位
	var c int //百位
	sum := 0
	for i := m; i < n; i++ {
		a = i % 10
		b = i % 100 / 10
		c = i / 100
		if a*a*a+b*b*b+c*c*c == i {
			fmt.Printf("%5d", i)
			sum++
			if sum%5 == 0 {
				fmt.Printf("\n")
			}
		}
	}
	time.Sleep(time.Second)
	fmt.Printf("\n")
	println("水仙花个数为:", sum)
}

//求阶乘，对于一个数n，求n的阶乘之和，即： 1！ + 2！ + 3！+…n!
func Factorial(a int) {
	var product int = 1
	var b int = 0
	for i := 1; i <= a; i++ {
		product = product * i
		b = b + product
	}
	time.Sleep(time.Second)
	//fmt.Printf("%d的阶乘之和为：%d",a,b)
	println(a, "的阶乘之和为：", b)
}
func main() {
	prime(101, 200)
	Narcissus(100, 999)
	//time.Sleep(3*time.Second)
	Factorial(10)
}
