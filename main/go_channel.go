package main

import (
	"fmt"
	"strconv"
	"time"
)

func main(){
	ch := make(chan int)

	go Read(ch)
	go Write(ch)
	time.Sleep(5)
}
func Write(ch chan int) {
	ch <- 10
}
func Read(ch chan int) {
	value := <- ch
	fmt.Println("value:" ,strconv.Itoa(value))
}
