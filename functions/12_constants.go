package main

import "fmt"

const Pi = 3.14

/*
常量


常量的声明和变量的声明类似， 只不过是使用 const 关键字。

常量可以是字符，字符串， 布尔值或者是数值。

常量不能用 := 语法声明。

*/
func main() {

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules", Truth)

}
