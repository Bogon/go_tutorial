package main

import "fmt"

func main() {

	/*
		类型推导

		在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或者 var = 表达式语法）， 变量的类型由右值推导得出。

		当右值声明了类型时，新变量的类型与其相同：

		var i int
		j := i /// j也是一个int

		不过当右边包含未指明类型的数值常量时，新变量的类型就可能是 int， float64 或者 complex128 了，这取决于常量的精度：

		i := 42 // int
		f := 3.14 // float64
		g := 0.86238 + 0.5i // complex128

	*/
	v := 42 /// 修改在这里

	fmt.Println("v is of type %T\n", v)

}
