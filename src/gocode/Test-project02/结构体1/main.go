package main

import "fmt"

type People struct {
	name string
	age  int
}

func main() {
	//people := new(People)
	//var people People
	//people.name = "小明"
	//people.age = 18
	//fmt.Println(people)
	people := &People{
		name: "小明",
		age:  18,
	}

	people1 := People{
		name: "小明",
		age:  18,
	}

	fmt.Printf("people的类型为: %T\n", people)
	fmt.Printf("people的类型为: %T\n", people1)

}
