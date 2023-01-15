package main

import "fmt"

func main() {

	slice1 := make([]int, 0, 1)
	for i := 0; i < 10; i++ {
		slice1 = append(slice1, i)
		fmt.Printf("切片的长度为: %d\t\t", len(slice1))
		fmt.Printf("切片的容量为: %d\n", cap(slice1))
	}

	slice2 := make([]int, 0, 1)

	slice2 = append(slice2, 1, 2, 3, 4, 5, 6)
	fmt.Println("slice2的长度为:", len(slice2))
	fmt.Println("slice2的容量为:", cap(slice2))

	// 申请一个1030大小的切片
	slice3 := make([]int, 1030, 1030)
	for i := 0; i < 1030; i++ {
		slice3[i] = i
	}
	slice4 := make([]int, 1, 1)
	copy(slice4, slice3)
	fmt.Println("slice4的容量为:", cap(slice4))

}
