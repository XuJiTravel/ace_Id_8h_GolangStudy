/*
结构体标签在json中的应用
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  float32  `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 6.66, []string{"周星驰", "张柏芝"}}

	//编码的过程  结构体 ---> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error: ", err)
		return
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)

	//解码的过程 jsonStr ---> 结构体
	//jsonStr = {"Title":"喜剧之王","year":2000,"price":100,"actors":["周星驰","张柏芝"]}
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error: ", err)
		return
	}

	fmt.Printf("%v\n", myMovie)
}
