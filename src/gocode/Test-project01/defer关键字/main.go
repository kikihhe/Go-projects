package main

import "fmt"

func main() {
	//
	fmt.Println(test(10, 20))
}
func test(n1 int, n2 int) int {
	// 程序遇到defer关键字，不会立即执行defer语句，而是将这些语句放入一个栈中，先去执行defer下的语句，执行完后再将栈中的语句取出.
	// 并且defer会记录语句放入栈里时的状态，如果变量在后续语句中更改，不会影响defer输出结果.
	defer fmt.Println("n1的值为", n1) // 此时是10
	defer fmt.Println("n2的值为", n2) // 此时是20
	n1 += 10                       // n1 = 20
	n2 += 10                       // n2 = 30
	sum := n1 + n2
	fmt.Println("n1 + n2 =", sum)
	return n1 * n2
}
