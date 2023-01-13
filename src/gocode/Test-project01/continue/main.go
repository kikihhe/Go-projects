package main

import "fmt"

func main() {
	// 打印从1 - 10 内，不是3的倍数的数
	for i := 1; i < 11; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
