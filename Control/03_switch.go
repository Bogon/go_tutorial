package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Println("Go runs on ")

	/*
	switch

	switch 是编写一连串的 if-else 语句的简便写法。它运行第一个值等于表达式的 case 语句。

	Go 的 switch 语句类似于C、 C++、Java、JavaScript 和 PHP 中的，不过 Go 只运行选定的case，而非之后的所有case。
	实际上， Go 自动提供了这些语言中每个 case 后面所需的 break 语句。除非以 fallthrough 语句结束，否则分支会自动终止。
	Go 的另外一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。

	*/
	switch OS := runtime.GOOS; OS {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		/// freebsd, openbsd,
		/// plan9, windows
		fmt.Println("%s.\n", OS)

	}

	fmt.Println("When's Statusday?")


	/*
	switch 的求值顺序

	（例如：
		switch i {
			case 0:
			case f():
		}
	在 i == 0时 f 不会被调用。
	）

	*注意：* Go 练习场中的时间总是从 2009-11-10 23:00:00 UTC 开始，该值的意义留给读者去发现。

	*/
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0 :
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")

	}

	/*
	没有条件的 switch

	没有条件的 switch 同 switch true 一样。
	这种形式能将一长串 if-then-else 写的更清晰。
	*/
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Goog mornning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")


	}

}
