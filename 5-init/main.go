/*
package main

import (
	"GolangStudy/5-init/lib1"
	"GolangStudy/5-init/lib2"
	"fmt"
)

func init() {
	fmt.Println("libmain.init()...")
}

func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
	fmt.Println("libmian.main()...")
}
*/

package main

import (
	_ "GolangStudy/5-init/lib1"
	mylib2 "GolangStudy/5-init/lib2"
	//. "GolangStudy/5-init/lib2"
	"fmt"
)

func init() {
	fmt.Println("libmain.init()...")
}

func main() {
	//lib1.Lib1Test()

	//lib2.Lib2Test()
	mylib2.Lib2Test()
	//Lib2Test()

	fmt.Println("libmian.main()...")
}
