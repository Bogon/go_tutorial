package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("====================================【切片】====================================")
	fmt.Println()
	fmt.Println("===============================================================================")
	/*
	切片

	每个数组的大小都是固定的。而切片则为数组元素提供了动态大小的、灵活的视角。在实践中，切片比数组更常用。

	类型 []T 表示一个元素类型为 T 的切片。
	切片通过两个下标来定界，即一个上界和一个下界，二者以冒号分割：
		a[low : high]

	它会选择一个半开取件，包括第一个元素，但排除最后一个元素。
	以下表达式创建了一个切片，它包含 a 中下标从1到3的元素：
		a[1:4]
	*/
	primes := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("====================================【切片就像数组的引用】====================================")

	/*
	切片就像数组的引用
	切片并不存储任何数据，他只是描述了底层数组中的一段。
	更改切片的元素会修改其底层数组中对应的元素。
	与其共享底层数组的切片都会观测到这些修改。
	*/
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:3]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println("a: %v b: %v", a, b)
	fmt.Println(names)

	fmt.Println("====================================【切片文法】====================================")

	/*
	切片文法

	切片文法类似于没有长度的数组文法。
	这是一个数组文法：
		[3]int{1,2,3}

	下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：
	 	[]bool{true, false, true, false}
	*/
	q := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(q)

	r := []bool{true, false, true, true, false}
	fmt.Println(r)

	struct1 := []struct {
		i int
		b bool
	}{
		{2, true},
		{2, false},
		{2, true},
		{2, true},
		{2, true},
	}
	fmt.Println(struct1)

	fmt.Println("====================================【切片默认行为】====================================")
	/// 切片的默认行为
	slicesBounds()

	fmt.Println("====================================【切片的长度和容量】====================================")
	/// 切片的长度和容量
	slicesLenCap()

	fmt.Println("====================================【nil 切片】====================================")
	/// nil 切片
	nilSlices()

	fmt.Println("====================================【使用 make 创建切片】====================================")
	/// 使用 make 创建切片
	makeSlices()

	fmt.Println("====================================【切片的切片】====================================")
	/// 切片的切片
	slicesOfSlices()

	fmt.Println("====================================【向切片追加元素】====================================")
	/// 向切片中追加元素
	appendSlices()
}

/*
切片的默认行为

在进行切片时，你可以使用它的默认行为来忽略边界。
切片的下界的默认值为 0， 上界则是该切片的长度。

对于数组：
	var s [10]int

来说， 以下切片是等价的：
	a[0:10]
	a[:10]
	a[0:]
	a[:]
*/
func slicesBounds() {
	s := []int{1, 2, 3, 4, 5, 6}

	s = s[1: 4]
	fmt.Println(s)

	s = s[:3]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

/*
切片的长度和容量

切片拥有 长度 和 容量。
切片的长度就是它所包含的元素个数。
切片的容量是从它的第一个元素开始，其底层数组元素末尾的个数。
切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。
你可通过重新切片来扩展一个切片，给他提供足够的容量。
*/
func slicesLenCap() {
	s := []int{1,2,3,4,5,6,7,8,9}
	printSlices(s)

	/// 截取切片使其长度为 0
	s  = s[:0]
	printSlices(s)

	/// 扩展其长度
	s = s[:4]
	printSlices(s)

	/// 舍弃前两个值
	s = s[2:4]
	printSlices(s)
}

func printSlices(s []int) {
	fmt.Println("len: %d cap: %d %v \n", len(s), cap(s), s)
}

/*
nil 切片

切片的零值是 nil。
nil 切片的长度和容量为 0 且没有底层数组。
*/
func nilSlices() {
	var s []int
	printSlices(s)
	if s == nil {
		fmt.Println("The slices is nil!")
	}
}

/*
用 make 创建切片

切片可以使用内建函数 make 来创建， 这也就是你创建动态数组的方式。

make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：
	a := make([]int, 5) // len(5) = 5
要指向它的容量，需向 make 传入第三个参数：
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4
*/

func makeSlices() {

	a := make([]int, 5)
	printSlices(a)

	b := make([]int, 0, 5)
	printSlices(b)

	c := b[:2]
	printSlices(c)

	d := c[2:5]
	printSlices(d)
}

/*
切片的切片

切片可包含任何类型的， 甚至包括其他的切片。
*/
func slicesOfSlices() {

	/// 创建一个井字板(经典游戏)
	board := [][]string {
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	/// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i<len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

/*
向切片中追加元素

为切片中追加新的元素是种常用的操作，为此 Go 提供了内建的 append 函数。
	func append(s []T, vs ...T) []T

append 的第一个参数 s 是一个元素的类型为 T 的切片， 其余类型为 T 的值将会追加到该切片的末尾。
append 的结果是一个包含元切片所有元素加上新添加元素的切片。
当 s 的底层数组大小， 不足以容纳所有给定的值时，他就会分配一个更大的数组。返回的切片会指向这个新分配的数组。
*/
func appendSlices() {
	var s []int
	printSlices(s)

	/// 添加一个空切片
	s = append(s, 0)
	printSlices(s)

	/// 这个切片会按需增长
	s = append(s, 1)

	/// 可以一次性添加多个元素
	s = append(s,2, 3, 4, 5)
	printSlices(s)

}