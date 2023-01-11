package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var isGood bool = true
	var isBad bool = false
	var a int = 11111
	fmt.Println("isGood的值为", isGood)
	fmt.Println("isBad的值为", isBad)
	fmt.Println(a)


	var age = 10
	fmt.Printf("age的类型为: %T\n", age)
	fmt.Println("age占用的字节:", unsafe.Sizeof(age))



}