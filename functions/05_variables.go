package main

import "fmt"

/*
var 关键字用于声明变量列表，跟函数的参数列表一样，类型在最后
var 语句可以出现在包或函数级别
*/
var c, python, java bool

func main() {
	var i int

	fmt.Println(i, c, python, java)
}
