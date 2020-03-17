package main

import "time"

/*
	用于区别值类型和引用数据类型
		值类型：int，float，boolean，数组，结构体,通常在栈内存中去分配
		引用类型：slice，chan，map，指针等，通常在堆内存进行分配
	难点：如何进行引用类型数据传递
*/
//用于观察值类型和应用类型的输出结果
//值类型输出的是结果值，引用类型输出的是结果值
func main() {
	var a = 3
	println(a) //期望值：3实际值：3
	modify(a)
	println(a) //期望值：3实际值：3
	/*	var pipe chan int
		pipe =make(chan int,1)
		pipe<-1
		println(&pipe)
		println(pipe)*/
	var b *int
	*b = 10
	println(&b) //期望值10，实际值1010
	modify_t(b)
	println(&b) //期望值100，实际值100
	time.Sleep(time.Second)
}
func modify(a int) {
	a = 10
}
func modify_t(a *int) {
	*a = 100
}
