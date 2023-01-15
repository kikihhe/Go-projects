package main

import (
	"fmt"
	"strconv"
)

type Phone struct {
	name  string
	price int
}

func NewPhone(name string, price int) Phone {
	return Phone{
		name:  name,
		price: price,
	}
}

type People struct {
	name   string
	phones []Phone
}

func NewPeople(name string, phones []Phone) People {
	return People{
		name:   name,
		phones: phones,
	}
}

func main() {
	// 创建一个People数组
	people := NewPeople("小明", make([]Phone, 5, 200))
	// 循环给people数组中的值赋值
	for i := 0; i < 5; i++ {
		phoneName := "小米" + strconv.Itoa(i)
		phone := NewPhone(phoneName, 1000*(i+1))
		//
		people.phones[i] = phone
		// 如果在初始化切片时指定长度为0或者小于5，使用append()扩容
		// people.phones = append(people.phones, phone)
	}
	fmt.Println(people)

}
