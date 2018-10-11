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
	"math/cmplx"

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

func basic_type_test() int {
	//
	// basic types:
	// -------------------
	// bool
	// string
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte ---> alias of uint8
	// rune ---> alias of int32
	//      ---> uniode code point
	// float32 float64
	// complex64 complex128
	//

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-10 + 12i)
	)

	fmt.Println("-----Basic Type demo-----")

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	return 0
}

func type_cast_test() int {
	i := 100
	// the type cast MUST be explictly
	var f float64 = float64(i*100) * 10.1

	//  cannot use i (type int) as type float32 in assignment
	//var x float32 = i

	fmt.Printf("the int is %v, the float64 is %v\n", i, f)

	return 0
}

func main() {
	//
	// := is the simple version of : var x int = 100
	// only available in the local context of a function
	//
	x := 100
	y := 200

	z := my_add(x, y)

	//
	// variables MUST be initialized
	//a, b int;
	//

	// b is 0 by default
	var a, b int
	a = 1
	//b = 2

	// var is ignored, because the default value declares the type
	int_a, float_b, string_c := 3, 3.14, " Pi"

	fmt.Print(int_a, float_b, string_c, "\n")

	//
	// if the variable has explict default value,
	// we can ignore the variable's type
	//
	var answer = "No!"
	fmt.Println("The answer is", answer)

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

	basic_type_test()
	type_cast_test()
}
