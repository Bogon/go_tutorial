package main

import "fmt"

// 当连续的参数2个或者是3个以上已命名的参数类型相同时，出最后一个类型之外，其他的都可以胜率类型
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(30, 40))

}
