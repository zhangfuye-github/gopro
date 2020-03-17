package main

import (
	"fmt"
	"time"
)

/*
	获取当前时间秒，如果可以被2整除，输出female，否则输出male
*/
func main() {
	const (
		MALE   = 1
		FEMALE = 2
	)
	scd := time.Now().Unix()
	fmt.Println(scd)
	if (scd % FEMALE) != 1 {
		fmt.Println(FEMALE)
	} else {
		fmt.Println(MALE)
	}

}
