package main

import "fmt"

func main() {
	// return语句直接结束函数的执行
	for i := 10; i < 20; i++ {
		if i == 15 {
			return
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("als")
}
