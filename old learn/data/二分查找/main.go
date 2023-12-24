package main

import (
	"math/rand"
	"fmt"
)
// 二叉搜索
func binarysearch(str []int,value int)  {

	low := 0
	high := len(str) - 1
	for low <= high {
		mid := (high + low) / 2 
		if value == str[mid]{
			fmt.Println(mid)
			return	
		}else if value < str[mid]{
			high = mid - 1
		}else{
			low = mid + 1	
		}
	}
	fmt.Println("not found")

}

// 冒泡排序
func maopao(str []int)  []int{
	for i:=0;i<len(str);i++{
		for x :=0;x < len(str)-1-i;x++{
			if str[x] > str[x+1]{
				str[x],str[x+1] = str[x+1],str[x]
			}
		}
	}
	return str
}


func main(){
	str := make([]int,0)
	for i:=0;i< 100;i++{
		str = append(str, rand.Intn(100))
	}
	str = maopao(str)
	fmt.Println(str)
	var a int
	fmt.Scanln(&a)
	binarysearch(str,a)
}