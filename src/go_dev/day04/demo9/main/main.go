package main

import (
	"fmt"
	"sort"
)

/*
	数组的排序和查找：排序都在sort包中，导入就可以使用了
		有如下方法:sort.Ints(),sort.Strings(),sort.Float64s()
	常用的查找方法有：sort.SearchFloat64s(),	sort.SearchInts(),sort.SearchStrings()
			注意：在查找之前要进行好排序
*/
func testSortInts() {
	var s = [...]int{1, 3, 4, 70, 5, 9, 2, 0}
	sort.Ints(s[:])
	fmt.Println(s)
}
func testSortStrings() {
	var s = [...]string{"A", "b", "zhou", "zhang", "zhao", "a"}
	sort.Strings(s[:])
	fmt.Println(s)
}
func testSortFloat64s() {
	var s = [...]float64{10.2, 3.2, 4.09, 70.56, 5.23, 9.14, 2.0, 0.0}
	sort.Float64s(s[:])
	fmt.Println(s)
}
func testSortSearchFloat64s() {
	var s = [...]float64{10.2, 3.2, 4.09, 70.56, 5.23, 9.14, 2.0, 0.0}
	sort.Float64s(s[:])
	i := sort.SearchFloat64s(s[:], 3.2)
	fmt.Println(i)
	//fmt.Println(len(s[:]))
}
func testSortSearchInts() {
	var s = [...]int{1, 3, 4, 70, 5, 9, 2, 0}
	sort.Ints(s[:])
	i := sort.SearchInts(s[:], 0)
	fmt.Println(i)
	i = sort.SearchInts(s[:], 2)
	fmt.Println(i)
	i = sort.SearchInts(s[:], 3)
	fmt.Println(i)
	i = sort.SearchInts(s[:], 4)
	fmt.Println(i)
	i = sort.SearchInts(s[:], 5)
	fmt.Println(i)
	i = sort.SearchInts(s[:], 9)
	fmt.Println(i)
	i = sort.SearchInts(s[:], 70)
	fmt.Println(i)
}
func main() {
	/*	testSortFloat64s()
		testSortStrings()
		testSortInts()*/
	testSortSearchInts()
	testSortSearchFloat64s()

}
