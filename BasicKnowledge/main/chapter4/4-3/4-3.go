package main

import "fmt"

/**
示例：九九乘法表
*/

func main() {
	// 遍历，决定处理第几行数据
	for y := 1; y <= 9; y++ {
		// 遍历，决定每一行有几列
		for x := 1; x <= y; x++ {
			fmt.Printf("%d * %d = %d", x, y, x*y)
		}
		fmt.Println()
	}
}
