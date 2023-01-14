package main

import "fmt"

func main() {

	// 数组的定义
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 4
	// 数组的初始化:
	// 1. 声明大小、类型
	var arr1 [10]int
	// 2. 声明大小、类型、初始数据  可以不指定全部值，指定一部分后，未初始化的位置填充默认值
	var arr2 = [10]int{1, 2, 3, 4, 5}
	// 3. 不知道长度, 在运行后获得长度
	var arr3 = [...]int{1, 2, 3}
	// 4. 指定下标
	var arr4 = [...]int{1: 1, 2: 2, 5: 5, 4: 4, 3: 3, 0: 0}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(len(arr3))
	fmt.Println(arr4)

	// 使用len(数组)得到的是数组的大小
	//fmt.Println(len(arr))

	var sum = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Printf("%d", sum)

	// 输入5个数，求他们的和,
	var nums [5]int
	for i := 0; i < len(nums); i++ {
		fmt.Scanln(&nums[i])
	}
	fmt.Println(nums)

	// 遍历数组
	// 方式一: for循环 + 数组下标遍历
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	fmt.Println("-----------------------")

	// 方式二: for range 遍历
	// key: 下标
	// value: 值
	for key, value := range nums {
		fmt.Println("key: ", key, "\tvalue: ", value)
	}

}
