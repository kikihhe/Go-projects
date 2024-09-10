package main

import "fmt"

//func testIf() int {
//	a := 1
//	if a == 1 {
//		return 3
//	} else if a == 2 {
//		return 4
//	} else {
//		return 7
//	}
//}

// 返回一个数的绝对值
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// 返回较大的数
func greater(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func testIf(n int) {
	if a := process(n); a < n {
		fmt.Println("n > 1")
	} else if a == 3 {

	} else {

	}

	fmt.Println()

}
func process(a int) int {
	return a + 4
}

// for循环累加1-n
func testFor(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
func printG(n int) {
	// 打印G, 形式如下:
	//    G
	//    G G
	//    G G G
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}

}
func printFizz() {
	// 写一个从 1 打印到 100 的程序，
	// 但是每当遇到 3 的倍数时，不打印相应的数字，但打印一次 "Fizz"。
	// 遇到 5 的倍数时，打印 Buzz 而不是相应的数字。
	// 对于同时为 3 和 5 的倍数的数，打印 FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func printT() {
	i := 2
	for i > 0 {
		i--
		fmt.Println("打印一次")
	}
}
func testForRange() {
	var str string = "你好帅"
	for i, value := range str {
		fmt.Printf("下标 %d 处的字符:%c\n", i, value)
	}
}
func main() {
	printG(5)
	testForRange()
}
