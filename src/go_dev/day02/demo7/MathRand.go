package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		//rand.Seed(time.Now().UnixNano()) 需要放在循环的外部，可以放在init函数当中实现初始化
		println(rand.Intn(100))
		println(rand.Float32())
	}

}
