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

func main() {
	x := 100
	fmt.Printf("The sum of from 1 to %d is %d\n", x, sum(x))
}
