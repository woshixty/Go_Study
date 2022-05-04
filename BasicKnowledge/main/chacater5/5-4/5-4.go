package main

import (
	"flag"
	"fmt"
)

/**
5.4 匿名函数---没有函数名字的函数
匿名函数经常用于回调函数、闭包等等
*/

// 遍历切片元素，通过给定函数解决
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

var skillParam = flag.String("skill", "", "skill to perform")

func main() {
	/**
	5.4.1 定义匿名函数
	*/
	// 1 定义时候嗲用匿名函数
	// 2 将匿名函数赋值给变量
	f := func(data int) {
		fmt.Println("hello world", data)
	}
	f(100)

	/**
	5.4.2 匿名函数用做回调
	*/
	visit([]int{1, 3, 2}, f)

	/**
	5.4.3 使用匿名函数实现操作封装
	*/
	flag.Parse()
	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angle fly")
		},
	}
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not fount")
	}
}
