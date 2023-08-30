package main

import (
	"fmt"

	"example.com/heaps"
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

// The function creates a new heap, inserts the numbers 1 to 7 into it, and returns the heap.
func getHeap() *heaps.Heap {
	heap := &heaps.Heap{}
	emptySlice := []int{}
	heap.Create(emptySlice, 7, 0)
	fmt.Println("\t"+colorCyan+"Insert 1, 2, 3, 4, 5, 6, & 7 in the heap", colorReset)
	heaps.Insert(heap, 1)
	heaps.Insert(heap, 2)
	heaps.Insert(heap, 5)
	heaps.Insert(heap, 3)
	heaps.Insert(heap, 4)
	heaps.Insert(heap, 6)
	heaps.Insert(heap, 7)
	return heap
}

func seven() {
	heap := getHeap()
	fmt.Println("\t"+colorCyan+"Printing the heap elements", colorReset)
	heap.Print()
	fmt.Println("\t"+colorCyan+"Maximum element in the given heap is: ", heaps.MaximumInMinHeap(heap), colorReset)
	printDottedLine()
}

func eight() {
	heap := getHeap()
	fmt.Println("\t"+colorYellow+"Printing the heap elements", colorReset)
	heap.Print()
	fmt.Println("\t"+colorYellow+"Deleting 3 from the heap", colorReset)
	fmt.Println("\t", colorCyan, heap.DeleteArbitrary(3), colorReset)
	heap.Print()
	printDottedLine()
}

func eleven() {
	heap := getHeap()
	fmt.Println("\t"+colorYellow+"Printing the heap elements", colorReset)
	heap.Print()
	fmt.Println("\t"+colorYellow+"All elements < 5", colorReset)
	fmt.Println("\t"+colorCyan, heaps.AllElementsLessThanK(heap, 5), colorReset)
	fmt.Println("\n\t"+colorYellow+"All elements < 1", colorReset)
	fmt.Println("\n\t"+colorCyan, heaps.AllElementsLessThanK(heap, 1), colorReset)
	fmt.Println("\n\t"+colorYellow+"All elements < 7", colorReset)
	fmt.Println("\n\t"+colorCyan, heaps.AllElementsLessThanK(heap, 7), colorReset)
	fmt.Println(colorYellow + "\tRECURSIVE: Elements < 7" + colorReset)
	heaps.RecursiveAllElementsLessThanK(heap, 7, 0)

	fmt.Println("\n\t"+colorYellow+"All elements > 7", colorReset)
	fmt.Println("\n\t"+colorCyan, heaps.AllElementsGreaterThanK(heap, 7), colorReset)
	fmt.Println("\n\t"+colorYellow+"All elements > 1", colorReset)
	fmt.Println("\n\t"+colorCyan, heaps.AllElementsGreaterThanK(heap, 1), colorReset)
	fmt.Println(colorYellow + "\tRECURSIVE: Elements > 1" + colorReset)
	heaps.RecursiveAllElementsGreaterThanK(heap, 1, 0)
	fmt.Println()
	printDottedLine()
}

// TODO WIP
func twelve() {

	emptySlice := []int{}

	h1 := &heaps.Heap{}
	h1.Create(emptySlice, 7, 1)

	heaps.Insert(h1, 1)
	heaps.Insert(h1, 2)
	heaps.Insert(h1, 5)
	heaps.Insert(h1, 3)
	heaps.Insert(h1, 4)
	heaps.Insert(h1, 6)
	heaps.Insert(h1, 7)

	h1.Print()

	h2 := &heaps.Heap{}
	h2.Create(emptySlice, 7, 1)

	heaps.Insert(h2, 15)
	heaps.Insert(h2, 13)
	heaps.Insert(h2, 14)
	heaps.Insert(h2, 9)
	heaps.Insert(h2, 10)
	heaps.Insert(h2, 11)
	heaps.Insert(h2, 12)

	h2.Print()

	fmt.Println(colorYellow + "\tMerge h1 with h2 and then heapify the array" + colorReset)
	heaps.MergeMaxHeaps(h1, h2)
	h1.Print()

	printDottedLine()

}

func executeExercises() {
	fmt.Println("PRIORITY QUEUE & HEAPS EXERCISES")
	fmt.Println("Problem 7: Given a min-heap, give an algorithm for finding the maximum element.")
	seven()
	fmt.Println("Problem 8: Give an algorithm for deleting an arbitrary element from min heap.")
	eight()
	fmt.Println("Problem 9: Give an algorithm for deleting the ith indexed element in a given min-heap.")
	fmt.Println("\t" + colorCyan + "SAME AS SOLUTION OF PROBLEM 8" + colorReset)
	fmt.Println("Problem 11: Give an algorithm to find all elements less than some value of k in a binary heap.")
	eleven()
	fmt.Println("Problem 12: Give an algorithm for merging two binary max-heaps. Let us assume that the size of the first heap is m + n and the size of the second heap is n.")
	twelve()
}

func main() {
	executeExercises()
}
