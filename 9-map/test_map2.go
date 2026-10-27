/*
map的使用方式
*/
package main

import "fmt"

func printMap(cityMap map[string]string) {
	//cityMap 是一个引用传递
	for key, value := range cityMap {
		fmt.Println("key = ", key)
		fmt.Println("value = ", value)
	}
}

func ChangeValue(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	//添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	// 遍历
	printMap(cityMap)

	//删除
	delete(cityMap, "USA")

	// 修改
	cityMap["China"] = "Shanghai"
	ChangeValue(cityMap)

	fmt.Println("=======")

	// 遍历
	// map存储是无序的，遍历 Map 时返回的键值对的顺序是不确定的
	printMap(cityMap)
}
