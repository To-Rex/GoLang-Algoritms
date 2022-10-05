package main
import (
	"fmt"
)

func main() {
	mArr :=[...][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12},{13,14,15,16},{17,18,19,20},{21,22,23,24}}
	for _,v:=range mArr{
		fmt.Println(v)
		for _,v1:=range v{
			fmt.Println(v1)
		}
	}
	fmt.Println("tomosha tamom")
}