/*
变量的内置pair结构详细说明
*/
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//tty: pair<type:*os.File , value:"11-reflect/dev/tty"文件描述符>
	tty, err := os.OpenFile("11-reflect/dev/tty", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	//r: pair<type: , value:>
	var r io.Reader
	//r: pair<type:*os.File , value:"11-reflect/dev/tty"文件描述符>
	r = tty

	//w: pair<type: , value:>
	var w io.Writer
	//w: pair<type:*os.File , value:"11-reflect/dev/tty"文件描述符>
	w = r.(io.Writer)

	w.Write([]byte("HELLO THIS IS A TEST!!!\n"))
}
