/*
知识点二： defer 和 return 谁先谁后
*/
package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}

func returnAndDefer() int {

	defer deferFunc()

	return returnFunc()
}

func main() {
	returnAndDefer()
}
