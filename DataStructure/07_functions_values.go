package main

import (
	"fmt"
	"math"
)

/*
函数值(高阶函数)

函数也是值。他们可以像其他值一样传递。
函数值可以用作函数的参数或返回值。
*/
func compute(fn func(float64, float64) float64) float64 {
	return fn(3,4)
}

/*
函数的闭包

Go 函数可以作为闭包。闭包是一个函数值，他引用了其他函数体之外的变量。该函数可以访问并赋予其引用的变量的值，换句话说，
该函数被这些变量"绑定"在一起。
*/
func adder() func(int) int {
	sun := 0
	return func(x int) int {
		sun += x
		return  sun
	}
}

func main() {

	/// 函数值
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	/// 函数闭包
	pos, neg := adder(), adder()

	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
			)
	}
}