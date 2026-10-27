/*
channel与select
*/
package main

import "fmt"

func fibonaci(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			// 如果成功向c中写入数据，则进行该case处理语句
			x = y
			y = x + y
		case <-quit:
			//// 如果从quit中成功读到数据，则进行该case处理语句
			fmt.Println("quit")
			return
		default:
			// 如果上面都没有成功，则进入default处理流程
			fmt.Println("没招了。。。")
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	//sub go
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("第", i+1, "次，读到的数据为：", <-c)
		}

		quit <- 0
	}()

	// main go
	fibonaci(c, quit)
}
