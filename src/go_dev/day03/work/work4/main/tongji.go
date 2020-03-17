package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	.输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数。
*/
func tongji(str string) {
	t := []rune(str)
	length := len(t)
	var a, b, c, d, e int = 0, 0, 0, 0, 0
	for i := 0; i < length; i++ {
		//println(t[i])
		if t[i] <= 'z' && t[i] >= 'a' {
			a++
		} else if t[i] <= 'Z' && t[i] >= 'A' {
			b++
		} else if t[i] == ' ' {
			c++
		} else if t[i] <= '9' && t[i] >= '0' {
			d++
		} else {
			e++
		}
	}
	fmt.Printf("%s字符串中包含%d个小写字母，%d个大写字符，%d空格，%d个数字,%d个特殊字符\n", str, a, b, c, d, e)
}
func main() {
	for {
		println("请输入一行字符串：")
		/*		var str string
				fmt.Scanf("%s\r\n",&str)*/
		reader := bufio.NewReader(os.Stdin) //声明一个实例，类似于java的创建对象
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Printf("输入的数据有误")
		} else {
			tongji(string(line))
		}
	}

	/*	println(int('0'))//48
		println(int('9'))//57
		println(int('a'))//97
		println(int('z'))//122
		println(int('A'))//65
		println(int('Z'))//90
		println(int(' '))//32*/
}
