package main

import "fmt"

func main() {
	var (
		age   int
		name  string
		isVIP bool
		score float32
	)

	// 方式一 fmt.Scanln()
	//fmt.Println("请输入学生的年龄:")
	//fmt.Scanln(&age)
	//
	//fmt.Println("请输入学生的姓名:")
	//fmt.Scanln(&name)
	//
	//fmt.Println("请输入学生的成绩")
	//fmt.Scanln(&score)
	//
	//fmt.Println("请输入学生是否为vip:")
	//fmt.Scanln(&isVIP)
	//
	//fmt.Println("年龄:", age)
	//fmt.Println("姓名:", name)
	//fmt.Println("成绩:", score)
	//fmt.Println("是否是vip:", isVIP)

	// 方式二
	fmt.Scanf("%d %s %f %t", &age, &name, &score, &isVIP)
	fmt.Println("年龄:", age)
	fmt.Println("姓名:", name)
	fmt.Println("成绩:", score)
	fmt.Println("是否是vip:", isVIP)

}
