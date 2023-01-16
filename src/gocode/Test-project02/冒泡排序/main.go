package main

import (
	"fmt"
)

func BubbleSort(arr *[5]int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	arr := [5]int{1, 43, 5, 94, 90}
	BubbleSort(&arr)
	fmt.Println(arr)
}
