package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

func main() {
	//a := 1
	//b := 2.0
	//fmt.Print(a <= b)
	//testString()
	isEquals := equalsString("abc", "abc")
	fmt.Println(isEquals)
}

// 将int数据转为无符号整数
func Uint8FromInt(n int) (uint8, error) {
	if n >= 0 && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d 超过了uint8的范围", n)
}

func testBool() {
	//a := true
	//b := true
	//var c int8 = int8(a)
	//fmt.Print(c == b)
}

type TT int
type TZ string

func testType() {
	var a, b TT = 2, 1
	fmt.Println(a, b)
}

func testString() {
	str := "1234"
	fmt.Println(len(str))

}

// 判断两个字符串是否相等
func equalsString(str1 string, str2 string) bool {
	return bytes.Equal([]byte(str1), []byte(str2))
}

// 判断字符串是否只包含空格
func stringEmpty(str string) bool {
	//return strings.TrimSpace(str) == ""
	return len(strings.TrimSpace(str)) == 0
}
