package main

import "fmt"

/**
3.2 切片（slice）---动态分配大小的连续空间
*/

func main() {
	/**
	3.2.1 从数组或者新的切片生成新的切片
	切片默认指向一段连续的内存空间，可以是数组或切片
	*/
	var a = [3]int{1, 2, 3}
	fmt.Println(a)
	// 1 从指定范围中生成新的切片
	var highReseBuilding [30]int
	for i := 0; i < 30; i++ {
		highReseBuilding[i] = i + 1
	}
	// 区间
	fmt.Println(highReseBuilding[10:15])
	// 中间到尾部的所有元素
	fmt.Println(highReseBuilding[20:])
	// 开头到中间的所有元素
	fmt.Println(highReseBuilding[:2])

	// 2 表示原有切片
	fmt.Println(highReseBuilding[:])

	/**
	3.2.2 声明切片
	*/
	// 声明字符串切片
	var strList []string
	fmt.Println(strList)

	/**
	3.2.3 使用make()函数构造切片
	*/
	c := make([]int, 2)
	d := make([]int, 2, 10)
	fmt.Println(c)
	fmt.Println(d)

	/**
	3.2.4 使用append()函数为切片添加元素
	切片扩容
	*/
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d   cap: %d   pointer: %p\n", len(numbers), cap(numbers), numbers)
	}

	/**
	3.2.4 切片的复制
	*/
	// 设置元素的数量为100
	const elementCount = 1000
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	// 将元素切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// 引用切片数据
	refData := srcData
	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据复制到新的切片空间中去
	copy(copyData, srcData)
	// 修改原始数据的第一个元素
	srcData[0] = 999
	// 打印引用切片的第一个元素
	fmt.Println(refData[0])
	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1])
	// 复制元素原始数据从4到6
	copy(copyData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}

	/**
	3.2.6 从切片中删除数据
	*/
	seq := []string{"a", "b", "c", "d", "e", "f"}
	// 指定删除未知
	index := 2
	// 查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index], seq[index+1:])
	// 将删除前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)
}
