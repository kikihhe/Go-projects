package main

import "fmt"

func main() {
	//var arr [5]int
	//fmt.Println(len(arr))
	//fmt.Println(cap(arr))
	//fmt.Println(arr)
	//
	//a := [...]string{"aa", "bb", "cc", "dd"}
	//for i := range a {
	//	fmt.Println("Array item", i, "is", a[i])
	//}

	arr1 := new([5]int)
	var arr2 [5]int
	fmt.Printf("arr1的类型: %T\n", arr1)
	fmt.Printf("arr2的类型: %T\n", arr2)

	updateArrValue(arr2)
	fmt.Println(arr2)

	updateArrPointValue(arr1)
	fmt.Println(*arr1)

	printArr(arr2)
	fmt.Println(arr2)

	arr3 := [4][4]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}
	printMulArr(arr3)

	str := reverseString("1234")
	fmt.Println(str)

}

// 直接操作数组，arr形参是实参的拷贝，修改arr不会影响原数组
func updateArrValue(arr [5]int) {
	arr[1] = 4
}

// 操作指针变量，修改arr会影响到原数组
func updateArrPointValue(arr *[5]int) {
	arr[1] = 4
}

// 将下标赋值给数组并打印
func printArr(arr [5]int) {
	for index := range arr {
		arr[index] = index
	}
	fmt.Println(arr)
}

// 打印多维数组
func printMulArr(arr [4][4]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Print(arr[i][j])
		}
		fmt.Println()
	}
}

func testSlice() {

	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("arr的长度为: %d\n", cap(arr))
	slice := arr[0:10]
	fmt.Printf("%p\n", &slice)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("&arr[%d] = %p\t&slice[%d] = %p\n", i, &arr[i], i, &slice[i])
	}

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
	slice4 := make([]int, 5, 5)
	i := copy(slice4, slice2)
	fmt.Println("copy的返回值为:", i)
	fmt.Println("slice4的值为:", slice4)

}

func testSliceCap() {
	slice1 := make([]int, 0, 10)
	// 一开始 slice1 的len为0，每次重组都变大一个
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("len(slice1) = %d\n", len(slice1))
	}

}

func reverseString(str string) string {
	//slice := make([]byte, 4)
	//for i := len(str) - 1; i >=0; i-- {
	//	slice[len(str) - i - 1] = str[i]
	//}
	//return string(slice)

	slice := make([]byte, 0)
	for i := len(str) - 1; i >= 0; i-- {
		slice = append(slice, str[i])
	}
	return string(slice)
}
