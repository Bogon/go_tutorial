package main

import "fmt"

/*
变量可以包含初始值，每个变量对应一个
如果变量未被初始化，需要指明变量的类型
如果初始化已存在，则可以省略类型， 变量会从初始值中获取类型(类型推断)
*/
var i, j int = 1, 2

func main() {
	var c, java, python = true, false, "no!"
	fmt.Println(i, j, c, java, python)
}
