package main

import "fmt"

/**
3.1 数组---固定大小的连续空间
*/

func main() {
	/**
	3.1.1 数组的声明
	*/
	var team [3]string
	team[0] = "hello"
	team[1] = "world"
	team[2] = "!"
	fmt.Println(team)

	/**
	3.1.2 初始化数组
	*/
	var team2 = [...]string{"hello", "world", "!"}
	fmt.Println(team2)

	/**
	3.1.3 遍历数组---访问每一个数组元素
	*/
	var team3 = [...]string{"hello", "world", "!"}
	for k, v := range team3 {
		fmt.Println(k, v)
	}
}
