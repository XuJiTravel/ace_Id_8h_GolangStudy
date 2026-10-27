/*
slice切片截取
*/
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4} // len = 4, cap = 4

	// 	取值范围：【0，1）左闭右开
	s1 := s[0:0] // []
	s2 := s[0:1] // [1]
	s3 := s[0:2] // [1 2]
	s4 := s[0:3] // [1 2 3]
	s5 := s[0:4] // [1 2 3 4]

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)

	s4[0] = 100

	fmt.Println(s)
	fmt.Println(s4)

	//copy 可以将底层数组的slice一起进行拷贝
	s6 := make([]int, 4) // s6 = [0, 0, 0, 0]

	//将s中的值 依次拷贝到s6中
	copy(s6, s)
	fmt.Println(s6)
}
