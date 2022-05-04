package main

import (
	"fmt"
)

/**
2.8 类型别名（Type Alias）
*/

func main() {
	/**
	2.8.1 区分类型别名和类型定义
	*/
	// 将NewInt定义为int类型
	type NewInt int

	// 将int取一个别名叫做IntAlias
	type IntAlias = int

	// 将a声明为NewInt类型
	var a NewInt
	// 查看a的类型名
	fmt.Printf("a type: %T\n", a)

	// 将a2声明为IntAlias类型
	var a2 IntAlias
	// 查看a2的类型名
	fmt.Printf("a2 type: %T\n", a2)

}

/**
2.8.2 非本地类型不能定义方法
*/
// 如下就是错误的
// Cannot define new methods on the non-local type 'time.Duration'
//type MyDuration = time.Duration
//
//func (m MyDuration) EasySet(a string) {
//
//}
