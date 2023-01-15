package main

import "fmt"

type Address struct {
	province string
	city     string
}

type People struct {
	id      int
	name    string
	age     int
	address Address
}

func main() {
	people := &People{
		id:   1,
		name: "小明",
		age:  18,
		address: Address{
			province: "河北",
			city:     "衡水",
		},
	}
	fmt.Println(people)

}
