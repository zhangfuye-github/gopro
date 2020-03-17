package main

import "fmt"

/*
	类型断言：就是一个接口类型的参数，传入任何类型的参数，通过判断，求出具体的传入类型
*/
func test(items ...interface{}) {
	for i, v := range items {
		//类型断言，用来判断传入的参数的类型。
		switch v.(type) {
		case int, int8, int16, int32, int64:
			fmt.Printf("param #%d is a int\n", i)
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		case student:
			fmt.Printf("param #%d is a struct\n", i)
		case *student:
			fmt.Printf("param #%d is a *student\n", i)
		}
	}
}

type student struct {
}

func main() {
	var s student
	var a *student
	test(1, "张夫业", 'A', 3.2, false, s, a)
}
