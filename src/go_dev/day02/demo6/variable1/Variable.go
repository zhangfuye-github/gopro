/*
	用来获取配置的GO的环境变量,通过OS包进行获取操作系统相关的数据
*/
package variable1

import (
	"fmt"
	"os"
)

func main() {
	var String string = os.Getenv("JAVA_HOME")

	var String1 = os.Getenv("path")
	fmt.Println(String)
	fmt.Println(String1)
	fmt.Println(os.Getenv("CLASSPATH"))
	fmt.Println(os.Getenv("GOROOT"))
	fmt.Println(os.Getenv("GOPATH"))
}
