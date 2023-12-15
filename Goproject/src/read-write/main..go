package main

import (
	// "fmt"
	// "io"
	// "os"
)

// func cat(dirname string){

// 	defer func(){
// 		var err =recover()
// 		if err !=nil{
// 			return	
// 		}
// 	}()

// 	catfile,err1:=os.Open(dirname)
// 	if err1!=nil{
// 		fmt.Println("open file err:",err1)
// 	}
// 	defer catfile.Close()

// 	var buf = make([]byte,128)
// 	var bufall []byte
	
// 	for {
// 		n,err2 := catfile.Read(buf)

// 		if err2 == io.EOF{
// 			break
// 		}

// 		if err2 !=nil{
// 			fmt.Println("read file err:",err2)
// 			return
// 		}

// 		bufall = append(bufall, buf[:n]...)
// 	}

// 	fmt.Println(string(bufall))
// }

func mai(){
	// 打开文件
	// fileone ,err:=os.Open("./main.go")

	// if err !=nil{
	// 	fmt.Println(err)
	// 	return 
	// }

	// defer fileone.Close()
	
	// 写文件
	//  创建文件
	// fileone,err1 := os.Create("./1.go")
	// if err1 != nil{
	// 	fmt.Println(err1)
	// 	return
	// }
	// defer fileone.Close()

	// n ,err2 :=fileone.Write([]byte("package main\nimport(\"fmt\")\nfunc main(){\nfmt.Println(\"hello world\")\n}\n"))
	// fmt.Println(n)
	// if err2 != nil{
	// 	fmt.Println(err2)
	// 	return
	// }

	// 读文件
	//  fileone ,err:=os.Open("./1.go")

	//  if err !=nil{
	// 	fmt.Println(err)
	// 	return
	//  }

	//  defer fileone.Close()

	//  var buf [128]byte
	//  var bufall []byte

	// for {
	// 	n1,err := fileone.Read(buf[:])
	// 	if err == io.EOF{

	// 		break

	// 	}

	// 	if err !=nil{
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	// fmt.Println(n1)
	// 	// fmt.Println(string(buf))
	// 	bufall = append(bufall, buf[:n1]...)
	// }	

	// fmt.Println(string(bufall))

	// 拷贝文件
	// 源文件
	// srcfile,err1 := os.Open("./1.go")
	// if err1 != nil{
	// 	fmt.Println(err1)
	// }
	// defer srcfile.Close()

	// // 复制目标文件
	// dirfile ,err2 := os.Create("./2.go")

	// if err2 != nil{
	// 	fmt.Println(err2)
	// }
	// defer dirfile.Close()

	// buf := make([]byte,128)
	// bufall := make([]byte,0)

	// for {

	// 	n,err3 := srcfile.Read(buf)
	// 	if err3 == io.EOF{
	// 		break
	// 	}
	// 	if err3 !=nil{
	// 		fmt.Println(err3)
	// 		return	
	// 	}

	// 	bufall = append(bufall, buf[:n]...)

	// }

	// dirfile.Write(bufall)
	// fmt.Println("拷行成功")

	// 制作linux命令中cat
	// cat("./1.go")

}

