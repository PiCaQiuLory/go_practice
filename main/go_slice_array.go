package main

import "fmt"

//go语言当中array传递不能改变其值，而是拷贝，一般我们用切片

func updateSlice(s []int) {
	s[0] = 100
}

func main(){
	arr := [...]int{0,1,2,3,4,5,6,7}

	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:6]
	fmt.Println("s1 = ", s1)
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])
	s2 := arr[:]
	fmt.Println("s2  = ", s2)
	fmt.Println("After update slice")
	updateSlice(s1)
	updateSlice(s2)
	fmt.Println("after s1 = ", s1)
	fmt.Println("after s2 = ", s2)
	fmt.Println(arr)


	fmt.Println("ReSlice")
	s2 = s2[0:2]
	fmt.Println(s2)
	s2 = s2[0:1]
	fmt.Println(s2)
}