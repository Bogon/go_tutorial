package main

import (
	"fmt"
	"math"
)

/*
类型转换

表达式T(v) 将值 v 转换为类型 T

一些关于数值的转换：

var i int  = 42
var f float64 = float64(i)
var u uint = uint(f)

或者，更加简单的方式：

i := 42
f := float64(i)
u := uint(f)


与 C 不同的是， Go在不同类型之间复制需要显示类型转换。

*/

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))

	var z uint = uint(f)

	fmt.Println(x, y, z)

}
