/*
map的三种声明定义方式
*/
package main

import "fmt"

func main() {
	//===> 第一种声明方式

	//声明myMap1是一种map类型，key是string,value是string
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map")
	}

	//在使用map前，需要先使用make给map分配数据空间
	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "Java"
	myMap1["two"] = "Python"
	myMap1["three"] = "Golang"
	myMap1["four"] = "C++"
	myMap1["five"] = "PHP"

	fmt.Println(myMap1)

	//===> 第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "Java"
	myMap2[2] = "Python"
	myMap2[3] = "Golang"
	myMap2[4] = "C++"
	myMap2[5] = "PHP"

	fmt.Println(myMap2)

	//===> 第三种声明方式
	myMap3 := map[string]string{
		"one":   "Java",
		"two":   "Python",
		"three": "Golang",
		"four":  "C++",
		"five":  "PHP",
	}

	fmt.Println(myMap3)
}
