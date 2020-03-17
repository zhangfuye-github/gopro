package main

import (
	"fmt"
	"time"
)

/*
	结构体中的匿名字段：结构体中的字段可以没有名字，如下：
	type student struct{
		name string
		age int
	}
	type person struct{
		student
		int
		second time.Time
	}
	但是，结构体中的匿名字段如何使用呢？
	匿名和继承：一个结构体嵌套了另一个匿名结构体，需要访问匿名结构体的方法时，实现了继承
*/
type person struct {
	name string
	age  int
	sex  string
}
type person1 struct {
	name string
}
type student struct {
	person
	person1
	int
	start time.Time
	//name string
}

type stu student

func main() {
	var s stu
	s.person1.name = "张天爱"
	s.age = 17
	s.int = 100
	s.start = time.Now()
	fmt.Println(s)
}
