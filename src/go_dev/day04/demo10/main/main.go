package main

import (
	"fmt"
)

/*
	map:key-value的数据结构，又叫字典或者关联数组
		map的声明方式: var map1 map[keytype]valuetype
					var map2 map[keytype] map[key]valuetype
		声明是不分配内存的，分配内存要通过make，想想切片和管道
*/
func testMap() {
	//var map1 map[string]string
	//map1=make(map[string]string,100)//第一种初始化方法，通过make
	var map1 map[string]string = map[string]string{"key": "value"} //d第二种初始化方法，直接声明的
	//初始化
	map1["key"] = "zhang"
	fmt.Println(map1["key"])
}
func main() {
	testMap()
}
