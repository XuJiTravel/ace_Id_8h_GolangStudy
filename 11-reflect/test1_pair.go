/*
变量的内置pair结构详细说明
*/
package main

import "fmt"

func main() {
	var a string
	//pair<statictype:string, value:"hello">
	a = "hello"

	//pair<type:string, value:"hello">
	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}
