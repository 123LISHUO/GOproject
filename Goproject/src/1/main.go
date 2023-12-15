package main
import(
	"fmt"
	"encoding/json"
)

type Class struct {
	ID	int
	Name string
	Age int
	Sex string
	
}
func main(){
	var str = `{"ID":1,"Name":"小明","Age":21,"Sex":"man"}`

	var s = &Class{}
	err := json.Unmarshal([]byte(str),s)
	if err != nil{
		fmt.Println(err)
	}else{
	   fmt.Printf("%v %T",s,s)
	}
}