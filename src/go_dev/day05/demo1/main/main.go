package main

import "fmt"

/*
	本节内容主要讲解结构体（struct）和方法(func)以及接口(interface)；
		1.结构体用来自定义复杂的数据结构
		2.方法用来对数据结构进行操作
		3.数据类型=数据结构+方法（相当于java中的类）
	接口是特殊的类，自己要实现接口中的方法。
		struct 定义方法和函数的区别：
*/
type student struct {
	name string
	age  int
	sex  string
}

func main() {
	//第一种初始化方法
	fmt.Println("第一种初始化方法")
	var stu1 student
	stu1.name = "张夫业"
	stu1.age = 23
	stu1.sex = "nan"
	fmt.Println(stu1)
	fmt.Println(&stu1.name, &stu1.age, &stu1.sex) //说明结构体的字段之间是连续的
	fmt.Println("==========================")
	fmt.Println("第二种初始化方法")
	var stu2 student = student{
		name: "赵辉",
		age:  22,
	}
	fmt.Println(stu2)
	fmt.Println("==========================")
	var stu3 *student = &student{
		name: "张三",
		age:  20,
		sex:  "待定",
	}
	fmt.Println(*stu3)
	fmt.Println("==========================")
}
