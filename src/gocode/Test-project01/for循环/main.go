package main

import "fmt"

func main() {
	var sum int = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("从1加到100的值为:", sum)
}
