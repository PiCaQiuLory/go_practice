package main

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
)

func convertToBin(n int) string {
	result := ""
	//省略起始条件
	for ; n >0 ; n/=2 {
		lsb := n/2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	//省略起始条件和递增条件
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever()  {
	//死循环
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(convertToBin(5),convertToBin(13),convertToBin(72387885), convertToBin(0))
	printFile("demo.txt")
}