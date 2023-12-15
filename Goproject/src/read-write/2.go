package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
)

// func wr(str string){
// 	dirfile,err := os.OpenFile("./1.go",os.O_CREATE|os.O_APPEND|os.O_RDWR ,0666)
// 	if err != nil{
// 		fmt.Println("openfile error:",err)
// 	}
// 	defer dirfile.Close()

// 	filewriter :=bufio.NewWriter(dirfile)
	
// 	for i:=0;i<10;i++{
// 		filewriter.WriteString(str)
// 	}
// 	filewriter.Flush()
// }

func wr(){
	file,err:=os.Create("./1.go")
	if err != nil{
		fmt.Println("create file error:",err)
		return
	}
	defer file.Close()
	for i:=0;i<10;i++{
		var s string = "hello world\n"
		fmt.Fprintf(file,"s",s)
	}
}


func re(){
	refile,err := os.Open("./1.go")
	if err != nil{
		fmt.Println("openfile error:",err)
		return
	}
	defer refile.Close()

	readfile := bufio.NewReader(refile)

	for {
		readline, _ ,err := readfile.ReadLine()
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println("readline error:",err)
		}
		fmt.Println(string(readline))
	}
	// fmt.Println("readfile success")
}
func ma(){
	wr()
	// re()
}