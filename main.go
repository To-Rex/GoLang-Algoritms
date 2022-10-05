package main

import (
	"fmt"
	"time"
)

func main() {
	functions1()
}

func functions() {
	weekday := time.Now().Weekday()
	fmt.Println(weekday)
	switch weekday {
	case 1:
		fmt.Println("Dushanba")
	case 2:
		fmt.Println("Seshanba")
	case 3:
		fmt.Println("Chorshanba")
	case 4:
		fmt.Println("Payshanba")
	case 5:
		fmt.Println("Juma")
	case 6:
		fmt.Println("Shanba")
	case 7:
		fmt.Println("Yakshanba")
	default:
		fmt.Println("Bunday hafta kuni yo'q")
	}
}
func functions1() {
	mArr:=[]string{"Dushanba","Seshanba","Chorshanba","Payshanba","Juma","Shanba","Yakshanba"}
	//for mArr size for loop
	for i:=0;i<len(mArr);i++{
		fmt.Println(mArr[i])
	}
	//for range loop
	for i,v:=range mArr{
		fmt.Println(i+1,v)
	}

}
func functions2() {
	mArr:=[]string{"Dushanba","Seshanba","Chorshanba","Payshanba","Juma","Shanba","Yakshanba"}
	//for mArr size for loop
	for i:=0;i<len(mArr);i++{
		fmt.Println(mArr[i])
	}
	//for range loop
	for i,v:=range mArr{
		fmt.Println(i+1,v)
	}
}
