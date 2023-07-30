// Go program to illustrate how to copy an array by value
package main

import "fmt"

func main() {

	// Creating and initializing an array
	// Using shorthand declaration

	shortArr := [3]int{1, 2, 3}

	// Copying the array into new variable
	copiumArr := shortArr

	copiumArr[1] = 1

	// print arrays
	fmt.Println(shortArr)
	fmt.Println(copiumArr)
}
