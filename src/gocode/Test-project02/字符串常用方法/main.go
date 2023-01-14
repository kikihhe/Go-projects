package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// len(str) 计算字符串字节长度
	length := len("你好,golang")
	println(length)

	// 字符串遍历
	//	1. 使用 range + for循环 对字符串遍历
	for _, value := range "你好哦" {
		fmt.Printf("%c", value)
	}
	fmt.Println()
	// 	2. rune + for 切片遍历
	r := []rune("你好哦")
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c", r[i])
	}

	// 字符串转整数
	atoi, _ := strconv.Atoi("1234")
	fmt.Printf("atoi的类型为: %T\n", atoi)
	fmt.Println(atoi)

	// 整数转字符串
	itoa := strconv.Itoa(123)
	fmt.Printf("itoa的类型为: %T\n", itoa)
	fmt.Println(itoa)

	// 计算str里有多少个sep strings.Count(str string, sep string)
	count := strings.Count("你好，你好，你很好啊，你好", "你好")
	fmt.Println("出现了", count, "次")

	// 区分大小写的比较: 直接用 ==
	fmt.Println("abc" == "abC")

	// 不区分大小写的比较 strings.EqualFold(s1 string, s2 string)
	isEqual := strings.EqualFold("ABCDE", "abcDe")
	fmt.Println(isEqual)

	// 返回sep在str出现的第一个下标 strings.Index(str string, sep string)  不存在就返回 -1
	index := strings.Index("abcdefg", "c")
	fmt.Println("第一次出现的下标为:", index)

	// 将字符串中前n个old子字符串替换为new新字符串  strings.Replace(s string, old string, new string, n int)
	// 如果str中不包含old字符串，返回原字符串
	// 如果n等于0，不替换；如果n小于0，替换所有old子字符串
	replace := strings.Replace("123456343434", "34", "12", -2)
	fmt.Println(replace)
	// 将str中所有old字符串替换为new字符串 strings.ReplaceAll(s string, old string, new string)
	// 底层为直接调用 strings.replace(s, old, new, -1)
	all := strings.ReplaceAll("121212", "12", "34")
	fmt.Println(all)
}
