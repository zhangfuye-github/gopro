package main

import "fmt"

/*
golang中的方法：golang中的方法是是作用在特定类型的变量上，因此自定义类型都可以有方法
	不仅仅是struct.
	方法的定义方式：
		这里要注意方法和函数的区别：recevier type
		func (receviver type) methodname(参数列表)(返回值列表){

		}
*/
type student struct {
	name string
	age  int
}

func (this *student) init(name string, age int) {
	this.name = name
	this.age = age
}
func (this student) get() student {
	return this
}
func main() {
	var stu, stu1 student
	stu.init("张天爱", 27)
	fmt.Println(stu)
	stu1 = stu.get()
	fmt.Println(stu1)
}
