package main
import(
	"fmt"
	"time"
	"sync"
)


func test(num int ){
	for x:=(num -1 ) * 30000 + 1; x<30000 * num; x++{
		var flag = true

		if x == 1{

			continue

		}else{

			for  y:=2;y<x;y++{

				if x % y == 0{
					flag = false
					break	
				}else{
					flag = true
				}
			}
			// fmt.Println(flag)
			if flag {

			fmt.Println(x,"是质数")

			}
		}
	}
	wg.Done()

}

var wg sync.WaitGroup

func main(){
	start:=time.Now().Unix()
	
	for i:=1;i<=4;i++{

		wg.Add(1)
		go test(i)

	}
	wg.Wait()
	end:=time.Now().Unix()
	println(end-start)
}