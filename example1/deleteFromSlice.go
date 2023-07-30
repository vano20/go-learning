package main

import "fmt"

func main() {
	// // Declare int slice/array
	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// delete 3rd value from slice
	slicedIndex := 3
	arr2 := append(arr1[:slicedIndex-1], arr1[slicedIndex:]...)
	// print slice
	fmt.Println(arr2)
}
