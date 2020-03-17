package main

import (
	"fmt"
)

/*
	continue  和 break
*/
func main() {
	str := "hello world,中国"
	for i, v := range str {
		//fmt.Printf("index[%d]val[%c]len[%d]\n",i,v,len([]byte(v)))
		fmt.Printf("index[%d] val[%c] len[%d]\n", i, v, len([]byte(string(v))))
	}
}
