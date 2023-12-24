package main

import (
	"encoding/json"
	"fmt"
)

type Class struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func main() {

	var p1 = Class{
		ID:   1,
		Name: "小明",
		Age:  21,
		Sex:  "man",
	}
	josnbyte, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("json err:", err)
	} else {
		json1 := string(josnbyte)
		fmt.Printf("%v", json1)
	}

}
