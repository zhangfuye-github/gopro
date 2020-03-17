package main

import "fmt"

/*
写一个方法实现string()方法
	如果一个函数实现了String()方法，那么println就会调用这个变量的string()方法，进行输出
接口的定义涉及关键字interface，类似于结构体的定义，接口中只能定义方法，不能存在变量，
实现接口就是实现接口里面的全部方法，如果一个结构体没有实现一个接口的全部方法，不能进行多态
的使用，多态就是接口指向实现该接口的结构体。
*/
type car struct {
	name  string
	price int
}

func (this car) run() {
	fmt.Println(this.name, "is runing")
}

type sail interface {
	sleep()
	run()
}

func (this car) String() string {
	return fmt.Sprintf("name:%s,price:%d", this.name, this.price)
}
func (this car) sleep() {
	fmt.Println(this.name, "is sleeping")
}
func main() {
	var c car
	var t sail
	c.name = "大众"
	c.price = 180000
	fmt.Println(c)
	t = c //此处便是多态
	t.sleep()
}
