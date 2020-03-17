package main

import "fmt"

/*
	判断一个类型是否实现了该接口：b, ok := b.(person) b 和person都是接口

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
func main() {
	var athl *athlete
	var b interface{} //因为判断是否实现接口时要两边都是接口，因此可以定义一个空接口
	//来接收要判断的结构体
	b = athl
	b, ok := b.(person)
	fmt.Println(b, ok)
}
