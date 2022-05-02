package main

import (
	"fmt"
)

/**
4.4 键值循环（for range）---直接获得对象索引和数据
*/

func main() {
	/**
	4.4.1 遍历数组、切片---获得索引和元素
	*/
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d value:%d\n", key, value)
	}

	/**
	4.4.2 遍历字符串---获得字符
	*/
	var str = "Hello World !"
	for key, value := range str {
		fmt.Printf("key:%d value:0x%x\n", key, value)
	}

	/**
	4.4.3 遍历map---获得map键值对
	*/
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}

	/**
	4.4.4 遍历通道（channel）---接收通道数据
	*/
	c := make(chan int)

	go func() {
		c <- 1
		c <- 2
		c <- 3
	}()

	for v := range c {
		fmt.Println(v)
	}

	/**
	4.4.5 在遍历中选择希望获得变量
	key变为了下划线，即匿名变量
	- 可以理解为占位符
	- 不进行空间分配
	*/
	for _, value := range m {
		fmt.Println(value)
	}
}
