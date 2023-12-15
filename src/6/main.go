package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func main() {

	var m1 = make(map[string]interface{})

	m1["name"] = "小王孩"
	m1["age"] = 18
	m1["merried"] = false
	m1["hobby"] = []interface{}{"唱", "跟", "rap",1,2,3}

	var s1 = Student{
		Name: "小明",
		Age: 18,
	}

	m1["student"] = s1

	li1 :=m1["hobby"].([]interface{})
	li2 :=m1["student"].(Student)

	fmt.Println(li1[3])
	fmt.Println(li2.Name)
	
}
