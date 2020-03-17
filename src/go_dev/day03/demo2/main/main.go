package main

import (
	"fmt"
	"time"
)

/*
	时间和日期函数：
		涉及time包，
*/
func main() {
	var time1 = time.Now().UnixNano()
	var time3 = time.Now().Format("2006/01/02 15:04:05")
	fmt.Println(time3)
	time2 := time.Now().UnixNano()
	fmt.Println(time2 - time1)
}
