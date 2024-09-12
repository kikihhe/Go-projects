package main

import "fmt"

var num int = 10
var numx2, numx3 int

func test(a int) {
	fmt.Println(a)
}

func test01(num1 int, num2 int, testFunc func(int)) {
	testFunc(num1)
	testFunc(num2)
}
func main() {
	//numx2, numx3 = getX2AndX3(num)
	//PrintValues()
	//numx2, numx3 = getX2AndX3_2(num)
	//PrintValues()
	//
	a := test
	fmt.Printf("a类型: %T, test函数类型: %T\n", a, test)
	fmt.Println(a)
	fmt.Println(test)
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}

// 接收两个整数，然后返回它们的和、积与差。编写两个版本，一个是非命名返回值，一个是命名返回值。
func getSumMulAndDel(a int, b int, c int) (int, int, int) {
	sum := a + b + c
	mul := a * b * c
	del := a - b - c
	return sum, mul, del
}

func getSumMulAndDelWithName(a int, b int, c int) (sum int, mul int, del int) {
	sum = a + b + c
	mul = a * b * c
	del = a - b - c
	return
}
