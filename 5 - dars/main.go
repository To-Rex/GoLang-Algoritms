package main

import (
	"fmt"
	"time"
)

func main() {
	loop5()
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
}

func loop1() {
	i := 0
	for i < 10 {
		fmt.Print(i)
		i++
	}
}

func loop2() {
	i := 0
	for {
		fmt.Print(i)
		i++
		if i == 10 {
			break
		}
	}
}

func loop3() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println(time.Now().Second())
	}
}

func loop4() {
	mArr := [5]int{1, 2, 3, 4, 5}
	for i, v := range mArr {
		fmt.Println(i, v)
	}


}
func loop5() {
	mMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 8,
		"g": 7,
		"h": 6,
		"i": 5,
		"j": 4,
	}

	for k, v := range mMap {
		fmt.Println(k, v)
	}
}
