package main

import (
	"fmt"
	"sync"
)

var wg sync.Mutex
type stack struct{
	data []interface{}
}

func (q *stack) push(pushdata interface{}){
	
	wg.Lock()
	defer wg.Unlock()
	q.data = append([]interface{}{pushdata},q.data...)
	// fmt.Println(q.data)

}	


func (q *stack) pop()  (interface{},bool) {
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

	p1 := stack{}
	p1.push(1)
	p1.push(2)
	p1.push(3)
	p1.push(4)

	fmt.Println(p1.pop())
	fmt.Println(p1.pop())
	fmt.Println(p1.pop())
	fmt.Println(p1.pop())
	fmt.Println(p1.pop())
	
}