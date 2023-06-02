package main

/**
* Slices in Go can be thought of as dynamic arrays.
* Using len() & cap() functions we get the length and capacity of slices.
* `_` is the blank identifier, when you don't want to use some returned values make use of underscore.
* Difference between len() & cap()
*	len() tells you how many elements are currently stored in the slice.
*	cap() tells you the maximum number of elements the slice can hold without reallocation.
*
 */

import "fmt"

func appendToSlice(originalSlice *[]string, name string) {
	*originalSlice = append(*originalSlice, name)
}

func twoDSlices() {
	rows := 3
	cols := 5

	twoDSlice := make([][]int, rows)

	for i := range twoDSlice {
		twoDSlice[i] = make([]int, cols)
	}

	twoDSlice[0][4] = 25896
	twoDSlice[1][3] = 457889
	twoDSlice[2][2] = 23364

	fmt.Println("Two dimensional slice is:")
	fmt.Println(twoDSlice)
}

func main() {
	nameless := []string{"Balaji", "Anjali", "Akash", "Ketan", "Ajit", "Aishwarya", "Rupali", "Sacchu", "Snehal", "Vishakha", "Vaishali"}

	fmt.Println("Capacity of nameless: ", cap(nameless))
	fmt.Println("Length of nameless: ", len(nameless))

	fmt.Println("Here is the nameless gang")
	fmt.Println(nameless)

	// You can append value to a slice at runtime
	appendToSlice(&nameless, "Akansha")
	fmt.Println("After append...")
	fmt.Println(nameless)

	twoDSlices()

}
