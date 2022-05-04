package main

import (
	"fmt"
	"runtime"
)

/**
5.11 宕机恢复（recover）---防止程序的崩溃
*/

/**
5.11.1 让程序在崩溃时能继续执行
*/

// 崩溃时需要传递的上下文信息
type panicContext struct {
	function string
}

// 保护方式允许一个函数
func ProtectRun(entry func()) {
	// 延迟处理的函数
	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error:", err)
		default:
			fmt.Println("error:", err)
		}
	}()
	entry()
}

func main() {
	fmt.Println("运行前")
	ProtectRun(func() {
		fmt.Println("手动宕机前")
		// 使用panic传递上下文
		panic(&panicContext{
			"手动触发panic",
		})
		fmt.Println("手动宕机后")
	})
	// 故意造成空指针访问错误
	ProtectRun(func() {
		fmt.Println("赋值宕机前")
		var a *int
		*a = 1
		fmt.Println("赋值宕机后")
	})
	fmt.Println("运行后")
}
