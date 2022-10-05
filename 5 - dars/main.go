package main

import (
	"fmt"
)

func main() {
	loop()
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
