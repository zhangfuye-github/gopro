package main

/*
	defer语句：
		1：用途：关闭句柄文件，解锁，关闭数据库连接
		2.先进后出，函数返回时执行该语句，
*/
func main() {
	i := 0
	i++
	defer println(i)
}
