package main

import "fmt"

/*
字符串与切片：字符串是类型为byte的数组，可以用切片区子串
内置函数：close，append，copy，len,cap,new,make,rocover,panic
如果想修改字符串的内容，要把字符串转换为数组，涉及到中文要用rune转换成字符数组
*/
func main() {
	var slice []int = []int{1, 2, 3, 4, 5, 6}
	var b []int = make([]int, 10)
	copy(b, slice)
	fmt.Println(b)
	var s string = "Hello World"
	s1 := s[:5]
	s2 := s[6:]
	fmt.Println(s1)
	fmt.Println(s2)
	s3 := []rune(s)
	s3[0] = '我'
	fmt.Println(string(s3))
}
