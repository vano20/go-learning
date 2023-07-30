package main

import "fmt"

// import required modules

// declare variables and define array content
// declare two array, integer and string
func main() {

	arrInt := [5]int{1, 2, 3, 4, 5}
	arrStr := [5]string{"a", "b", "c", "d", "e"}

	// do it five times
	for i := 0; i != 5; i++ {
		// print the $th value of the intarray and the strarray
		fmt.Println("int of", i, arrInt[i])
		fmt.Println("string of", i, arrStr[i])
	}

	// Print both array
	fmt.Println(arrInt)
	fmt.Println(arrStr)
}
