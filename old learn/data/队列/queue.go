package main

import (
	"fmt"
	"sync"
)

var wg sync.Mutex
type queue struct{
	data []interface{}
}

func (q *queue) push(pushdata interface{}){
	
	wg.Lock()
	defer wg.Unlock()
	q.data = append(q.data,pushdata)

}	


func (q *queue) popone()  (interface{},bool) {
	wg.Lock()
	defer wg.Unlock()
	if len(q.data) >0 {
		o := q.data[0]
		q.data = q.data[1:]
		return o, true
	}
	return "nil", false
}


func main(){

	p1 := queue{}
	p1.push(1)
	p1.push(2)
	p1.push(3)
	p1.push(4)

	fmt.Println(p1.popone())
	fmt.Println(p1.popone())
	fmt.Println(p1.popone())
	fmt.Println(p1.popone())
	fmt.Println(p1.popone())
	
}