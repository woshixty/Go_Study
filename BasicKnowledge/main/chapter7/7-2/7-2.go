package main

import "fmt"

/**
7.2 接口的实现条件
*/

/**
7.2.1 条件1：接口方法与实现接口的类型方法一致
即签名一，包括：方法名、参数列表、返回参数列表
*/
// 定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{}) error
}

// 定义文件结构，用于实现接口
type file struct {
}

// 实现DataWriter中的方法
func (d *file) WriteData(data interface{}) error {
	// 模拟数据写入
	fmt.Println("Write Data：", data)
	return nil
}

func main() {
	// 示例化file
	f := new(file)
	// 声明一个DataWriter接口
	var writer DataWriter
	// 将接口赋值f，也就是file指针类型
	writer = f
	// 使用DataWriter接口进行数据写入
	writer.WriteData("data")
}
