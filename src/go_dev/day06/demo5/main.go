package main

import "fmt"

/*
	由于接口是一般类型，不知道具体类型 ，如果转成具体类型，可以如下：
		var a int
		var b interface{}
		b.(int)
*/
type student struct {
	name string
	age  int
}

func test(a interface{}) {
	s, ok := a.(student)
	if ok == false {
		fmt.Println("can't convert")
	}
	fmt.Println(s)
}
func main() {
	/*	var a int
		fmt.Printf("%d%T\n",a,a)
		var b interface{}
		//a=b.(int)  错误，进口是空类型，不可以转换成int类型。
		//可以通过如下方法进行类型转换
		b=a//多态
		fmt.Printf("%d%T\n",b,b)*/
	var stu student
	test(stu)
}
