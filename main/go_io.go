package main

import (
	"io/ioutil"
	"fmt"
)

func main(){
	const filename  = "demo.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n",contents)
	}
}
