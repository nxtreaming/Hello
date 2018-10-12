package main

import (
	"fmt"
	"math"
)

//
// there is no '(' and ')' in 'for' or 'if'
// but MUST have '{' and '}' even if the
// statement has only one line
//
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
	//	for {
	//		//infinite loop
	//	}

	return s
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// simple calculation is allowed in 'if' statement
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	//undefined 'v': out of the scope of v
	//v += 1

	return lim
}

func my_sqrt(x float64) float64 {
	var z float64 = 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("loop %d: fake sqrt(%v): %v\n", i+1, x, z)
	}

	return z
}

func my_sqrt2(x float64) float64 {
	var z float64 = x / 2

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("loop %d: fake sqrt(%v): %v\n", i+1, x, z)
	}

	return z
}

func main() {
	x := 100
	fmt.Printf("The sum of from 1 to %d is %d\n", x, sum(x))
	fmt.Printf("The sum of from 1 to %d is %d\n", x, sum2(x))
	fmt.Printf("The sum of from 1 to %d is %d\n", x, sum3(x))

	fmt.Printf("The value is %s, %s\n", sqrt(4), sqrt(-4))

	fmt.Printf("The limit power: limit %v, %v^%v = %v\n", 10, 3, 2, pow(3, 2, 10))
	fmt.Printf("The limit power: limit %v, %v^%v = %v\n", 20, 3, 3, pow(3, 3, 20))

	my_sqrt(10)
	my_sqrt2(10)

	fmt.Printf("The official sqrt(%d) = %v\n", 10, math.Sqrt(10))
}
