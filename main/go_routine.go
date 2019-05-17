package main

import (
	"fmt"
	"time"
)

func test_routine(){
	fmt.Println("routine is running")
}

func main(){
	//启动协程
	go test_routine()
	time.Sleep(1)
}
