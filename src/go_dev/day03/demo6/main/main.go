package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	练习：猜数字，写一个程序，随机生成一个0到100的整数n，然后用户在终端，
		输入数字，如果和n相等，则提示用户猜对了。如果不相等，则提示用户，大于
		或小于n。
*/
func main() {
	rand.Seed(time.Now().Unix())
	var a = rand.Intn(100)
	var b int
	for {
		fmt.Println("请输入一个整数")
		fmt.Scanf("%d", &b)
		switch a < 100 {
		case a < b:
			println("您输入的值大了")
		case a > b:
			println("您输入的值小了")
		default:
			println("恭喜你，猜对了", a)
			return
		}
	}

}
