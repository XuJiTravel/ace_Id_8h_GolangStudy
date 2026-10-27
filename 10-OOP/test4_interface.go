/*
Golang中面向对象多态的实现和基本要素
*/
package main

import "fmt"

// 本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  //H获取动物的类型
}

// 具体的猫类
type Cat struct {
	color string //猫的颜色
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleeping")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 具体的狗类
type Dog struct {
	color string //狗的颜色
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleeping")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep() //多态
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}

func main() {
	/*
		var animal AnimalIF //接口的数据类型，父类指针
		animal = &Cat{"Green"}

		animal.Sleep() //调用的就是Cat的Sleep()方法，多态的现象

		animal = &Dog{"Yellow"}

		animal.Sleep() //调用Dog的Sleep()方法，多态的现象
	*/

	cat := Cat{"Black"}
	dog := Dog{"White"}

	showAnimal(&cat)
	showAnimal(&dog)
}
