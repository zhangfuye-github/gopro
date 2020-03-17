package main

import "fmt"

/*
	此案例介绍接口嵌套：接口嵌套就是接口当中含有匿名接口，也就是含有了匿名接口的所有方法
本案例的重点就是：了解认识多态的重要性，如果生命一个接口类型的变量，那么要想使该接口的变量
指向某结构体变量，那么该结构体变量要实现该接口的全部方法，因此接口和变量不要重名，否则系统会
无法识别，会有编译错误。
*/
type person interface {
	names()
	ages()
}
type student interface {
	scores()
	run()
}
type runner interface {
	person
	student
	eat()
}
type athlete struct {
	name  string
	age   int
	food  string
	score float64
}

//实现person的names方法，也就实现了runner接口嵌套person的names方法。
func (this *athlete) names() {
	fmt.Println("my name is:", this.name)
}
func (this *athlete) ages() {
	fmt.Println("my age is:", this.age)
}
func (this *athlete) run() {
	fmt.Println(this.name, "is running")
}
func (this *athlete) scores() {
	fmt.Println("my score is:", this.score)
}
func (this *athlete) eat() {
	fmt.Println(this.name, "is eatting", this.food)
}

//定义一个传输接口类型参数的函数
func test(rw runner) {
	rw.names()
	rw.eat()
	rw.run()

}
func main() {
	//var p runner
	var athlete1 athlete
	athlete1.name = "张天爱"
	athlete1.age = 27
	athlete1.score = 95.5
	athlete1.food = "humberger"
	//p=&athlete1
	/*	athlete1.names()
		p.eat()*/
	test(&athlete1)
}
