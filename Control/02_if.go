package main

import (
	"fmt"
	"math"
)

/*
	if

	Go 的 if 语句与 for 循环类似，表达式外无需小括号（），为大括号{}则是必须的。
*/
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

/*
if 的简短语句

同 for 一样， if 语句可以在条件表达式前执行一个简短的语句。

该语句声明的变量作用域仅在 if 之内。

*/

/*
if 和 else

在 if 的简短语句中声明的变量同样可以在任何对应的 else 块中使用。

*/

func pow(x, n, lim float64) float64 {

	if v := math.Pow(x, n); v < lim {
		return  v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	/// 这里开始就不能使用v了
	return  lim
}


func main() {

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 10),
		)

}
