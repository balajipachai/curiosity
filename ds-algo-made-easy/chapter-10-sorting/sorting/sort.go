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

// The InsertionSort function sorts an array of integers using the insertion sort algorithm and prints
// the number of shifts made during the sorting process.
func InsertionSort(array []int) {
	numberOfElements := len(array)
	shifts := 0
	fmt.Println(colorMagenta, "\tPRINTING SHIFTS", colorReset)
	for i := 1; i <= numberOfElements-1; i++ {
		v := array[i]
		j := i
		// Shifting elements to the right side of the smallest element
		for j >= 1 && (array[j-1] > v) {
			array[j] = array[j-1]
			j--
			shifts++
		}
		array[j] = v
		PrintArray(array)
	}
	fmt.Println(colorMagenta+"\tTotal Shifts: ", shifts, colorReset)
}

func ShellSort(array []int) {
	numberOfElements := len(array)
	fmt.Println(colorMagenta, "\tPRINTING SHIFTS", colorReset)
	for gap := numberOfElements / 2; gap > 0; gap /= 2 {
		for i := gap; i < numberOfElements; i++ {
			temp := array[i]
			var j int
			for j = i; j >= gap && array[j-gap] > temp; j -= gap {
				array[j] = array[j-gap]
			}
			array[j] = temp
			PrintArray(array)
		}
	}
}

func merge(array []int, temp []int, left, mid, right int) {
	leftEnd := mid - 1
	tempPos := left
	size := right - left + 1

	// iske andar merge k 3 conditions hai
	// pehla, jaha pe left, leftEnd se chhota hai & mid, right se chhota hai
	for (left <= leftEnd) && (mid <= right) {
		// iske andar do conditions hai
		// pehla, array[left] <= array[mid]
		// dusra, array[mid] <= array[left]
		if array[left] <= array[mid] {
			temp[tempPos] = array[left]
			tempPos += 1
			left += 1
		} else {
			temp[tempPos] = array[mid]
			tempPos += 1
			mid += 1
		}
	}
	// dusra condition jaha pe sirf left jo hai wo leftEnd se chhota hai
	for left <= leftEnd {
		temp[tempPos] = array[left]
		tempPos += 1
		left += 1
	}

	// teesra condition jaha pe sirft mid jo hai wo right se chhota hai
	for mid <= right {
		temp[tempPos] = array[mid]
		tempPos += 1
		mid += 1
	}

	// elements ko apne sahi jagah pe daalo resultant array mai
	for i := 0; i < size; i++ {
		array[right] = temp[right]
		right -= 1
	}
}

func MergeSort(array []int, temp []int, left, right int) {
	var mid int

	if right > left {
		mid = (right + left) / 2
		MergeSort(array, temp, left, mid)
		MergeSort(array, temp, mid+1, right)
		merge(array, temp, left, mid+1, right)
	}
}
