package main

import (
	"fmt"
	"time"
)

func main() {
	loop3()
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
