package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

/*
映射

映射将键映射到值。
映射的零值为nil。nil 映射既没有键，也不能添加键。
make 函数会返回给定类型的映射，并将其初始化备用。
*/
var m map[string]Vertex

func main() {

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])

	/// 映射的文法
	mapLiterals()
	mapLiterals2()

	/// 修改映射
	modifyMap()
}

/*
映射的文法

映射的文法与结构体相似，不过必须有键名。
*/
var m1 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
func mapLiterals() {
	fmt.Println(m1)
}

/*
映射的文法

若定级类型只有一个类型名，你可以在文法的元素中省略它。
*/

var m2 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967,},
	"Google": { 37.42202, -122.08408, },
}
func mapLiterals2() {
	fmt.Println(m2)
}

/*
修改映射

在映射 m 中插入或修改元素：
	m[key] = elem
获取元素：
	elem = m[key]
删除元素：
	delete(m, key)
通过双赋值检测某个键是否存在:
	elem, ok = m[key]

若 key 在 m 中， ok 为true， 否则， ok 为 false。
若 key 不在映射中，那么 elem 是该映射元素类型的零值。

同样的，当映射中读取某个不存在的键时，结果是映射的元素类型的零值。

注意：若 elem 或 ok 还未声明，你可以使用短变量声明：
	elem, ok := m[key]
*/
func modifyMap() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The values:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println(m)

	v, OK := m["Answer"]
	fmt.Println("The value:", v, "Present?", OK)
}