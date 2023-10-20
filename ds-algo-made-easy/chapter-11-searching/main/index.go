package main

import (
	"fmt"

	"example.com/searching"
)

const (
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorReset  = "\033[0m"
	colorCyan   = "\033[36m"
	colorGreen  = "\033[32m"
	colorBlue   = "\033[34m"
)

func printDottedLine() {
	fmt.Println(colorBlue + "---------------------------------------------------------------------------------------------------------------------------------------------------------------" + colorReset)
}

func unorderedLinearSearch() {
	fmt.Println(colorCyan + "\t************UNORDERED LINEAR SEARCH************" + colorReset)
	array := []int{12, 34, 54, 2, 3}
	fmt.Println(colorYellow + "\tPrinting the array:" + colorReset)
	searching.PrintArray(array)
	fmt.Println(colorYellow + "\tSearch 258 in the array:" + colorReset)
	data := 258
	result := searching.UnorderedLinearSearch(array, data)
	if result < 0 {
		fmt.Println(colorCyan, "\t", data, " is not found", colorReset)
	} else {
		fmt.Println(colorCyan, "\t", data, " is found at index ", result, colorReset)
	}

	fmt.Println(colorYellow + "\tSearch 34 in the array:" + colorReset)
	data = 34
	result = searching.UnorderedLinearSearch(array, data)
	if result < 0 {
		fmt.Println(colorCyan, "\t", data, " is not found", colorReset)
	} else {
		fmt.Println(colorCyan, "\t", data, " is found at index ", result, colorReset)
	}

	fmt.Println(colorRed, "\t===> Time complexity: O(n), in the worst case we need to scan the complete array.\n\t===> Space complexity: O(1).", colorReset)
	printDottedLine()
}

func sortedOrOrderedLinearSearch() {
	fmt.Println(colorCyan + "\t************SORTED/ORDERED LINEAR SEARCH************" + colorReset)
	array := []int{12, 34, 54, 2, 3, 504, 84, 69}
	fmt.Println(colorYellow + "\tPrinting the array:" + colorReset)
	searching.PrintArray(array)
	fmt.Println(colorYellow + "\tSearch 584 in the array:" + colorReset)
	data := 584
	result := searching.SortedOrOrderedLinearSearch(array, data)
	if result < 0 {
		fmt.Println(colorCyan, "\t", data, " is not found", colorReset)
	} else {
		fmt.Println(colorCyan, "\t", data, " is found at index ", result, colorReset)
	}

	fmt.Println(colorYellow + "\tSearch 54 in the array:" + colorReset)
	data = 54
	result = searching.SortedOrOrderedLinearSearch(array, data)
	if result < 0 {
		fmt.Println(colorCyan, "\t", data, " is not found", colorReset)
	} else {
		fmt.Println(colorCyan, "\t", data, " is found at index ", result, colorReset)
	}

	fmt.Println(colorRed, "\t===> Time complexity: O(n), in the worst case we need to scan the complete array.\n\t===> Space complexity: O(1).", colorReset)
	printDottedLine()
}

func binarySearchIterative() {
	fmt.Println(colorCyan + "\t************BINARY SEARCH ITERATIVE************" + colorReset)
	array := []int{2, 3, 12, 34, 54, 84, 69, 504}
	fmt.Println(colorYellow + "\tPrinting the array:" + colorReset)
	searching.PrintArray(array)
	fmt.Println(colorYellow + "\tSearch 14 in the array:" + colorReset)
	data := 14
	result := searching.BinarySearchIterative(array, data)
	if result < 0 {
		fmt.Println(colorCyan, "\t", data, " is not found", colorReset)
	} else {
		fmt.Println(colorCyan, "\t", data, " is found at index ", result, colorReset)
	}

	fmt.Println(colorYellow + "\tSearch 504 in the array:" + colorReset)
	data = 504
	result = searching.BinarySearchIterative(array, data)
	if result < 0 {
		fmt.Println(colorCyan, "\t", data, " is not found", colorReset)
	} else {
		fmt.Println(colorCyan, "\t", data, " is found at index ", result, colorReset)
	}

	fmt.Println(colorRed, "\t===> Time complexity: O(log n).\n\t===> Space complexity: O(1).", colorReset)
	printDottedLine()
}

func binarySearchRecursive() {
	fmt.Println(colorCyan + "\t************BINARY SEARCH RECURSIVE************" + colorReset)
	array := []int{2, 3, 12, 34, 54, 69, 84, 504}
	fmt.Println(colorYellow + "\tPrinting the array:" + colorReset)
	searching.PrintArray(array)
	fmt.Println(colorYellow + "\tSearch 140 in the array:" + colorReset)
	data := 140
	result := searching.BinarySearchRecursive(array, data)
	if result < 0 {
		fmt.Println(colorCyan, "\t", data, " is not found", colorReset)
	} else {
		fmt.Println(colorCyan, "\t", data, " is found at index ", result, colorReset)
	}

	fmt.Println(colorYellow + "\tSearch 69 in the array:" + colorReset)
	data = 69
	result = searching.BinarySearchRecursive(array, data)
	if result < 0 {
		fmt.Println(colorCyan, "\t", data, " is not found", colorReset)
	} else {
		fmt.Println(colorCyan, "\t", data, " is found at index ", result, colorReset)
	}

	fmt.Println(colorRed, "\t===> Time complexity: O(log n).\n\t===> Space complexity: O(1).", colorReset)
	printDottedLine()
}

func interpolationSearch() {
	fmt.Println(colorCyan + "\t************INTERPOLATION SEARCH ************" + colorReset)
	array := []int{2, 3, 12, 34, 54, 69, 84, 504}
	fmt.Println(colorYellow + "\tPrinting the array:" + colorReset)
	searching.PrintArray(array)
	fmt.Println(colorYellow + "\tSearch 140 in the array:" + colorReset)
	data := 140
	result := searching.InterpolationSearch(array, data)
	if result < 0 {
		fmt.Println(colorCyan, "\t", data, " is not found", colorReset)
	} else {
		fmt.Println(colorCyan, "\t", data, " is found at index ", result, colorReset)
	}

	fmt.Println(colorYellow + "\tSearch 18 in the array:" + colorReset)
	data = 504
	result = searching.InterpolationSearch(array, data)
	if result < 0 {
		fmt.Println(colorCyan, "\t", data, " is not found", colorReset)
	} else {
		fmt.Println(colorCyan, "\t", data, " is found at index ", result, colorReset)
	}

	fmt.Println(colorRed, "\t===> Time complexity: O(log(log n)).\n\t===> Space complexity: O(1).", colorReset)
	printDottedLine()
}

func searchingBasics() {
	unorderedLinearSearch()
	sortedOrOrderedLinearSearch()
	binarySearchIterative()
	binarySearchRecursive()
	interpolationSearch()
}

func main() {
	searchingBasics()
}
