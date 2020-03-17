package main

import (
	"fmt"
	"reflect"
)

//反射就是获取运行时的变量的类型和数据值主要有typeof 和valueof 两个函数
type student struct {
	name  string
	age   int
	score int
}

func test(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println(t)
	p := reflect.ValueOf(a)
	kind := p.Kind()
	fmt.Println(p)
	fmt.Println(kind)
}
func main() {
	var a student
	test(a)
}
