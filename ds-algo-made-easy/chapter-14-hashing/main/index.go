package main

import (
	"fmt"

	"example.com/hashing"
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

func two() {
	fmt.Println(colorBlue, "\tProblem 2: Given an array of characters, give an algorithm for removing the duplicates.", colorReset)
	chars := []int{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'g', 'h', 'i', 'j', 'a', 'b', 'c', 'd', 'e', 'f'}
	fmt.Println(colorYellow + "\tPrinting the array:" + colorReset)
	hashing.PrintArray(chars, true)
	withoutDuplicates := hashing.RemoveDuplicates(chars)
	fmt.Println(colorYellow + "\n\tPrinting the array after removing all the duplicates:" + colorReset)
	hashing.PrintArray(withoutDuplicates, true)
	fmt.Println(colorGreen, "\tTime Complexity: O(n^2). Space Complexity: O(1).", colorReset)
	printDottedLine()
}

func three() {
	fmt.Println(colorBlue, "\tProblem 3: Can we find any other idea to solve this problem in better time than O(n2)? Observe that the order of characters in solutions do not matter.", colorReset)
	chars := []int{1, 2, 2, 3, 4, 4, 4, 5, 5}
	fmt.Println(colorYellow + "\tPrinting the array:" + colorReset)
	hashing.PrintArray(chars, false)
	withoutDuplicates := hashing.RemoveDuplicatesWithSorting(chars)
	fmt.Println(colorYellow + "\n\tPrinting the array after removing all the duplicates:" + colorReset)
	hashing.PrintArray(withoutDuplicates, false)
	fmt.Println(colorGreen, "\tTime Complexity: Θ(nlogn). Space Complexity: O(1).", colorReset)
	printDottedLine()
}

func four() {
	fmt.Println(colorBlue, "\tProblem 4: Can we solve this problem in a single pass over given array?", colorReset)
	fmt.Println(colorCyan, "\t\tWe can use hash table to check whether a character is repeating in the given string or not. If the current character is not available in hash table, then insert it into hash table and keep that character in the given string also. If the current character exists in the hash table then skip that character.\n", colorReset)
	chars := []int{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'g', 'h', 'i', 'j', 'a', 'b', 'c', 'd', 'e', 'f'}
	fmt.Println(colorYellow + "\tPrinting the array:" + colorReset)
	hashing.PrintArray(chars, true)
	withoutDuplicates := hashing.RemoveDuplicatesWithSorting(chars)
	fmt.Println(colorYellow + "\n\tPrinting the array after removing all the duplicates:" + colorReset)
	hashing.PrintArray(withoutDuplicates, true)
	fmt.Println(colorGreen, "\tTime Complexity: Θ(n) on average. Space Complexity: O(n).", colorReset)
	printDottedLine()
}

func five() {
	fmt.Println(colorBlue, "\tProblem 5: Given two arrays of unordered numbers, check whether both arrays have the same set of numbers?", colorReset)
	a := []int{2, 5, 6, 8, 10, 11, 12}
	b := []int{2, 5, 5, 8, 10, 5, 6}
	c := []int{2, 5, 6, 8, 10, 11, 12}
	// Check if arrays A & B are same
	fmt.Println(colorYellow + "\tPrinting the array A & B:" + colorReset)
	hashing.PrintArray(a, false)
	hashing.PrintArray(b, false)
	fmt.Println(colorYellow, "\n\tAre arrays `A & B` same: ", hashing.AreArraysSame(a, b), colorReset)
	// Check if arrays A & C are same
	fmt.Println(colorYellow + "\tPrinting the array A & C:" + colorReset)
	hashing.PrintArray(a, false)
	hashing.PrintArray(c, false)
	fmt.Println(colorYellow, "\n\tAre arrays `A & C` same: ", hashing.AreArraysSame(a, c), colorReset)

	fmt.Println(colorGreen, "\tTime Complexity: Θ(n^2) since for each element of A we have to scan B.", colorReset)
	printDottedLine()
}

func six() {
	fmt.Println(colorBlue, "\tProblem 6: Can we improve the time complexity of Problem-5?", colorReset)
	fmt.Println(colorCyan, "\t\tYes. To improve the time complexity, let us assume that we have sorted both the lists. Since the sizes of both arrays are n, we need O(n log n) time for sorting them. After sorting, we just need to scan both the arrays with two pointers and see whether they point to the same element every time, and keep moving the pointers until we reach the end of the arrays.\n", colorReset)

	a := []int{2, 5, 6, 8, 10, 2, 2}
	b := []int{2, 5, 5, 8, 10, 5, 6}
	c := []int{2, 5, 6, 8, 10, 2, 2}

	// Check if arrays A & B are same
	fmt.Println(colorYellow + "\tPrinting the array A & B:" + colorReset)
	hashing.PrintArray(a, false)
	hashing.PrintArray(b, false)
	fmt.Println(colorYellow, "\n\tAre arrays `A & B` same: ", hashing.AreArraysSameUsingSorting(a, b), colorReset)
	// Check if arrays A & C are same
	fmt.Println(colorYellow + "\tPrinting the array A & C:" + colorReset)
	hashing.PrintArray(a, false)
	hashing.PrintArray(c, false)
	fmt.Println(colorYellow, "\n\tAre arrays `A & C` same: ", hashing.AreArraysSameUsingSorting(a, c), colorReset)

	fmt.Println(colorGreen, "\tTime Complexity: Θ(nlogn) since sorting requires nlogn and scanning requires O(n), ultimately making the time complexity to be O(nlogn).", colorReset)
	printDottedLine()
}

func seven() {
	fmt.Println(colorBlue, "\tProblem 7: Can we further improve the time complexity of Problem-5?", colorReset)
	fmt.Println(colorCyan, "\t\tYes, by using a hash table.\n", colorReset)

	a := []int{2, 5, 6, 8, 10, 2, 2}
	b := []int{2, 5, 5, 8, 10, 5, 6}
	c := []int{2, 5, 6, 8, 10, 2, 2}
	d := []int{2, 8, 5, 5, 5, 6, 10}

	// Check if arrays A & B are same
	fmt.Println(colorYellow + "\tPrinting the array A & B:" + colorReset)
	hashing.PrintArray(a, false)
	hashing.PrintArray(b, false)
	fmt.Println(colorYellow, "\n\tAre arrays `A & B` same: ", hashing.AreArraysSameUsingHashTable(a, b), colorReset)
	// Check if arrays A & C are same
	fmt.Println(colorYellow + "\tPrinting the array A & C:" + colorReset)
	hashing.PrintArray(a, false)
	hashing.PrintArray(c, false)
	fmt.Println(colorYellow, "\n\tAre arrays `A & C` same: ", hashing.AreArraysSameUsingHashTable(a, c), colorReset)

	// Check if arrays B & D are same
	fmt.Println(colorYellow + "\tPrinting the array B & D:" + colorReset)
	hashing.PrintArray(b, false)
	hashing.PrintArray(d, false)
	fmt.Println(colorYellow, "\n\tAre arrays `B & D` same: ", hashing.AreArraysSameUsingHashTable(b, d), colorReset)
	fmt.Println(colorGreen, "\tTime Complexity: Θ(n).", colorReset)
	printDottedLine()
}

func problemsAndSolutions() {
	two()
	three()
	four()
	five()
	six()
	seven()
}

func main() {
	problemsAndSolutions()
}
