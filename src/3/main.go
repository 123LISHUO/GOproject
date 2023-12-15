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
	// 往结构体里面写数据
	var p1 = Class{

		Title: "class1",
		Student: make([]Student,0,10),

	}
	 for i := 0; i < 10; i++ {

		s := Student{
			Name: fmt.Sprintf("stu%d",i),
			Sex: "man",
		}

		p1.Student = append(p1.Student,s)

	 }
	// 对结构体数据进行序列化

	// 转换
	jsonone , err := json.Marshal(p1)
	
	// 判断
	if err != nil {
		fmt.Println(err)
	}else{
		jsontwo := string(jsonone)
		fmt.Printf("%v %T\n",jsontwo,jsontwo)
	}


}
