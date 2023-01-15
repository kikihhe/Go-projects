package main

type People struct {
	// name、age是小写，其他包无法访问，
	// 但是可以提供get/set方法实现封装
	name string
	age  int
}

func NewPeople(name string, age int) *People {
	return &People{
		name: name,
		age:  age,
	}
}
func (people *People) SetAge(age int) {
	people.age = age
}

func (people *People) SetName(name string) {
	people.name = name
}
