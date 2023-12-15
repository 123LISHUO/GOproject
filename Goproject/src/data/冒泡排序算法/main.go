package main

import (
	"fmt"
	"math/rand"
)
func main(){
	cheshi := make([]int, 0)
	for i := 0; i < 30; i++ {
		cheshi = append(cheshi, rand.Intn(50))
	}
	// fmt.Println(cheshi)
	maopao(cheshi)
}


func maopao(cheshi []int) {


	for x:=0;x<len(cheshi);x++{
		for y:=0;y<len(cheshi)-x-1;y++{
			if cheshi[y] > cheshi[y + 1]{
				cheshi[y],cheshi[y+1] = cheshi[y+1],cheshi[y]
			}
		}
		fmt.Println("中间状态",cheshi)
	}
	fmt.Println("最终状态",cheshi)
}
