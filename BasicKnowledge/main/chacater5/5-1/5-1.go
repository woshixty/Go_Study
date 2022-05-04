package main

import "fmt"

/**
第五章 函数（function）
Go语言支持普通函数、匿名函数和闭包
*/

/**
5.1 声明函数
*/

/**
5.1.1 普通函数声明方式
*/

func foo(a int, b string) int {
	return 0
}

/**
5.1.2 参数类型简写
*/
func foo2(a, b string) int {
	return 0
}

/**
5.1.3 函数的返回值
Go语言支持多返回值
*/
// 1 同一类型的返回值
func typedTwoValues() (int, int) {
	return 0, 0
}

// 2 带有变量名的返回值
func typedTwoValues1() (a, b int) {
	a = 1
	b = 1
	return
}

/**
5.1.6 示例：函数中的参数传递效果测试
Go语言中传入和返回参数在调用和返回时都使用值传递
需要注意！！！
指针、切片和map等引用型对象指向的内容在参数传递时不会发生复制，而是将指针进行一次复制，也就是说，创建了一次引用
*/

// Data 1 测试数据类型
type Data struct {
	complex  []int
	instance InnerData
	ptr      *InnerData
}

type InnerData struct {
	a int
}

// 2 值传递的函数测试
func passByValue(inFunc Data) Data {
	// 输出参数的成员情况
	fmt.Errorf("inFunc value: %+v\n", inFunc)
	// 打印funcData指针
	fmt.Errorf("inFUnc ptr: %p\n", &inFunc)
	return inFunc
}

func main() {
	// 3 测试流程
	// 准备传入函数的结构
	in := Data{
		complex: []int{1, 2, 3},
		instance: InnerData{
			5,
		},
		ptr: &InnerData{1},
	}
	// 输入结构的成员情况
	fmt.Printf("in value: %+v\n", in)
	// 输入结构的指针地址
	fmt.Printf("in ptr: %p\n", &in)
	// 传入结构体，返回同类型的结构体
	out := passByValue(in)
	// 输出结构的成员情况
	fmt.Printf("out value: %+v\n", out)
	// 输出结构的指针地址
	fmt.Printf("out ptr: %p\n", &out)
}
