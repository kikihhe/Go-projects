package main

import "fmt"

func main() {
	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("arr的长度为: %d\n", cap(arr))
	slice := arr[0:10]
	//slice3 := arr[0:8]
	fmt.Printf("%p\n", &slice)
	//fmt.Printf("%p\n", &slice3)

	fmt.Println(slice)
	fmt.Printf("slice的长度为: %d\n", len(slice))
	fmt.Printf("slice的容量为: %d\n", cap(slice))
	fmt.Printf("slice的类型为: %T\n", slice)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("&arr[%d] = %p\t&slice[%d] = %p\n", i, &arr[i], i, &slice[i])
	}

	//ints := append(slice, 1)
	//fmt.Println(ints)

	var slice1 []int
	if slice1 == nil {
		fmt.Println("切片的元素数:", len(slice1))
		fmt.Println("切片的容量数:", cap(slice1))
		fmt.Println("切片的所有元素:", slice1)
		fmt.Println("slice1是空切片")
	}

	slice2 := make([]int, 5, 20)

	fmt.Println(slice2) //此时为空切片
	slice2[0] = 1
	slice2[1] = 23
	slice2[2] = 32
	slice2[3] = 154
	slice2[4] = 6
	//fmt.Println(len(slice2))
	//slice2[5] = 6
	//fmt.Println(len(slice2))
	//slice3 := append(slice2, 6)
	//fmt.Println(slice3) //此时已经有值
	slice4 := make([]int, 5, 5)
	i := copy(slice4, slice2)
	fmt.Println("copy的返回值为:", i)
	fmt.Println("slice4的值为:", slice4)

}
