package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	// "unicode/utf16"
)


func cat(r *bufio.Reader) {
	for {
		file, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		
		fmt.Fprintf(os.Stdout, "%s", file)
	}
}

func main() {
	// 解析命令行参数
	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())

	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		file, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Println(err)	
			continue
		}
		defer file.Close()
		cat(bufio.NewReader(file))
	}
}
