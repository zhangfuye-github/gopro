package main

/*
	1.练习字符串的拼接

*/
func main() {
	var s1 string = "hello"
	var s2 string = "world"
	var s string
	s = s1 + " " + s2
	println(s)
	str := s[0:5] //类型是切片，和数组一样，取数范围表示从第几位后开始取数
	println(str)
	/*
		字符串类型不包含翻转，自己实现一个翻转
			两种方法：
				1：字符串拼接
				2：强制转换成数组，然后用拼接的方法，append
	*/
	reverse(str)
}
func reverse(s string) {
	var a string
	var s1 string
	k := len(s)
	for i := 0; i < k; i++ {
		a = s[k-i-1 : k-i] //可以通过切片实现取出每个数据元素
		s1 = s1 + a
	}
	println(s1)
}
