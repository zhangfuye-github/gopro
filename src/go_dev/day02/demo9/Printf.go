package main

import "fmt"

/*
	格式化输入输出的一个重要应用，把其他的类型转换成自己想要的类型
*/
func main() {
	var a int
	a = 100
	fmt.Printf("%v\n", a)
	fmt.Printf("%q\n", a) //会转换成相应的ASCALL编码，进行输出
	fmt.Printf("%c\n", a) //也可以进行占位符进行相应的转换，与%q占位符的区别在于，是否有引号
	fmt.Printf("%T\n", a)
	str := fmt.Sprintf("%d", a)
	fmt.Printf("%v\n", str)
	fmt.Printf("%q\n", str) //如果是字符串，会加上双引号输出
	fmt.Printf("%T\n", str) //查看类型
}
