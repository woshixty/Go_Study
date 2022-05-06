package main

import (
	"fmt"
	"math"
)

/**
6.3 初始化结构体的成员变量
*/

type people struct {
	name  string
	child *people
}

/**
6.5.3 示例：二维矢量模拟玩家移动
*/
type Vec2 struct {
	X, Y float32
}

// 使用矢量加上另外一个矢量，生成新矢量
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

// 使用矢量减去另外一个矢量，生成新矢量
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

// 使用矢量乘以另外一个矢量，生成新矢量
func (v Vec2) Scale(s float32) Vec2 {
	return Vec2{
		v.X * s,
		v.Y * s,
	}
}

// 计算两个矢量的距离
func (v Vec2) DistanceTo(other Vec2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

// 返回当前矢量的标准化矢量
func (v Vec2) Normalize() Vec2 {
	mag := v.X*v.X + v.Y*v.Y
	if mag > 0 {
		oneOverMag := 1 / float32(math.Sqrt(float64(mag)))
		return Vec2{v.X * oneOverMag, v.Y * oneOverMag}
	}
	return Vec2{0, 0}
}

func main() {
	/**
	6.3.1 使用键值对初始化结构体
	*/
	relation := &people{
		name: "爷爷",
		child: &people{
			name: "爸爸",
			child: &people{
				name: "我",
			},
		},
	}
	fmt.Println(relation)
}
