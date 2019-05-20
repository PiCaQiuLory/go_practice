package main

import "fmt"

func main() {
	m := map[string]string{
		"name":"lory",
		"course": "golang",
	}
	fmt.Println(m)//map[course:golang name:lory]

	m2 := make(map[string]int)
	fmt.Println(m2)//map[]

	fmt.Println("Traversing map")
	for k,v := range m{
		fmt.Println(k,v)
	}

	fmt.Println("Getting values")
	courseName,ok := m["course"]
	fmt.Println(courseName,ok)
	if causeName,ok := m["cause"]; ok{
		fmt.Println(causeName)//不存在的go会拿到zero value 这里是空
	}else {
		fmt.Println("不存在")
	}

	fmt.Println("Deleting values")
	delete(m, "name")
	name,ok := m["name"]
	fmt.Println(name, ok)

}
