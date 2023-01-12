package main

import "fmt"

func main() {
	if 5 > 4 {
		fmt.Println("5 > 4 正确")
	}

	var age int
	fmt.Println("请输入你的年龄")
	fmt.Scanln(&age)
	if age < 18 {
		fmt.Println("未成年，犯法哦~ 三年起步~")
	} else if age > 18 && age <= 45 {
		fmt.Println("成年，而且是中年..")
	} else if age > 45 && age <= 80 {
		fmt.Println("老年人...")
	} else if age > 80 && age <= 120 {
		fmt.Println("能过这么大岁数也不容易", age)
	} else {
		fmt.Println("不要乱输哦，谁能活这么长时间?")
	}

	// 在if后面可以加上变量的定义
	if count := 20; count > 10 {
		fmt.Println("20 大于 10")
	}

}
