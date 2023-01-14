package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Now() 返回的是一个结构体，类型为 time.Time 这个Time有很多方法
	now := time.Now()
	fmt.Printf("now的类型为: %T\n", now)
	fmt.Println(now)

	// 获取Time中的 年月日.
	// 年: int类型
	// 月: time.Month类型，可以调用String()方法转换为string类型, 可以直接使用 int(month)转为数字
	// 日: int类型
	year, month, day := now.Date()
	fmt.Println("year:", year, "month:", int(month), "day:", day)
	fmt.Printf("年的类型为:  %T\n", year)
	fmt.Printf("月的类型为:  %T\n", month)
	fmt.Printf("日的类型为:  %T\n", day)

	// 获取Time中的 时分秒
	// 时分秒都为int类型,
	// 直接输出Time结构体，会发现: 2023-01-14 11:38:12.5628644 +0800 CST m=+0.004337001, 秒精确到七位数，但是使用Time.Clock()获得的秒截取整数
	hour, min, sec := now.Clock()
	fmt.Println("hour:", hour, "min:", min, "sec:", sec)
	fmt.Printf("hour的类型: %T\n", hour)
	fmt.Printf("min的类型: %T\n", min)
	fmt.Printf("sec的类型: %T\n", sec)

	// Time.YearDay() 返回今天是今年的第几天
	yearDay := now.YearDay()
	fmt.Println(yearDay)

	// Time.WeekDay 返回今天是这周的第几天, 返回值为time.Weekday类型，使用int(weekday)转为数字.
	weekday := now.Weekday()
	fmt.Printf("weekday的类型是: %T\n", weekday)
	fmt.Println((int)(weekday))

	// 日期的格式化, 输出特定格式的日期. (需要固定 2006-01-02 15:04:05)   返回值是一个字符串
	format := now.Format("2006-01-02 15:04:05")
	fmt.Println(format)
}
