package main

import (
	"fmt"
	"math/rand"
)

/*
链表是一个线性数据结构，
和数组的区别是逻辑上的顺序存贮，在物理存储上地址既可以连续也可以不连续
*/
type student struct {
	name string
	age  int
	sex  string
	next *student
}

//头部插入
func insertfont(stu *student, p *student) {
	stu.next = p
	p = stu
	/*	for p!=nil{
		fmt.Println(*p)
		p=p.next
	}*/
}

//尾部插入
func insertback(stu *student, p *student) {
	p.next = stu
	p = stu
}

//循环遍历
func recoverlink(stu *student) {
	var p *student
	p = stu
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}
func main() {
	var stu2 student = student{
		name: "赵辉",
		age:  18,
		next: nil,
	}
	var stu1 student = student{
		name: "zhangfuye",
		age:  18,
		next: &stu2,
	}
	var stu3 student
	stu3.name = "张天爱"
	stu3.age = 27
	//recoverlink(&stu1)
	var p *student
	p = &stu2
	insertback(&stu3, p)
	var stu5 student
	for i := 0; i < 10; i++ {
		stu5.name = "zhangfuye" + string(i)
		stu5.age = rand.Intn(30)
	}
	recoverlink(&stu1)
	var stu4 student
	stu4.name = "张天爱"
	stu4.age = 27
	var q *student
	q = &stu1
	insertfont(&stu4, q)
	recoverlink(&stu4)
}
