package main

import "fmt"

/*
.  输入一个字符串，判断其是否为回文。回文字符串是指从左到右读和从右到
左读完全相同的字符串。
*/
func huiwen(str string) {
	//首先考虑中文：中文占三个字节，英文占一个字节
	//首先取字节改成取字符
	t := []rune(str)
	i := len(t)
	for j := 0; j <= (i / 2); j++ {
		if t[i-j-1] != t[j] {
			fmt.Printf("%s不是回文序列\n", str)
			break
		} else {
			if j == (i / 2) {
				fmt.Printf("%s是回文序列\n", str)
			} else {
				continue
			}
		}
	}
}
func main() {
	for {
		fmt.Println("请输入一个字符串：")
		var str string
		fmt.Scanf("%s", &str)
		huiwen(str)
	}
}
