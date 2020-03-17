package add

import "fmt"

var Name string = "张夫业"
var Age int = 25
var Hobby string

func Test() {
	Hobby = "学习"
	Name = "张天爱"
	Age = 27
}
func init() {
	fmt.Println("这是引用包的初始函数")
}
