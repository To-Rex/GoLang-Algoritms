package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(function2(2003))

}

func function(x float64) int {
	return int(math.Pow(x, 2003))
}

func function1(x int) int {
	return x * x
}

func function2(x int) int {
	return x * x * x
}
