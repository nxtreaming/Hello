//
// Firstly, setup GOROOT and GOPATH environment variables
//
// --------------------------------
// How to compile?
// go build
//
// How to install?
// go install github.com/nxtreaming/Hello
//
// How to remove?
// go clean -i github.com/nxtreaming/Hello
//
// ---------------------------------
//

//
// 'package name' MUST be the first statement
// Executable commands must always use 'package main'.
//
package main

import (
	"fmt"
	"math"

	"github.com/nxtreaming/stringutil"
)

//
// 'int {' MUST be in the same line with 'func'
//
//func my_add(x int, y int) int
//{
//	return x+y
//}

//func my_add(x int, y int) int {
func my_add(x, y int) int {
	return x + y
}

// multi-returned values
func string_swap(str1, str2 string) (string, string) {
	return str2, str1
}

func main() {
	x := 100
	y := 200
	str1, str2 := string_swap("me", "Exchange")

	z := my_add(x, y)

	fmt.Println("The Pi value:", math.Pi)
	fmt.Println("the sum of is:", z)
	fmt.Println(str1, str2)
	fmt.Println("Hello world.")
	fmt.Println((stringutil.Reverse(".dlrow olleH")))
}
