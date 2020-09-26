package main

import (
	"fmt"
	"math/cmplx"
)

/*
基本类型

Go 的基本类型有

bool

string

int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64

byte // uint8 的别名

rune // int32 的别名
	 // 表示一个 Unicode的码点

float32	float64

complex64 complex128

本例中展示了集中类型的变量。同导入语句一样，变量生命也可以 “分组”成一个语句块

int, uint 和 uintptr 在32位系统上通常为32位宽，在64位系统中则为64位宽。当你需要一个整数值时应使用 int类型，
除非你有特殊的理由使用固定大小或无符号的整数类型

*/
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	fmt.Println("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Println("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Println("Type: %T Value: %v\n", z, z)

}
