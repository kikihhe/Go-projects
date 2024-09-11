package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "你好，世界！欢迎来到 Go 语言的世界。"

	// strings.HasPrefix(str, prefix) bool ：str 是否以 prefix 开头
	prefix := "你好"
	fmt.Println("是否以 '你好' 开头:", strings.HasPrefix(str, prefix)) //true

	// strings.HasSuffix(str, suffix) bool ：str 是否以 suffix 结尾
	suffix := "世界。"
	fmt.Println("是否以 '世界。' 结尾:", strings.HasSuffix(str, suffix))

	// strings.Contains(str1, str2) bool ：str1 是否包含 str2
	substr := "Go 语言"
	fmt.Println("是否包含 'Go 语言':", strings.Contains(str, substr))

	// strings.Index(s, str) int ：str 在 s 中出现的首位置
	index := strings.Index(str, "世界")
	fmt.Println("‘世界’首次出现的位置:", index)

	// strings.LastIndex(s, str) int ：str 在 s 中出现的末位置
	lastIndex := strings.LastIndex(str, "世界")
	fmt.Println("‘世界’最后一次出现的位置:", lastIndex)

	// strings.IndexRune(s, rune) int ：非ASCII码字符在 s 中出现的首位置
	indexRune := strings.IndexRune(str, '世')
	fmt.Println("‘世’首次出现的位置:", indexRune)

	// strings.Replace(s, oldStr, newStr, n) string ：将 s 中前 n 个 oldStr 替换为 newStr，返回新字符串
	replaced := strings.Replace(str, "世界", "地球", 1)
	fmt.Println("替换后的字符串:", replaced)

	// strings.Count(s, str) int ：返回 str 在 s 中出现的非重叠次数
	count := strings.Count(str, "世界")
	fmt.Println("‘世界’出现的次数:", count) // 输出：2

	// strings.TrimSpace(str) string ：去除 str 开头和结尾的空格
	trimmed := strings.TrimSpace(" 你好，世界！ ")
	fmt.Println("去除空格后的字符串:", trimmed)

	trimmedRegex := strings.Trim("你好，世界！##", "#")
	fmt.Println("去除 '#' 后的字符串:", trimmedRegex)
	trimmedLeft := strings.TrimLeft("你好，世界！##", "#")
	fmt.Println("去除左边 '#' 后的字符串:", trimmedLeft)

	// strings.TrimRight(str, regex) string：去除 str 结尾的 regex
	trimmedRight := strings.TrimRight("你好，世界！##", "#")
	fmt.Println("去除右边 '#' 后的字符串:", trimmedRight)

	// strings.Join([]string, sep) string ：用 sep 当分隔符将字符串数组修改为字符串
	joined := strings.Join([]string{"你好", "世界", "欢迎", "来到"}, ",")
	fmt.Println("连接后的字符串:", joined)

}
