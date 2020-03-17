package main

import "fmt"

/*
	const 用来修饰常量，常量就是在运行过程当中，其值不会发生变化的量
	construction可以用来修饰：boolean,number,string
	多常量声明的三种方式：
		1.const a=123
		const b=false
		const c="张夫业"
		2. const(
				const a=123
				const b=false
				const c="张夫业"
			)
		3.const(
			a=iota		//表示从0开始进行常量的赋值
			b
			c
			d
			)
*/
func main() {
	/*const a=123
	const b=false
	const c="张夫业"*/
	const (
		a = iota //表示从0开始进行常量的赋值
		b
		c
		d
	)
	//期望结果a=0,b=1,c=2,d=3
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
