package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a int = 10

	fmt.Println("a的类型为: %T\n", a)

	var b float32 = float32(a)

	fmt.Printf("a的类型为: %T\n", a)
	fmt.Printf("b的类型为: %T\n", b)

	// 大值变小值，会出现转圈圈
	// 小值变大值，可以，不报错，只要显式指定
	var c int32 = 128
	var d int8 = int8(c)
	fmt.Println("转化之后的值为", d)

	var e float32 = 2.2
	var f int = int(e)
	fmt.Println("f的值为", f)

	var g string = fmt.Sprintf("%f", e)
	fmt.Printf("g的类型为: %T\n", g)
	fmt.Println("g的值为:", g)
	fmt.Printf("g的值为: %v\n", g)

	var h bool = false
	var i string = fmt.Sprintf("%t", h)
	fmt.Printf("i的类型为: %T\n", i) // string
	fmt.Println("i的值为:", i)      // false

	// j := true
	// k, _:= strconv.ParseBool(j)
	var j string = "tue"
	var k, l = strconv.ParseBool(j)
	fmt.Println("k的值为:", k)
	fmt.Printf("k的类型为: %T\n", k)

	fmt.Println("第二个参数:", l)

	var m int64 = 8
	var n string = strconv.FormatInt(m, 2)

	fmt.Println("n的值为", n)

	var o string = "8"
	var p, _ = strconv.ParseInt(o, 10, 64)
	fmt.Println("p的值为", p)

	var o1 string = "-8"
	var p1, _ = strconv.ParseUint(o1, 10, 64)
	fmt.Println("p的值为", p1)

	var q string = "23.47"
	var r, _ = strconv.ParseFloat(q, 64)
	fmt.Println("r的值为:", r)

	//var s string = "GoLang"
	//var t int32  = 20
	//t, _ = strconv.ParseInt(s, 10, 32)
	//fmt.Println("t的值为:", t)

}
