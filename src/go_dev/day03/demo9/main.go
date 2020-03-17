package main

/*
练习14：写一个函数add，支持1个或多个int相加，并返回相加结果
练习15：写一个函数concat，支持1个或多个string相拼接，并返回结果
*/
func add(a int, arg ...int) (sum int) {
	sum = a
	lenth := len(arg)
	println(lenth)
	for i := 0; i < len(arg); i++ {
		sum = sum + arg[i]
	}
	return
}
func contact(a string, arg ...string) (s string) {
	s = a
	for i := 0; i < len(arg); i++ {
		//s=s+string(arg[i])
		s = s + arg[i]
		//time.Sleep(time.Millisecond)
	}
	return
}
func main() {
	sum := add(1, 3, 4, 5, 6)
	s := contact("1", "sss", "ddd", "zhang", "fu", "ye")
	println(sum)
	println(s)
}
