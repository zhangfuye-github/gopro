package main

import (
	"fmt"
	"strings"
)

/*
	闭包：一个函数和其相关的引用环境组成的实体,闭包的返回值是一个函数，返回的函数涉及
		到调用该函数的的变量
*/
func Adder() func(int) int {
	var x int
	return func(dela int) int {
		x += dela
		return x
	}
}
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func main() {
	var f = Adder()
	fmt.Print(f(20), "-")
	fmt.Print(f(20), "-")
	func1 := makeSuffixFunc(".doc")
	println(func1("HelloWorld"))
	func2 := makeSuffixFunc(".go")
	println(func2("HelloWorld"))
	println(func2("HelloWorld.go"))
}
