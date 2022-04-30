package main

/**
2.7 常量---恒定不变的值
在编译期间确定的值
*/

import (
	"fmt"
)

type Weapon int

// 常量的声明
const (
	pi = 3.1415926
	e  = 2.718218
)

/**
2.7.1 枚举---一组常量的值
Go中没有枚举，使用常量配合iota来模拟枚举
*/
const (
	Arrow Weapon = iota
	Shuriken
	SniperRifle
	Rifle
	Blower
)

// 还可以使用iota来做一些强大的枚举常量生成器
const (
	FlagNone = 1 << iota
	FlagRed
	FlagGreen
	FlagBlue
)

/**
2.7.2 将枚举值转换为字符串
通过枚举值获得对应的字符串
*/
type ChipType int

const (
	None ChipType = iota
	CPU
	GPU
)

func (c ChipType) String() string {
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

func main() {
	// 输出所有的枚举值
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)

	// 使用枚举类型并赋初值
	var weapon Weapon = Blower
	fmt.Println(weapon)

	fmt.Printf("%d %d %d\n\n", FlagRed, FlagGreen, FlagBlue)
	fmt.Printf("%b %b %b\n\n", FlagRed, FlagGreen, FlagBlue)

	// 输出CPU的值并以整数格式显示
	fmt.Printf("%s %d", CPU, GPU)
}
