package main

import (
	"fmt"
)

func sum(v int) int {
	s := 0
	for i := 1; i <= v; i++ {
		s += i
	}

	return s
}

func sum2(v int) int {
	s := 0
	i := 1
	for ; i <= v; i++ {
		s += i
	}

	return s
}

func sum3(v int) int {
	s := 0
	i := 1
	// same as while loop
	for i <= v {
		s += i
		i++
	}

	return s
}

func main() {
	x := 100
	fmt.Printf("The sum of from 1 to %d is %d\n", x, sum(x))
	fmt.Printf("The sum of from 1 to %d is %d\n", x, sum2(x))
	fmt.Printf("The sum of from 1 to %d is %d\n", x, sum3(x))
}
