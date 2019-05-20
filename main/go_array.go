package main

import "fmt"
//go 数组

func updateArray(s [5]int){
	s[0] = 200
}

func main() {
	var arr1 []int//等价于 var arr1 [0]int
	arr2 := [3]int{1,3,4}
	arr3 := [...]int{1,3,45,677,889}//编译器自动推断数组大小
	var grid [4][5]int//4行5列数组 [[0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0] [0 0 0 0 0]]
	arr1 = append(arr1, 1)
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	for i,v := range arr3 {
		fmt.Println(i, v)
	}
	fmt.Println("After update array")
	updateArray(arr3)
	//updateArray(arr1)//编译错误
	fmt.Println(arr3)//[1 3 45 677 889]
}
