package main

import (
	"fmt"
)

func printer(c chan int) {
	// 开始无限循环等待数据
	for {
		// 从channel中获取一个数据
		data := <-c
		// 将0视为结束
		if data == 0 {
			break
		}
		// 打印数据
		fmt.Println(data)
	}
	// 通知main函数已经循环结束
	c <- 0
}

func main() {
	// 创建一个channel
	c := make(chan int)
	// 并发执行printer，传入channel
	go printer(c)
	for i := 1; i <= 10; i++ {
		// 通过channel将数据传给printer
		c <- i
	}
	// 通知并发的printer结束循环等待
	c <- 0
	// 等待printer结束
	<-c

	ch := make(chan int)
	// 声明一个只能发送的通道类型，并赋值未ch
	var chSendOnly chan<- int = ch
	// 声明一个只能接受的通道类型
	var chRecvOnly <-chan int = ch
	chSendOnly <- 0
	<-chRecvOnly
}
