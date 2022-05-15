package main

import (
	"errors"
	"fmt"
	"time"
)

// 模拟RPC客户端请求和接收
func RPCClient(ch chan string, req string) (string, error) {
	// 向服务器发送请求
	ch <- req
	// 等待服务器返回
	select {
	// 接收到数据
	case ack := <-ch:
		return ack, nil
	// 超时
	case <-time.After(time.Second):
		return "", errors.New("time Out")
	}
}

// 模拟RPC服务端接收客户端请求和相应
func RPCServer(ch chan string) {
	for {
		// 接收客户端请求
		data := <-ch
		// 打印接受的数据
		fmt.Println("server received :", data)
		time.Sleep(time.Second * 2)
		// 向客户端反馈已收到
		ch <- "roger"
	}
}

func main() {
	ch := make(chan string)
	go RPCServer(ch)
	str, err := RPCClient(ch, "Hello")
	if err != nil {

	}
	go fmt.Println("client received :", str)
	time.Sleep(time.Second * 3)
}
