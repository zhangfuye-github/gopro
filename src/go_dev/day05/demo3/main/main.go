package main

import (
	"fmt"
	"math/rand"
)

type student struct {
	name string
	age  int
	next *student
}

func insert(stu *student) {
	p := stu
	for i := 0; i < 10; i++ {
		stu := &student{
			name: fmt.Sprintf("stu%d", i),
			age:  rand.Intn(30),
		}
		p.next = stu
		p = stu
	}
}
func main() {
	var stu *student = &student{
		name: "张天爱",
		age:  27,
	}
	insert(stu)
	var p *student
	p = stu
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}
