package main

import (
	"fmt"
	"reflect"
)

type not struct {
	s1 string
	s2 int
	s3 int
}

func (this not) String() string {
	return this.s1
}

var a interface{} = not{"zhangfuye", 20, 50}

func main() {
	value := reflect.ValueOf(a)
	fmt.Println(value, "reflect.ValueOf(a)")
	typeOf := reflect.TypeOf(a)
	fmt.Println(typeOf, "reflect.TypeOf(a)")
	kind := value.Kind()
	fmt.Println(kind)
	fmt.Println(value.NumField())
	fmt.Println(value.NumMethod())
	fmt.Println(value.Method(0).Call(nil))
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		//value.Field(i).SetString("C#")
	}
}
