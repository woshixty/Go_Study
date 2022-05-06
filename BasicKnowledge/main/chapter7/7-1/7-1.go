package main

/**
第七章 接口（interface）
接口本身是调用方和实现方均需遵守的一种协议
*/

/**
7.1.2 接口的常见写法
*/
type Writer interface {
	Write(p []byte) (n int, err error)
}

func main() {

}
