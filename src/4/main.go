package main

import (
	"encoding/json"
	"fmt"
)

type Class struct {
	Title   string `json:"title"`
	Student []Student
}

type Student struct {

	Name string `json:"name"`
	Sex string `json:"age"`

}

func main() {

	// json 数据
	jsontwo := `{"title":"class1","Student":[{"name":"stu0","age":"man"},{"name":"stu1","age":"man"},{"name":"stu2","age":"man"},{"name":"stu3","age":"man"},{"name":"stu4","age":"man"},{"name":"stu5","age":"man"},{"name":"stu6","age":"man"},{"name":"stu7","age":"man"},{"name":"stu8","age":"man"},{"name":"stu9","age":"man"}]}`

	// 反序列化
	var p1 = &Class{}
	err := json.Unmarshal([]byte(jsontwo), p1)

	// 判断
	if err != nil {
		fmt.Println(err)
	}else{

		fmt.Printf("%#v",p1)
		fmt.Println(p1.Title)
	}

}