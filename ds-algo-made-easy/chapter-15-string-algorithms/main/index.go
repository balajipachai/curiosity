package main

import (
	"fmt"

	"example.com/stringalgo"
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

func bruteForceStringMatch() {
	fmt.Println(colorRed, "\tString pattern matching by brute force approach", colorReset)
	T := []int{'b', 'a', 'c', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'c', 'a', 'c', 'a'}
	P := []int{'a', 'b', 'a', 'b', 'a', 'c', 'a'}
	fmt.Println(colorYellow + "\tPrinting the T(text) & Pattern(P) array (P is found in T)" + colorReset)
	stringalgo.PrintArray(T, true)
	stringalgo.PrintArray(P, true)
	isFound := stringalgo.StringMatchBruteForce(T, P)
	if isFound >= 0 {
		fmt.Println(colorYellow, "\n\tPattern `P` is found in Text `T` starting from index: ", isFound, colorReset)
	} else {
		fmt.Println(colorYellow, "\n\tPattern `P` is not found in Text `T` ", colorReset)
	}
	printDottedLine()
	P = []int{'a', 'b', 'a', 'b', 'a', 'c', 'a', 'a', 'a'}
	fmt.Println(colorYellow + "\n\tPrinting the T(text) & Pattern(P) array (P is not found in T)" + colorReset)
	stringalgo.PrintArray(T, true)
	stringalgo.PrintArray(P, true)
	isFound = stringalgo.StringMatchBruteForce(T, P)
	if isFound >= 0 {
		fmt.Println(colorYellow, "\n\tPattern `P` is found in Text `T` starting from index: ", isFound, colorReset)
	} else {
		fmt.Println(colorYellow, "\n\tPattern `P` is not found in Text `T` ", colorReset)
	}

	fmt.Println(colorGreen, "\tTime Complexity: O((n – m + 1) × m) ≈ O(n × m). Space Complexity: O(1).", colorReset)
	printDottedLine()
}

func robinKarp() {
	fmt.Println(colorRed, "\tString pattern matching Robin-Karp algorithm", colorReset)
	T := []int{'G', 'E', 'E', 'K', 'S', 'F', 'O', 'R', 'G', 'E', 'E', 'K', 'S'}
	P := []int{'G', 'E', 'E', 'K'}
	fmt.Println(colorYellow + "\tPrinting the T(text) & Pattern(P) array (P is found in T)" + colorReset)
	stringalgo.PrintArray(T, true)
	stringalgo.PrintArray(P, true)
	isFound := stringalgo.RobinKarp(T, P)
	if isFound >= 0 {
		fmt.Println(colorYellow, "\n\tPattern `P` is found in Text `T` starting from index: ", isFound, colorReset)
	} else {
		fmt.Println(colorYellow, "\n\tPattern `P` is not found in Text `T` ", colorReset)
	}
	printDottedLine()

	T = []int{'b', 'a', 'c', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'c', 'a', 'c', 'a'}
	P = []int{'a', 'b', 'a', 'b', 'a', 'c', 'a', 'a', 'a'}
	fmt.Println(colorYellow + "\n\tPrinting the T(text) & Pattern(P) array (P is not found in T)" + colorReset)
	stringalgo.PrintArray(T, true)
	stringalgo.PrintArray(P, true)
	isFound = stringalgo.RobinKarp(T, P)
	if isFound >= 0 {
		fmt.Println(colorYellow, "\n\tPattern `P` is found in Text `T` starting from index: ", isFound, colorReset)
	} else {
		fmt.Println(colorYellow, "\n\tPattern `P` is not found in Text `T` ", colorReset)
	}

	fmt.Println(colorGreen, "\tTime Complexity: Average and best case: O(n+m). Worst case: O(m*n). Space Complexity: O(1).", colorReset)
	printDottedLine()

}

func kmp() {
	fmt.Println(colorRed, "\tString pattern matching KMP algorithm", colorReset)
	T := []int{'G', 'E', 'E', 'K', 'S', 'F', 'O', 'R', 'G', 'E', 'E', 'K', 'S'}
	P := []int{'G', 'E', 'E', 'K'}
	fmt.Println(colorYellow + "\tPrinting the T(text) & Pattern(P) array (P is found in T)" + colorReset)
	stringalgo.PrintArray(T, true)
	stringalgo.PrintArray(P, true)
	isFound := stringalgo.KMP(T, P)
	if isFound >= 0 {
		fmt.Println(colorYellow, "\n\tPattern `P` is found in Text `T` starting from index: ", isFound, colorReset)
	} else {
		fmt.Println(colorYellow, "\n\tPattern `P` is not found in Text `T` ", colorReset)
	}
	printDottedLine()

	T = []int{'b', 'a', 'c', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'c', 'a', 'c', 'a'}
	P = []int{'a', 'b', 'a', 'b', 'a', 'c', 'a', 'c', 'a'}
	fmt.Println(colorYellow + "\n\tPrinting the T(text) & Pattern(P) array (P is not found in T)" + colorReset)
	stringalgo.PrintArray(T, true)
	stringalgo.PrintArray(P, true)
	isFound = stringalgo.KMP(T, P)

	if isFound >= 0 {
		fmt.Println(colorYellow, "\n\tPattern `P` is found in Text `T` starting from index: ", isFound, colorReset)
	} else {

		fmt.Println(colorYellow, "\n\tPattern `P` is not found in Text `T` ", colorReset)
	}

	fmt.Println(colorGreen, "\tTime Complexity: O(m + n), where m is the length of the pattern and n is the length of the text to be searched. Space Complexity: O(m).", colorReset)
	printDottedLine()

}

func boyerMoore() {
	fmt.Println(colorRed, "\tString pattern matching Boyer-Moore algorithm", colorReset)
	T := []int{'G', 'E', 'E', 'K', 'S', 'F', 'O', 'R', 'G', 'E', 'E', 'K', 'S'}
	P := []int{'G', 'E', 'E', 'K'}
	fmt.Println(colorYellow + "\tPrinting the T(text) & Pattern(P) array (P is found in T)" + colorReset)
	stringalgo.PrintArray(T, true)
	stringalgo.PrintArray(P, true)
	stringalgo.BoyerMoore(T, P)
	printDottedLine()

	T = []int{'A', 'B', 'A', 'A', 'A', 'B', 'C', 'D'}
	P = []int{'A', 'B', 'C'}
	fmt.Println(colorYellow + "\n\tPrinting the T(text) & Pattern(P) array (P is not found in T)" + colorReset)
	stringalgo.PrintArray(T, true)
	stringalgo.PrintArray(P, true)
	stringalgo.BoyerMoore(T, P)

	fmt.Println(colorGreen, "\tTime Complexity: O(n x m). Space Complexity: O(1).", colorReset)
	printDottedLine()
}

func fundamentals() {
	bruteForceStringMatch()
	robinKarp()
	kmp()
	boyerMoore()
}

func main() {
	fundamentals()
}
