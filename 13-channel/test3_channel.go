/*
channel的关闭特点
*/
package main

import "fmt"

/*
func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		//close可以关闭一个channel
		//close(c)
	}()

	for {
		//ok如果为true表示channel没有关闭，如果为false表示channel已经关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("Main Finished...")
}
*/

/*
func main() {
	c := make(chan int)

	//向一个已经关闭的channel再次发送数据时会发生panic错误：panic: send on closed channel
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			//close可以关闭一个channel
			close(c)
		}
	}()

	for {
		//ok如果为true表示channel没有关闭，如果为false表示channel已经关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("Main Finished...")
}
*/

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		//close可以关闭一个channel
		close(c)
	}()

	for {
		//ok如果为true表示channel没有关闭，如果为false表示channel已经关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("Main Finished...")
}
