package main

import (
	"fmt"
	"os"
)

// import required modules

// main function
func main() {

	// print each argument from os
	fmt.Println(os.Args[1:])
	fmt.Println(os.Args[1])
}
