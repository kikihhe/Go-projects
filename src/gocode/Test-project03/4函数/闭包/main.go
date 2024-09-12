package main

import "fmt"

// 函数功能 : 求和
// 函数名 : sum
// 参数 : 无
// 返回值 : 一个函数，这个函数接受一个int参数，返回一个int值
func sum() func(int) int {
	a := 0

	return func(num int) int {
		a += num
		return a
	}
}

func main() {
	function := sum()
	result1 := function(1)
	result2 := function(2)
	fmt.Println(result1)
	fmt.Println(result2)
}
