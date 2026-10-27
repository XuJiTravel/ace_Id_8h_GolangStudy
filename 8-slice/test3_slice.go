/*
slice 切片的4种声明定义方式
*/
package main

import "fmt"

func main() {
	//第一种声明方式：声明slice是一个切片，并且初始化，默认值是1，2，3.长度len是3
	//slice := []int{1, 2, 3}

	//第二种声明方式：声明slice是一个切片，但是并没有给slice分配空间
	var slice []int
	//slice1 = make([]int, 3) // 开辟3个空间，默认值是0

	//第三种声明方式：声明slice是一个切片，同时给slice分配空间，3个空间，初始化值是0
	//var slice []int = make([]int, 3)

	//第四种声明方式：声明slice是一个切片，同时给slice分配空间，3个空间，初始化值是0，通过:=推导出slice是一个切片
	//slice := make([]int, 3)

	fmt.Printf("len = %d, slice = %v\n", len(slice), slice)

	if slice == nil {
		fmt.Println("slice 是一个空切片")
	} else {
		fmt.Println("slice 是有空间的")
	}
}
