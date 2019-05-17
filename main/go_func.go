package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

//go语言函数介绍

//返回一个值
func eval(a int, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q,_ :=div(a , b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

//函数多返回值
func div(a int, b int) ( q, r int) {
	return a / b, a % b
}

//函数作为参数
func apply(op func(int, int) int, a int, b int) int {
	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("Calling function %s with args(%d, %d)\n", opName, a, b)
	return op(a, b)
}

//可变参数
func sum(numbers ...int) int {
	s := 0
	for i := range numbers{
		s += numbers[i]
	}
	return s
}

func main() {
	/*result,err := eval(3, 4, "a")
	if err != nil{
		fmt.Println(err)
	}*/
	if result,err := eval(3, 4, "a"); err != nil {
		fmt.Printf("Errors: %s\n", err)
	}else {
		fmt.Println(result)
	}
	fmt.Println(div(13, 4))

	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 5))
	fmt.Println(sum(5,6,7,8))
}