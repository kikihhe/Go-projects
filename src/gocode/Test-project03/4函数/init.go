package main

import "fmt"

var num int = test()

func test() int {
	fmt.Println("test方法")
	return 1
}
func init() {
	fmt.Println("初始化函数执行!!!")
}

func main() {
	fmt.Println("main函数执行!!!")
}
