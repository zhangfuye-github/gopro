package main

import (
	"time"
)

func main() {
	var i int
	for i = 0; i < 100; i++ {
		go GO_ROUTE(i)
	}
	time.Sleep(10 * time.Second)
}
