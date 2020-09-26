package main

import "fmt"

func main() {

	/*
		零值

		没有明确的初始值变量声明会被赋予他们 零值

		零值是：
			数值类型为 0；
			布尔类型为 false；
			字符类型为 ""(空字符串)。

	*/
	var i int
	var f float64
	var b bool
	var s string

	fmt.Println(i, f, b, s)

}
