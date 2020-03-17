package main

import (
	"fmt"
	"os"
)

func main() {
	var goos string = os.Getenv("PATH")
	fmt.Println(goos)
}
