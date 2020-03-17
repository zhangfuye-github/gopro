package main

import (
	"encoding/json"
	"fmt"
)

/*
二叉树：二叉树的遍历，要用到递归
结构体是用户单独定义的数据类型，不能和基本数据类型进行转换，
给结构体进行别名：
		type stu student
		stu和student类型虽然字段一样，但是属于两种不同的类型
		golang结构体中没有构造函数，一般使用工厂模式来解决这个问题。
		我们可以为struct中的每个字段写上一个tag，这个tag可以通过反射的机制获取到，
	最长用的就是json的序列化和反序列化
		go中反引号的应用：1.tag  2.反引号用来创建 原生的字符串字面量
*/
type student struct {
	Name string `json:"name"` //要想被json访问，首字母要大写
	Age  int    //`json:"age"` tag可以用来控制json字符串key的大小写
}
type stu student

func main() {
	var stu1 stu
	stu1.Name = "张天爱"
	stu1.Age = 27
	bytes, e := json.Marshal(stu1)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(string(bytes))
}
