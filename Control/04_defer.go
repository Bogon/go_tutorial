package main

import "fmt"

func main() {

	/*
	defer

	defer 语句会将函数推迟到外层函数返回值后执行。
	推迟代用的函数其参数会立即求值，但知道外层函数返回前该函数都不会被调用。
	*/
	defer fmt.Println("world")
	fmt.Println("Hello")

	/*
	defer 栈

	推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	*/
	fmt.Println("counting……")

	for i := 0; i < 10; i++ {
		defer  fmt.Println(i)
	}

	fmt.Println("Done.")

}