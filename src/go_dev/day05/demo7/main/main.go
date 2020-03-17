package main

import "fmt"

/*
继承：一个结构体套用了一个匿名结构体，需要调用匿名结构体的方法时，就实现了继承。
*/
type person struct {
	name string
	age  int
	sex  string
}
type student struct {
	person
	class string
}

func (this person) run() {
	fmt.Println(this.name, "is runing")
}
func main() {
	var stu student
	stu.name = "张天爱"
	stu.run()
}
