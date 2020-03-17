package main

import "fmt"

/*
	string 的两种表示方式：
			1."" 双引号括起来,都是经过转义的，串中不能换行
			2.``反引号括起来，与双引号的区别，反引号括起来的字符串不会转义，可以换行，原封原样
	byte 字符类型
			1.字符类型用单引号''
			2.字符类型默认输出是Ascall编码，要想输出字符形式，用格式化输入输出，此处体现出
				格式化输出的优势
	下面的程序运行过程中结果出现的顺序不一致，a的位置提前了，是否与格式化输入输出有关系，
			答:无关，这和go语言的特性，高并发有关系，会起一个携程
			张夫业
			a
			hello world!
			let's go !!!!!!!!!!!!

			97

*/
func main() {
	var a string
	a = "张夫业"
	var b string
	b = `hello world!
let's go !!!!!!!!!!!!
`
	println(a)
	println(b)

	var char byte
	char = 'a'
	println(char)
	fmt.Printf("%c\n", char)
}
