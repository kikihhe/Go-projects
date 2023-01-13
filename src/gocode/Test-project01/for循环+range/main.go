package main

import "fmt"

func main() {
	var name string = "你好"

	for i, value := range name {
		fmt.Println("索引:", i, "值为:", value)
	}
}
