package main

import (
	"fmt"
	"math"
)

type Vertex_s struct {
	X,Y float64
}

/*
指针接收者

你可以为指针接收者声明方法。
这意味着对于某类型 T ，接收者的类型可以用 *T 的文法，。（此外，T 不能是像 *int 这样的指针。）
指针接收者的方法可以修改接收者指向的值。由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。

若使用值接收者，那么 Scale 方法会对原始 Vertex 值的副本进行操作。（对于函数的其他参数也是如此。）Scale方法
必须使用指针接收者来更改 main 函数中声明的 Vertex 的值。

*/
func (v Vertex_s) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex_s) Scale(f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {

	// 指针接收者
	v := Vertex_s{3,4}
	v.Scale(10)
	fmt.Println(v)

	// 方法与指针
	v1 := Vertex_s{3,4}
	fmt.Println(S_Abs(v1))
	S_Scale(v1,10)
	fmt.Println(v1)

	/// 方法与指针重定向
	methodPointerIndirection()

	/// 方法与指针重定向 2
	methodPointerIndirection2()

	/// 选择值或者是指针作为接收者
	methodPointerReceivers()
}

/*
指针与函数

*/
func S_Abs(v Vertex_s) float64 {
	return math.Sqrt(v.X * v.X + v.Y*v.Y)
}

func S_Scale(v Vertex_s, f float64) {
	v.X *= f
	v.Y *= f
}

/*
方法与指针重定向

比较前两个程序，你会发现带指针参数的函数必须接收一个指针。
	var v Vertex
	ScaleFunc(v, 5)  // 编译错误！
	ScaleFunc(&v, 5) // OK

而以指针为接收者的方法被调用时，接收者既能为值，也能为指针：
	var v Vertex
	v.Scale(5)  // OK
	p := &v
	p.Scale(10) // OK
对于语句 v.Scale(5)，即便 v 是值而非指针，带指针接收者的方法也能被直接调用。
也就是说，由于 Scale 方法是一个指针接收者，为了方便起见，Go会将语句 v.Scale(5)
解释为 &v.Scale(5).

*/

func (v *Vertex_s) SS_Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func SS_ScaleFunc(v *Vertex_s, f float64) {
	v.X *= f
	v.Y *= f
}

func methodPointerIndirection() {

	fmt.Println("================================================================================")
	v := Vertex_s{3,4}
	v.SS_Scale(2)
	SS_ScaleFunc(&v, 10)

	p := &Vertex_s{4,3}
	p.SS_Scale(3)
	SS_ScaleFunc(p, 8)

	fmt.Println(v, p)
}

/*
方法与指针重定向 2

接受一个值作为参数的函数必须接受一个指定类型的值。
	var v Vertex
	fmt.Println(AbsFunc(v))  // OK
	fmt.Println(AbsFunc(&v)) // 编译错误！

而以值为接收者的方法被调用时，接收者既能为值也能为指针：
	var v Vertex
	fmt.Println(v.Abs()) // OK
	p := &v
	fmt.Println(p.Abs()) // OK

这种情况下，方法调用 p.Abs() 会被解释为 (*p).Abs()。
*/

func (v Vertex_s)a_Abs() float64 {
	return  math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func a_AbsFunc(v Vertex_s) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func methodPointerIndirection2() {
	fmt.Println("================================================================================")
	v := Vertex_s{3,4}
	fmt.Println(v.a_Abs())
	fmt.Println(a_AbsFunc(v))

	p := &Vertex_s{4,3}
	fmt.Println(p.a_Abs())
	fmt.Println(a_AbsFunc(*p))

}

/*
选择值或者指针为接收者

使用指针接收者原因有二：
	首先，方法能够修改其接收者指向的值。
	其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。

在本例中，Scale 和 Abs 接收者的类型为 *Vertex，即便 Abs 并不需要修改其接收者。
通常来说，所有给定类型的方法都应该有值或指针接收者但不应该二者混用。

*/

func (v *Vertex_s)pr_Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func (v *Vertex_s)pr_Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func methodPointerReceivers() {
	v := &Vertex_s{3,4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.pr_Abs())

	v.pr_Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.pr_Abs())

}

