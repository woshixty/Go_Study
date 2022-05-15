package main

import (
	"fmt"
	"time"
)

func running() {
	var times int
	// 构建一个无限循环
	for {
		times++
		fmt.Println("tick", times)
		// 延迟一秒
		time.Sleep(time.Second)
	}
}

func main() {
	// 并发执行程序
	go running()
	// 接收命令行输入，不做任何事
	var input string
	fmt.Scanln(&input)

	go func() {
		var times int
		for {
			times++
			fmt.Println("tick", times)
			time.Sleep(time.Second)
		}
	}()
}
