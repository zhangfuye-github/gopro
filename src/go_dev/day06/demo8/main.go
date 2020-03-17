package main

import "fmt"

/*
	实现一个链表的通用接口
*/
func main() {
	var link link
	for i := 0; i < 10; i++ {
		link.insertfont(fmt.Sprintf("stu%d", i))
		//link.inserttail(fmt.Sprintf("stu%d",i))
	}
	link.trans()
}
