package main

import "fmt"

func main() {
	map1 := make(map[string]int)
	map1["xiaoming"] = 13
	map1["xiaohe"] = 15
	map1["xiaolan"] = 20

	_, xiaoMingIsPresent := map1["xiaoming"]
	fmt.Println("小明存在:", xiaoMingIsPresent)
	if xiaoMingIsPresent {
		fmt.Println("小明的值", map1["xiaoming"])
	}

	delete(map1, "xiaoming")
	_, xiaoMingIsPresent = map1["xiaoming"]
	if !xiaoMingIsPresent {
		fmt.Println("小明已经被删了")
	}

	printMap(map1)
}

func printMap(map1 map[string]int) {
	for key := range map1 {
		value := map1[key]
		fmt.Println(key, value)
	}
}

// 对调map的key和value
func reverseMap(map1 map[string]int) map[int]string {
	reverse := make(map[int]string, len(map1))
	for key, value := range map1 {
		reverse[value] = key
	}
	return reverse
}
