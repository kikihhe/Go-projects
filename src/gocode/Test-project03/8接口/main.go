package main

import "fmt"

type Shaper interface {
	area() int
}

type Rectangle struct {
	length int
	width  int
}

func (rec *Rectangle) area() int {
	return rec.length * rec.width
}

func main() {
	var shaper Shaper
	rec := &Rectangle{4, 2}
	shaper = rec
	fmt.Printf("shaper的类型: %T\n", shaper)
	println(shaper.area())

	// 可以通过类型断言得到 变量的真正类型
	if _, ok := shaper.(*Rectangle); ok {
		fmt.Println("shaper是Rectangle类型的")
	} else {
		fmt.Println("shaper不是Rectangle类型的")
	}
}
