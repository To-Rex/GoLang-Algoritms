package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	start1()
}

func start() {
	var wg sync.WaitGroup
	wg.Add(2)
	go display("salom", &wg)
	go display("hello", &wg)
	wg.Wait()
}

func start1() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func ()  {
		display1("salom")
		wg.Done()
	}()
	go func ()  {
		display1("hello")
		wg.Done()
	}()
	wg.Wait()
}

func display(input string, wg *sync.WaitGroup) {
	for i := 0; i<=5; i++ {
		fmt.Println(i, input)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func display1(input string) {
	for i := 0; i<=5; i++ {
		fmt.Println(i, input)
		time.Sleep(1 * time.Second)
	}
}
