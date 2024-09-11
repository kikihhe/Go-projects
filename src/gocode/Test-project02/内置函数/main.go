package main

import "fmt"

func main() {
	// fmt、time、strconv这种函数是需要导包的函数
	// 像 print、println、len，这样不需要导包的函数就是Go内置的函数

	// new(Type) *Type 分配空间  主要分配值类型，如: int float bool string 数组 7结构体
	// 形参为需要分配的类型
	// 返回值为该类型的指针，指向地址中存储一个该类型的默认值  new(int) 分配一个int大小的内存，返回的值是一个 *int指针。
	i := new(int)
	fmt.Printf("i的类型为: %T\n", i)
	fmt.Println("i的值为", *i)

	// make(t Type, size ...IntegerType) 主要分配引用类型
	// 未完待续...

}
