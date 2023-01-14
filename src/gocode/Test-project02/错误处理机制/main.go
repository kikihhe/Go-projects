package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("错误被捕获")
			fmt.Printf("err的类型是: %T\n", err)
			fmt.Println("具体的错误信息为:", err)
		}
	}()

	test2(10, 1)
	fmt.Println("执行test()下的语句")
}

// 程序错误
func test1() int {
	num1 := 10
	num2 := 0
	return num1 / num2
}

// 自定义错误，自定义的错误，需要在调用方法时由程序员处理
// 如果想让程序报错，可以调用panic(err error) 这样就可以被 defer + recover机制 捕获
func test2(num1 int, num2 int) (int, error) {
	// 如果被除数为1，抛出"被除数为1，没有意义"的错误
	if num2 == 1 {
		err := errors.New("被除数为1，没有意义")
		panic(err)
	}
	return num1 / num2, nil
}
