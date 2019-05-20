package main

import "fmt"

//go语言当中array传递不能改变其值，而是拷贝，一般我们用切片

func updateSlice(s []int) {
	s[0] = 100
}

func main(){
	/*arr := [...]int{0,1,2,3,4,5,6,7}

	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:6]
	fmt.Println("s1 = ", s1) //s1 =  [2 3 4 5]
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])
	s2 := arr[:]
	fmt.Println("s2  = ", s2)//s2  =  [0 1 2 3 4 5 6 7]
	fmt.Println("After update slice")
	updateSlice(s1)
	updateSlice(s2)
	fmt.Println("after update slice then s1 = ", s1)//after update slice then s1 =  [100 3 4 5]
	fmt.Println("after update slice then s2 = ", s2)//after update slice then s2 =  [100 1 100 3 4 5 6 7]
	fmt.Println("the original array is arr = ", arr)//the original array is arr =  [100 1 100 3 4 5 6 7]


	fmt.Println("ReSlice")
	s2 = s2[0:2]
	fmt.Println(s2)
	s2 = s2[0:1]
	fmt.Println(s2)*/

	arr := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println("extending slice")
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n ",s1, len(s1), cap(s1))//s1 = [2 3 4 5], len(s1) = 4, cap(s1) = 6
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d\n ",s2, len(s2), cap(s2))//s2 = [5 6], len(s2) = 2, cap(s2) = 3
}