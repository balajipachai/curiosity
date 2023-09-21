package main

import (
	"fmt"

	"example.com/sorting"
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

// The function bubbleSort() sorts an array using the bubble sort algorithm and prints the unsorted and
// sorted arrays.
func bubbleSort() {
	fmt.Println(colorCyan + "\t************BUBBLE SORT************" + colorReset)
	array := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println(colorYellow + "\tPrinting the unsorted array:" + colorReset)
	sorting.PrintArray(array)
	sorting.BubbleSort(array)
	fmt.Println(colorYellow + "\tPrinting the sorted array:" + colorReset)
	sorting.PrintArray(array)
	fmt.Println(colorYellow + "\tInvoking bubble sort (UnOptimized) for alreay sorted array:" + colorReset)
	sorting.BubbleSort([]int{11, 12, 22, 25, 34, 64, 90})
	fmt.Println(colorYellow + "\tInvoking bubble sort (Optimized) for alreay sorted array:" + colorReset)
	sorting.OptimizedBubbleSort([]int{11, 12, 22, 25, 34, 64, 90})
	printDottedLine()
}

/*
Notes:
- Selection sort works by finding the minimum element and then inserting it in its correct position by swapping with the element which is in the position of this minimum element. This is what makes it unstable.
- This algorithm is called selection sort since it repeatedly selects the smallest/largest element.
*/
// The selectionSort function sorts an array of integers using the selection sort algorithm and prints
// the sorted array in ascending and descending order.
func selectionSort() {
	fmt.Println(colorCyan + "\t************SELECTION SORT************" + colorReset)
	array := []int{64, 25, 12, 22, 11}
	fmt.Println(colorYellow + "\tPrinting the unsorted array:" + colorReset)
	sorting.PrintArray(array)
	sorting.SelectionSort(array, true)
	fmt.Println(colorYellow + "\tPrinting the sorted array (SORT BY ASCENDING ORDER)" + colorReset)
	sorting.PrintArray(array)
	fmt.Println(colorYellow + "\tPrinting the sorted array (SORT BY DESCENDING ORDER)" + colorReset)
	array = []int{64, 25, 12, 22, 11}
	sorting.SelectionSort(array, false)
	sorting.PrintArray(array)
	printDottedLine()
}

func sortingBasics() {
	bubbleSort()
	selectionSort()
}

func main() {
	sortingBasics()
}
