/*
package main

import "fmt"

func swap(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}

func main() {
	var a int = 10
	var b int = 20

	swap(a, b)

	fmt.Println("a = ", a, " b = ", b)
}
*/

/*
package main

import "fmt"

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}

func main() {
	var a int = 10
	var b = 20

	swap(&a, &b)

	fmt.Println("a = ", a, " b = ", b)
}
*/

package main

import "fmt"

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp
}

func main() {
	var a int = 10
	var b = 20

	swap(&a, &b)

	fmt.Println("a = ", a, " b = ", b)

	var p *int //一级指针

	p = &a
	fmt.Println("&a = ", &a)
	fmt.Println("p = ", p)

	var pp **int //二级指针

	pp = &p

	fmt.Println("&p = ", &p)
	fmt.Println("pp = ", pp)
}
