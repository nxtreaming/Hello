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

func my_add(x int, y int) int {
	return x + y
}

func main() {
	x := 100
	y := 200

	fmt.Println("The Pi value:", math.Pi)
	fmt.Println("the sum of is:", my_add(x, y))
	fmt.Println("Hello world.")
	fmt.Println((stringutil.Reverse(".dlrow olleH")))
}
