package main

import (
	"fmt"
	"strings"
)

/**
2.6 字符串的应用
*/

func main() {
	/**
	2.6.1 计算字符串长度
	*/
	var tip1 = "genji is a ninja"
	fmt.Println(len(tip1))

	var tip2 = "忍者"
	fmt.Println(len(tip2))

	/**
	2.6.2 遍历字符串---获取每一个字符串元素
	*/
	// 1 遍历每一个ASCII字符
	// := 左边的变量必须是没有定义过的，否则会报错
	var theme = "狙击 start"
	for i := 0; i < len(theme); i++ {
		fmt.Printf("Ascii: %c %d\n", theme[i], theme[i])
	}

	// 2 按照Unicode字符遍历字符串
	for _, str := range theme {
		fmt.Printf("Unicode: %c %d\n", str, str)
	}

	/**
	2.6.3 获取字符串某一段字符
	获取某一段字符的字串
	*/
	tracer := "死神来了，死神来了bye bye"
	comma := strings.Index(tracer, "，")

	pos := strings.Index(tracer[comma:], "死神")
	fmt.Println(comma, pos, tracer[comma+pos:])
}
