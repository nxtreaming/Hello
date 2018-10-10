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

// return named-values;
func string_swap2(str1, str2 string) (s1, s2 string) {
	//
	// the implict defintion in the function context
	// so we can not redefine them as by 's1 :=' and 's1 :='
	//
	s1 = fmt.Sprintf("Original string is: %s, %s", str1, str2)
	s2 = fmt.Sprintf("Exchanged string is: %s, %s", str2, str1)

	// MUST be empty
	return
}

//global var: default false(0)
var c, python, java bool

func main() {
	x := 100
	y := 200
	z := my_add(x, y)

	//
	// variables MUST be initialized
	//a, b int;
	//

	var a, b int
	a = 1
	//b = 2

	fmt.Println(a, b, c, python, java)

	str1, str2 := string_swap("me", "Exchange")
	str_a, str_b := string_swap2("github.com", "gitlab.com")

	fmt.Println("The Pi value:", math.Pi)
	fmt.Println("the sum of is:", z)
	fmt.Println(str1, str2)
	fmt.Println(str_a)
	fmt.Println(str_b)
	fmt.Println("Hello world.")
	fmt.Println((stringutil.Reverse(".dlrow olleH")))
}
