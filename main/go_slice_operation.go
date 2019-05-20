package main

import "fmt"

//Slice的操作
/*
	author: lory
 */
func main() {
	/*arr := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2,10)
	s4 := append(s3,11)
	s5 := append(s4,12)
	fmt.Println("s3, s4, s5 = ", s3, s4, s5)//s3, s4, s5 =  [5 6 10] [5 6 10 11] [5 6 10 11 12]
	fmt.Println(arr)//[0 1 2 3 4 5 6 10]*/

	fmt.Println("creating slice")
	var s []int
	for i := 0; i<100 ; i++ {
		printSlice(s)
		//1.create slice
		s = append(s, 2 * i +1)
	}
	fmt.Println(s)
	//2. create slice
	s1 := []int{2,4,6,8}
	printSlice(s1)
	//3.create slice with len
	s2 := make([]int, 16)
	printSlice(s2)
	//4. create slice with len and cap
	s3 := make([]int, 12, 32)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting element from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Poping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)
	fmt.Println("Poping from back")
	tail := s2[len(s2) -1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)
}
func printSlice(s []int) {
	fmt.Printf("%v,len= %d, cap= %d\n", s, len(s), cap(s))
}