package main

import "fmt"

func main() {

	/*
		for

		Go只有一种循环结构： for 循环
		基本的 for 循环有三部分组成，他们用分号隔开：
			初始化语句：在第一行迭代前执行；
			条件表达语句：在每次迭代前求值；
			后置语句：在每次迭代末尾执行。

		初始化语句通常为依据短变量声明，该变量声明仅在 for 语句的作用域中可见。

		一旦条件表达式的布尔值为 false，循环迭代就会终止。
		注意： 和C， Java， JavaScript之类的语言不同，CG 的 for 语句后面的三个构成部分外没有小括号，大括号 {} 则是必须的。

	*/

	sum := 0

	for i := 0; i < 100; i++ {
		sum += i
	}

	fmt.Println("sum:", sum)

	/*

		初始化语句和后置语句是可选的。
	*/
	sum1 := 1

	for sum1 < 1000 {
		sum1 += sum1
	}

	fmt.Println("sum1:", sum1)

	/*
		for 是 Go 中的 “while”
		此时你可以去掉分号，因为 C 的 while 中的 Go 中叫做 for 。
	*/

	sum3 := 1

	for sum3 < 10 {
		sum3 += sum3
	}

	fmt.Println(sum3)

	/*
		无限循环

		如果省略循环条件，该循环就不会结束，因此无限循环可以写的很紧凑。
	*/

	for {
		/// TODO……
	}

}
