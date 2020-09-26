package main

import "fmt"

func main() {

	/*
	数组

	类型 [n]T 表示拥有 n 个 T 类型的值的数组。
	表达式
		var a [10]int

	会将变量a声明为拥有10个整数的数组。
	数组的长度是类型的一部分，因此数组不能改变大小。这看起来是个限制，不过没关系，Go 提供了更加便利的方式来使用数组。
	*/
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes)
}