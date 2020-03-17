package main

//调用其他包下的变量
/*
	注意事项：
		1.注意变量首字母大写，go语言中没有权限访问控制关键字，通过首字母大小写，来确定变量的作用范围
		2.必须声明在函数外部（主）；不能声明在函数当中，声明在函数当中，该包下必须使用，否则报错
		3.声明在函数外部，在函数内部赋值的变量，想要引用赋值后的变量，需要引入赋值函数
		4.字符串类型string，s小写
		5.在函数之外不能存在单独的语句，如：i：=123，是不合法的
*/

/*func main() {
	fmt.Println(add.Name)		//张夫业
	fmt.Println(add.Age)//25
	fmt.Println(add.Hobby)//""
	add.Test()//作用赋值
	fmt.Println(add.Name)//张天爱
	fmt.Println(add.Age)//27
	fmt.Println(add.Hobby)//学习


}
*/
