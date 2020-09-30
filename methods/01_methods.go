package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
方法

Go中没有类。不过你可以为结构体类型定义方法。
方法就是一类带有特殊的 接收者 参数的函数。
方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
*/
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
方法即函数。

记住：方法只是个带接收者参数的函数。
*/
func s_Abs(v Vertex) float64 {
	return  math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {


	v := Vertex{3,4}
	fmt.Println(v.Abs())

	fmt.Println(s_Abs(v))

	/// 非结构类型的方法
	otherMethods()
}

/*
方法（续）

你也可以为非结构体类型声明方法。

你只能为在同一包内定义的类型的接收者声明方法，而不能为其他包内定义的类型(包括 int 之类的内件类型。)
的接收者声明方法。

(注: 就是接收者的类型定义和方法声明必须在一个包内；不能为内建类型声明方法。)
*/
type MineFloat float64

func (f MineFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func otherMethods() {
	f := MineFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
