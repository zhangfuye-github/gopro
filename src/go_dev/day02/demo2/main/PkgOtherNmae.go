package main

import (
	"fmt"
	a "go_dev/day02/demo2/add"
)

/*
	1.包的别名引用，注意事项：一个包下值呢有一个main函数，多了会报错，乌龟的屁股，规定
*/
func main() {
	fmt.Println(a.Name)  //张夫业
	fmt.Println(a.Age)   //25
	fmt.Println(a.Hobby) //""
	a.Test()             //作用赋值
	fmt.Println(a.Name)  //张天爱
	fmt.Println(a.Age)   //27
	fmt.Println(a.Hobby) //学习
}
