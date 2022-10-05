package main

import (
	"fmt"
)

func main() {
	maps1()
}

func maps() {
	//create map
	var map1 map[string]int
	map1 = make(map[string]int)
	map1["a"] = 1
	map1["b"] = 2
	map1["c"] = 3
	map1["d"] = 4
	map1["e"] = 5
	map1["f"] = 6
	map1["g"] = 7
	map1["h"] = 8

	fmt.Println(map1)
}

func maps1() {
	//create map
	map1 := make(map[string]int)
	map1["a"] = 1
	map1["b"] = 2
	map1["c"] = 3
	map1["d"] = 4
	map1["e"] = 5
	map1["f"] = 6
	map1["g"] = 7
	map1["h"] = 8

	fmt.Println(map1)
}

func maps2() {
	//create map
	map1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
	}

	fmt.Println(map1)
}

func maps3() {
	//create map
	map1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
	}

	fmt.Println(map1["a"])
}