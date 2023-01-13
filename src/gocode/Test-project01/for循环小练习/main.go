package main

import "fmt"

func main() {
	//var sum int

	//// 求1-100的和，当和到300时清零
	//for i := 1; i < 101; i++ {
	//	sum += i
	//	if sum >= 300 {
	//		sum = 0
	//	}
	//}
	//fmt.Println(sum)

	//// 求1-100的和，当和到300时停止程序
	//for i := 1; i < 101; i++ {
	//	sum += i
	//	if sum >= 300 {
	//		break
	//	}
	//}
	//fmt.Println(sum)

	//for i := 1; i < 4; i++ {
	//	for j := 2; j <= 4; j++ {
	//		fmt.Println("i为:", i, " j为:", j)
	//		if i == 2 && j == 3 {
	//			break
	//		}
	//	}
	//}

	// 使用标签指定哪一个for循环结束

lable:
	for i := 1; i < 4; i++ {

		for j := 2; j <= 4; j++ {
			fmt.Println("i为:", i, " j为:", j)
			if i == 2 && j == 3 {
				break lable
			}
		}
	}

}
