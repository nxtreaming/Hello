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

	"github.com/nxtreaming/stringutil"
)

func main() {
	//	i := 100
	//	fmt.Println("The value:", i)
	fmt.Println("Hello world.")
	fmt.Println((stringutil.Reverse(".dlrow olleH")))
}
