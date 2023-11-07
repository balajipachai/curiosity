package main

import (
	"fmt"

	"example.com/selectionalgorithms"
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

func maxMinUsingSorting() {
	fmt.Println(colorCyan + "\t************MAX/MIN USING SORTING************" + colorReset)
	array := []int{1000, 11, 445, 1, 330, 3000}
	fmt.Println(colorYellow + "\tPrinting the array:" + colorReset)
	selectionalgorithms.PrintArray(array)
	max, min := selectionalgorithms.MaxMinUsingSorting(array)
	fmt.Println(colorCyan, "\tMax, Min using Sorting = ", max, " , ", min, colorReset)
	fmt.Println(colorGreen, "\tTIME COMPLEXITY: O(nlogn)", colorReset)
	printDottedLine()
}

func maxMinUsingLinearSearch() {
	fmt.Println(colorCyan + "\t************MAX/MIN USING LINEAR SEARCH************" + colorReset)
	array := []int{1000, 11, 445, 1, 330, 3000}
	fmt.Println(colorYellow + "\tPrinting the array:" + colorReset)
	selectionalgorithms.PrintArray(array)
	max, min := selectionalgorithms.MaxMinUsingLinearSearch(array)
	fmt.Println(colorCyan, "\tMax, Min using Linear Search = ", max, " , ", min, colorReset)
	fmt.Println(colorGreen, "\tTIME COMPLEXITY: O(n)", colorReset)
	printDottedLine()
}

func maxMinUsingTournamentMethod() {
	fmt.Println(colorCyan + "\t************MAX/MIN USING TOURNAMENT METHOD************" + colorReset)
	array := []int{1000, 11, 445, 1, 330, 3000}
	fmt.Println(colorYellow + "\tPrinting the array:" + colorReset)
	selectionalgorithms.PrintArray(array)
	max, min := selectionalgorithms.MaxMinUsingTournamentMethod(array, 0, len(array)-1)
	fmt.Println(colorCyan, "\tMax, Min using Tournament Method = ", max, " , ", min, colorReset)
	fmt.Println(colorGreen, "\tTIME COMPLEXITY: O(n)", colorReset)
	fmt.Println(colorGreen, "\tTOTAL COMPARISONS: T(n)  = 3n/2 -2", colorReset)
	printDottedLine()
}

func maxMinByComparingPairs() {
	fmt.Println(colorCyan + "\t************MAX/MIN BY COMPARING PAIRS************" + colorReset)
	array := []int{1000, 11, 445, 1, 330, 3000}
	fmt.Println(colorYellow + "\tPrinting the array:" + colorReset)
	selectionalgorithms.PrintArray(array)
	max, min := selectionalgorithms.MaxMinByComparingPairs(array)
	fmt.Println(colorCyan, "\tMax, Min by comparing pairs = ", max, " , ", min, colorReset)
	fmt.Println(colorGreen, "\tTIME COMPLEXITY: O(n)", colorReset)
	printDottedLine()
}

func kthSmallestUsingPartition() {
	fmt.Println(colorCyan + "\t************K-th SMALLEST ELEMENT USING PARTITION************" + colorReset)
	array := []int{1000, 11, 445, 1, 330, 3000}
	k := 5
	fmt.Println(colorYellow + "\tPrinting the array:" + colorReset)
	selectionalgorithms.PrintArray(array)
	fmt.Println(colorYellow, "\tK = ", k, colorReset)
	kthSmallestElement := selectionalgorithms.KthSmallestUsingPartition(array, 0, len(array)-1, k)
	fmt.Println(colorCyan, "\tKth smallest element = ", kthSmallestElement, colorReset)
	fmt.Println(colorGreen, "\tTIME COMPLEXITY: O(n)", colorReset)
	printDottedLine()
}

func selectionAlgorithmFundamentals() {
	maxMinUsingSorting()
	maxMinUsingLinearSearch()
	maxMinUsingTournamentMethod()
	maxMinByComparingPairs()
	kthSmallestUsingPartition()
}

func main() {
	selectionAlgorithmFundamentals()
}
