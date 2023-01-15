package main

import "fmt"

type People struct {
	name   string
	age    int
	height float32
}

type Monkey struct {
	name   string
	age    int
	height float32
}

func main() {
	// 声明一个People类型的结构体
	people := &People{
		name:   "小明",
		age:    18,
		height: 180.4,
	}
	// 将上述结构体变量转换为 *Monkey 类型
	monkey := (*Monkey)(people)
	fmt.Println("monkey:", monkey)

}
