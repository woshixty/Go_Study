package main

import "fmt"

/**
5.9 处理运行时发生的错误
*/

/**
5.9.4 示例：在解析中使用自定义错误
*/

// 声明一个解析错误
type ParseError struct {
	// 文件名
	Filename string
	// 行号
	Line int
}

// 实现error接口，返回错误描述
func (e *ParseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}

// 创建一些解析错误
func newParseError(filename string, line int) error {
	return &ParseError{Filename: filename, Line: line}
}

func main() {
	var e error
	// 创建一个错误实例
	e = newParseError("main.go", 1)
	// 通过error接口查看错误描述信息
	fmt.Println(e.Error())
	// 根据错误接口的具体类型，获取详细的错误信息
	switch detail := e.(type) {
	case *ParseError:
		fmt.Printf("Filename:%s  Line:%d\n", detail.Filename, detail.Line)
	default:
		fmt.Println("other error")
	}
}
