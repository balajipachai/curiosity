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
	fmt.Println("\t"+colorYellow+"All elements less than 5", colorReset)
	heaps.AllElementsLessThanK(heap, 5, 0)
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
}

func main() {
	executeExercises()
}
