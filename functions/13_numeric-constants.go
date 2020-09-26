package main

import "fmt"

const (
	// 将1左移 100 位来创建一个非常大的数字
	// 这个数的二进制是1 后面跟100个0
	Big = 1 << 100
	// 在右移 99 位， 即Small = 1 << 1 或者Small = 2
	Small = Big >> 99
)

/*
数值常量

数值常量是高精度的值。
一个未指定类型的变量由上下文决定其类型。

(int 类型最大可以存储一个 64 位的整合素，有时会更小，根据平台不同有时会更小)

*/

func needInt(x int) int { return x*10 + 1 }
func needFLoat(x float64) float64 {
	return x * 0.1
}
func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFLoat(Small))
	fmt.Println(needFLoat(Big))
}
