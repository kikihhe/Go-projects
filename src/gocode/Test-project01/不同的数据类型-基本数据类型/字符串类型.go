package main

import "fmt"

func main() {
	var a = "123何"
	var b = "123"
	fmt.Println(a == b)

	// 直接使用len判断的是字节数
	// 现在先不说如何判断字符数
	fmt.Println(len(a))
}
