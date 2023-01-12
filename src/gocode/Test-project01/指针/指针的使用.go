package main

import "fmt"

func main() {
	var a = 10
	var p1 *int = &a
	fmt.Println(&p1)
	fmt.Println(&a)
	fmt.Println(p1)
	fmt.Println(*p1)

	// 可以通过指针改变变量的值
	*p1 = 20
	fmt.Println("a变为", a)
	fmt.Println("*p为", *p1)

	// 指针的值必须是地址值
	var b int = 10 
	var p2 *int = &b 

	fmt.Println(*p2)


	// 指针的类型必须与变量的类型相同
	// var c int = 10 
	// var p3 *float32 = &c
	// fmt.Println(*p3)


	// 不同类型之间不能运算
	var d int32 = 20
	var e int64 = 21
	var f = int64(d) - e
	fmt.Println(f)

	// 
	var g uint32 = 100
	var h uint32 = 99
	var i uint32 = g - h
	fmt.Println(i)

}