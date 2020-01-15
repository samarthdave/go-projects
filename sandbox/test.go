package main

import (
	"fmt"
)

func main() {
	var x int
	x = 5
	var y int = 6;
	// Go will infer type if not specified
	sum := x + y

	fmt.Println(sum)

	// control flow
	if sum > 10 {
		fmt.Println("sum is greater than 10")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}