package sorting

import "fmt"

const (
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
)

// The PrintArray function prints each element of an integer array in a formatted manner.
func PrintArray(array []int) {
	fmt.Println(colorMagenta)
	for _, value := range array {
		fmt.Printf("\t%d", value)
	}
	fmt.Print("\n")
	fmt.Println(colorReset)
}

// The BubbleSort function sorts an array of integers using the bubble sort algorithm and prints the
// total number of passes.
func BubbleSort(array []int) {
	numberOfElements := len(array)
	totalPasses := 0
	for pass := 0; pass < numberOfElements-1; pass++ {
		for j := 0; j < numberOfElements-pass-1; j++ {
			if array[j] > array[j+1] {
				// Swap Elements
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
		totalPasses = pass
	}
	fmt.Println(colorMagenta, "\n\tTotal passes: ", totalPasses, colorReset)
	fmt.Print("\n")
}

// The function performs an optimized version of the bubble sort algorithm on an array of integers,
// keeping track of the total number of passes made.
func OptimizedBubbleSort(array []int) {
	numberOfElements := len(array)
	totalPasses := 0
	for pass := 0; pass < numberOfElements-1; pass++ {
		swapped := false
		for j := 0; j < numberOfElements-pass-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}
		// If no two elements are swapped then break out of the inner loop
		if !swapped {
			break
		}
		totalPasses = pass
	}
	fmt.Println(colorMagenta, "\n\tTotal passes: ", totalPasses, colorReset)
	fmt.Print("\n")
}

// The SelectionSort function sorts an array of integers in either ascending or descending order using
// the selection sort algorithm.
func SelectionSort(array []int, sortInAscending bool) {
	numberOfElements := len(array)
	for i := 0; i < numberOfElements-1; i++ {
		minOrMax := i
		for j := i + 1; j < numberOfElements; j++ {
			if sortInAscending {
				// Sort by minimum index
				if array[j] < array[minOrMax] {
					minOrMax = j
				}
			} else {
				// Sort by maximum index
				if array[j] > array[minOrMax] {
					minOrMax = j
				}
			}

		}
		if minOrMax != i {
			// Swap A[i] with A[minOrMax]
			array[i], array[minOrMax] = array[minOrMax], array[i]
		}
	}
}
