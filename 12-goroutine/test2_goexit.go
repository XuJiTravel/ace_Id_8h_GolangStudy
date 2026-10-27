/*
创建goroutine
*/
package main

import (
	"fmt"
	"time"
)

/*
func main() {

	//用go创建承载一个形参为空，返回值为空的一个函数
	go func() {
		defer fmt.Println("A.defer")

		return

		func() {
			defer fmt.Println("B.defer")
			//推出当前的goroutine
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
*/

/*
func main() {

	//用go创建承载一个形参为空，返回值为空的一个函数
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			//退出当前的goroutine
			return
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
*/

/*
func main() {

	//用go创建承载一个形参为空，返回值为空的一个函数
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			//退出当前的goroutine
			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
*/

func main() {

	go func(a int, b int) bool {
		fmt.Println("a = ", a, ", b = ", b)
		return true
	}(10, 20)

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
