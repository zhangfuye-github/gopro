package main

import (
	"strings"
)

/*
	主要是讲解：strings的主要函数和strconv包的使用情况
		1.string:
			1)HasPrefix函数：用于判断字符串是否存在指定的""开头
					func HasPrefix(s, prefix string) bool {
						return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
					}
			2)HasSuffix函数：用于判断字符串是否以suffix结尾
					func HasSuffix(s, suffix string) bool {
						return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
					}
			3). strings.Index(s string, str string) int：判断str在s中首次出现的位置，如果没有

*/
//练习1：判断一个url是否以http://开头，如果不是，则加上http://
//练习2：判断一个路径是否以“/”结尾，如果不是，则加上/。
func main() {
	/*	var str string="zhangfuye"
		a:=strings.HasPrefix(str,"zh")
		println(a)
		strings.HasSuffix()*/
	/*	var str string
		str1:="http://"
		println("请输入网址")
		fmt.Scanf("%s\n",&str)
		//str="www.baidu.com"
		if(strings.HasPrefix(str,str1)){
			println(str)
		}else{
			println(str1+str)
		}
		println("请输入路径：")
		str2:="\\"
		fmt.Scanf("%s\n",&str)
		//str="D:\project\src\go_dev\day03\demo1"
		if(strings.HasPrefix(str,str2)){
			println(str)
		}else{
			println(str+str2)
		}
	*/ //. strings.Index(s string, str string) int：判断str在s中首次出现的位置，如果没有
	var name string = "zzhangfuye"
	a := strings.Index(name, "f")
	b := strings.Index(name, "b")
	c := strings.LastIndex(name, "z")
	println(a, b, c)
	println()
	var name1 = "     zh  ang     "
	name2 := strings.TrimSpace(name1) //去掉字符串的前后的空格
	println(name2)
	name3 := strings.Trim(name2, " ")
	println(name3)
	name4 := strings.Trim(name2, "zh")
	println(name4)
	name5 := strings.Trim(name4, "")
	println(name5)
	name6 := strings.Replace(name, "z", "a", 1)
	name7 := strings.Replace(name, "z", "a", 2)
	name8 := strings.Replace(name, "z", "b", -1) //-1表示替换所有的，
	// n表示替换的次数
	println(name6, name7, name8)
	number := strings.Count(name, "a") //统计字符出现的次数
	println(number)
	println(strings.Contains(name, "z"))
	println(strings.Contains(name, "zz"))
	println(strings.ContainsAny(name, "bddd"))
	println(strings.ToUpper(name))
	println(strings.ToLower(strings.ToUpper(name)))
	println(strings.Repeat(name, 3))
	name9 := "eaeae"
	println(strings.TrimLeft(name9, "e")) //注意事项必须是去掉首字母，填写的字符必须和首字母一致，然后从左输出
	println(strings.TrimRight(name9, "e"))
	println(strings.TrimLeft(name9, "a"))
	println(strings.TrimRight(name9, "a"))
	println(strings.Fields("123 321 123"))
}
