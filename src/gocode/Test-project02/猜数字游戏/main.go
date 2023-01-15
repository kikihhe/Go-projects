package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixMilli())
	answer := rand.Intn(maxNum)
	fmt.Println("生成的数字为:", answer)
	var i int
	for i = 0; i < 3; i++ {
		fmt.Print("请输入你猜的数字: ")
		var a int
		fmt.Scanln(&a)
		if a == answer {
			fmt.Println("你猜对啦，游戏结束!")
			break
		}
		if a > answer {
			fmt.Println("比答案大哦~")
		}
		if a < answer {
			fmt.Println("比答案小哦~")
		}
	}
	if i == 3 {
		fmt.Println("三次机会已经用完~游戏结束!")
	}

}
