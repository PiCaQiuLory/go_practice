package main

import "fmt"

type Animal interface {
	Fly()
	Run()
}

type Animal2 interface {
	Fly()
}

type Bird struct {
}

func (bird Bird) Fly()  {
	fmt.Println("bird is flying")
}

func (bird Bird) Run(){
	fmt.Println("bird is running")
}

func main(){
	var animal Animal
	var animal2 Animal2

	bird := new(Bird)

	//对象赋值给接口
	animal = bird
	//接口赋值给接口
	animal2 = animal

	animal2.Fly()
}
