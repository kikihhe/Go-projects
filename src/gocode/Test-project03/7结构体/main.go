package main

import (
	"fmt"
	"strings"
)

type User struct {
	id       int
	name     string
	password string
}

func NewUser(id int, name string, password string) *User {
	if id < 0 {
		return nil
	}
	return &User{id, name, password}
}

// 格式化id, 变成string并在后面加上俩空格
func (user *User) FormatId() string {
	return string(user.id) + "  "
}

func main() {
	user := new(User)
	user.id = 1
	user.name = "xiaoming"
	user.password = "123456"

	fmt.Println(user)
	fmt.Println(user.id)
	fmt.Println(user.name)
	fmt.Println(user.password)

	user1 := User{1, "xiaohong", "123"}
	fmt.Println(user1)

	upperName1(user)
	upperName2(user1)
	fmt.Println(user)
	fmt.Println(user1)

}

func upperName1(user *User) {
	user.name = strings.ToUpper(user.name)
}
func upperName2(user User) {
	user.name = strings.ToUpper(user.name)
}
