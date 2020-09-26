package main

import "fmt"

type Vertes struct {
	X int
	Y int
}

var (
	v2 = Vertes{1,2}	// 创建一个 Vertexs 类型的 结构体
	v3 = Vertes{X:1}			// Y:0 被隐式地赋予
	v4 = Vertes{}				// X:0 Y:0
	p1 = &Vertes{1,2}	// 创建一个 *Vertexs 类型的结构体(指针)
)

func main() {

	/*
	结构体

	一个结构体（structs）就是一组字段(field)。
	*/
	fmt.Println(Vertes{1,2})

	/*
	结构体字段

	结构体字段使用点号来访问。
	*/
	v := Vertes{1,2}
	v.X = 4
	fmt.Println(v.X)

	/*
	结构体指针

	结构体字段可以通过结构体指针来访问。
	如果我们有一个指向结构的指针 p，那么可以通过 (*p).X 来访问其字段 X。 不过这么写太啰嗦了，所以语言也允许我们使用
	隐式间接引用，直接写p.X 就可以。
	*/
	v1 := Vertes{1,2}
	p := &v1
	p.X = 1e9
	fmt.Println(v1)

	/*
	结构体文法

	结构体文法通过直接列出字段的值来新分配一个结构体。
	使用 Name： 语法可以仅列出部分字段。(字段名的顺序无关)
	特殊的前缀 & 返回一个指向结构体的指针。
	*/
	fmt.Println(v2, v3, v4, p1)
}
