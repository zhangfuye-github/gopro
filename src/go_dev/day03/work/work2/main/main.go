package main

import "fmt"

/*
2.  一个数如果恰好等于它的因子之和，这个数就称为“完数”。例如6=1＋2＋3.
编程找出1000以内的所有完数。
*/
//求完数
func wanshu(a int) {
	if a == 0 {
		fmt.Printf("%d	", a)
	} else {
		var sum int = 0
		for i := 1; i < a; i++ {
			if a%i == 0 {
				sum = sum + i
			}
		}
		if a == sum {
			fmt.Printf("%d	", a)
		}
	}

}
func main() {
	fmt.Println("请输入一个数")
	var a int
	fmt.Scanf("%d", &a)
	for i := 0; i <= a; i++ {
		wanshu(i)
	}
}
