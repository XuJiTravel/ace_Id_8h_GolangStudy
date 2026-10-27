/*
interface空接口万能类型与类型断言机制
*/
package main

import "fmt"

// interface{}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is Called...")
	fmt.Println(arg)

	//interface{} 该如何区分 此时引用的底层数据结构类型到底是什么？

	//给 interface{} 提供“类型断言”的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Println("arg id string type, value = ", value)
		fmt.Printf("value type = %T\n", value)
	}

	fmt.Println("=====================")
}

type Book5 struct {
	title  string
	author string
}

func main() {
	book := Book5{title: "Golang", author: "xvchaoxi"}

	myFunc(book)
	myFunc(100)
	myFunc("abcdefg")
	myFunc(3.1415926)
	myFunc(true)
}
