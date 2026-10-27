/*
创建goroutine
*/
package main

import (
	"fmt"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine
func main() {
	//创建一个go程 去执行newTask()流程
	go newTask()

	//此时主程序只是帮助开启了一个go协程，就继续执行下一行了，接着执行完毕退出
	fmt.Println("main Goroutine exit!!!!")

	/*
	       i := 0
	       for {
	   		i++
	   		fmt.Printf("main Goroutnie : i = %d\n", i)
	   		time.Sleep(1 * time.Second)
	   	}
	*/
}
