package main

import (
	"fmt"
)

func mySum(x int, y int) int {
	return x + y
}

func main() {
	// integer addition
	var x int = 5
	var y int = 6

	var sum int = x + y
	shortHand := x + y

	fmt.Println(sum)
	fmt.Println(shortHand)

	// arrays (fixed length)
	var a [5]int
	a[2] = 7

	fmt.Println(a)

	manualArr := [5]int{4, 5, 6, 2, 5}
	fmt.Println(manualArr)

	// slices (non-fixed length arrays)
	bSlice := []int{5, 8, 1, 3, 6}
	// now use built in append method
	bSlice = append(bSlice, 13)

	// maps (like dict in Python)
	// string is key, int is value
	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 4
	vertices["dodecagon"] = 6

	fmt.Println(vertices)

	fmt.Println("======================")
	// for loops
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop --> a for loop with some changes
	g := 0
	for g < 5 {
		fmt.Println("hello!")
		g++
	}

	// loop through an array using range
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	// same with a map
	for key, value := range vertices {
		fmt.Println("key", key, "value", value)
	}

	fmt.Println("======================")

	fmt.println(mySum(4, 5))

}
