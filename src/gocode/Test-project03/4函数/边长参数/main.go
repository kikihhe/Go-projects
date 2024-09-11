package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	minVal := min(slice...)
	fmt.Println(minVal)
}

// 返回一个正整数数组中的最小值
func min(slice ...int) int {
	if len(slice) <= 0 {
		return 0
	}
	minVal := 0
	for _, val := range slice {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}
