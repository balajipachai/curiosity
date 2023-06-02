package main

import "fmt"

/**
* Arrays are the most common datastructure that one come across in various programming languages.
* Arrays store homogeneous data types.
* In Go, len() returns the length of the array.
* Arrays can be traversed using a traditional for loop OR
* Arrays can be traversed by using the `range` keyword.
* While using `range`, both `index` & `value` can be returned.
* In Go, arrays have a fixed size.
 */

// The function prints the index and value of each element in an array of 10 integers.
func main() {
	arrays := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for index, value := range arrays {
		fmt.Println("index\t", index, "\tvalue\t", value)
	}
}
