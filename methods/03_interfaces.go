package main

import (
	"fmt"
	"math"
)

type Abser interface {
	i_Abs() float64
}

func main() {

	/// 接口
	interfaces()

	/// 接口与隐式实现
	interfaceSatisfyImplicity()
}

/*
接口值

接口也是值。他们也可以像其他值一样传递。
接口值可作为函数的参数和返回值。
在内部，接口值可以看作是包含值和类型的元组：
	(value, type)
接口值保存了一个具体底层类型的具体值。
接口值调用方法时会执行其底层类型的同名方法。
*/

func interfaceValue() {

	v := Vertex_i{1,3}

}

/*
接口与隐式实现

类型通过实现一个接口的所有方法来实现该接口。既然无需专门显式声明，也就没有"implements"关键字。
隐式接口从接口的实现中解耦了定义，这样的接口实现可以出现在任何包中，无需提前准备。
因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。
*/

type I interface {
	M()
}

type T struct {
	S string
}

/// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (t T)M() {
	fmt.Println(t.S)
}

func interfaceSatisfyImplicity() {
	var i I = T{"Hello world!~"}
	i.M()
}


/*
接口

接口类型是由一组方法签名定义的集合。
接口类型变量可以保存任何实现了这些方法的值。

*注意*: 示例代码的 22 行存在一个错误。由于 Abs 方法只为 *Vertex （指针类型）定义，因此 Vertex（值类型）并未实现 Abser。
*/

type MyFloat float64

func (f MyFloat)i_Abs() float64  {
	if f <0 {
		return  float64(-f)
	}
	return float64(f)
}

type Vertex_i struct {
	X, Y float64
}

func (v *Vertex_i) i_Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func interfaces() {
	var a_i Abser
	f := MyFloat(-math.Sqrt2)
	v_i := Vertex_i{3,4}

	a_i = f		/// a_i MyFloat 实现了 Abser
	a_i = &v_i	/// a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	//a_i = v_i

	fmt.Println(a_i.i_Abs())

}
