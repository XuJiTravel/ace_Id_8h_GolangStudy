package main

import (
	"fmt"
)

// const 来定义枚举类型
const (
	//可以在const() 添加一个关键字 iota，每行的iota都会累加1，第一行的iota默认值是0
	BEIJING  = iota //iota = 0
	SHANGHAI        //iota = 1
	SHENZHEN        //iota = 2
)

const (
	a, b = iota + 1, iota + 2 // iota = 0, a = iota + 1 = 1, b = iota + 2 = 2
	c, d                      // iota = 1, c = iota + 1 = 2, d = iota + 2 = 3
	e, f                      // iota = 2, e = iota + 1 = 3, f = iota + 2 = 4

	g, h = iota + 2, iota + 3 // iota = 3  g = iota + 2 = 5, h = iota + 3 = 6
	i, k                      // iota = 4, i = iota + 2 = 6, k = iota + 3 = 7
)

func main() {
	//常量（只读属性）
	const length int = 10

	fmt.Println("length = ", length)

	//length = 100 //常量是不允许修改的。

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)

	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("i = ", i, "k = ", k)

	// iota 只能够配合const() 一起使用， iota 只有在const进行累加效果
	// var a int = iota
}
