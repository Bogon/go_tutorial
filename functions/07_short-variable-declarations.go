package main

import "fmt"

func main() {

	/*
		短变量声明
		在函数中， 简洁的赋值语句 := 可在类型明确的地方代替 var
		函数外的每个语句都必须以关键字开始（var func等）， 因此 := 结构不能在函数外使用
	*/
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
