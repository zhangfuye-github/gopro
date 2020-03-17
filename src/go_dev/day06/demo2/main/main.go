package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/*
	该方法用来实现sort包下的sort接口，实现排序，首先去官网了解sort
*/
type Student struct {
	name  string
	age   int
	score float64
}
type studentArrray []Student

func (this studentArrray) Len() int {
	return len(this)
}
func (this studentArrray) Less(i, j int) bool {
	return this[i].name < this[j].name
}
func (this studentArrray) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func main() {
	var stu studentArrray
	for i := 0; i < 10; i++ {
		stu1 := Student{
			name:  fmt.Sprintf("stu%d", rand.Intn(100)),
			age:   rand.Intn(100),
			score: 100 * rand.Float64(),
		}
		stu = append(stu, stu1)
	}
	for _, v := range stu {
		fmt.Println(v)
	}
	sort.Sort(stu)
	fmt.Println("按姓名排序后：")
	for _, v := range stu {
		fmt.Println(v)
	}
}
