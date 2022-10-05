package main

import (
	"fmt"
	"time"
)

func main() {
	go display("salom")
	go display("hello")
	fmt.Scanln()
}

func display(input string) {
	for i := 0; true; i++ {
		fmt.Println(i, input)
		time.Sleep(1 * time.Second)
	
	}
}
