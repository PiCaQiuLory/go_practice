package main

import "fmt"
// go语言可以灵活的创建函数，并作为另一个函数的实参
type cb func(int) int

func main(){
	testCallBack(1, callback)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})
}

func testCallBack(x int, f cb) {
	f(x)
}

func callback(x int) int{
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}