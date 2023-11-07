package sorting

import (
	"fmt"
	"math"

	"example.com/trees"
)

const (
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
	colorYellow  = "\033[33m"
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

// The merge function takes in an array, a temporary array, and three indices (left, mid, right) and
// merges the two subarrays within the given indices in a sorted manner.
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

// The MergeSort function recursively divides an array into smaller subarrays, sorts them, and then
// merges them back together.
func MergeSort(array []int, temp []int, left, right int) {
	var mid int

	if right > left {
		mid = (right + left) / 2
		MergeSort(array, temp, left, mid)
		MergeSort(array, temp, mid+1, right)
		merge(array, temp, left, mid+1, right)
	}
}

// The partition function takes an array, selects a pivot element, and rearranges the elements such
// that all elements smaller than the pivot are placed before it, and all elements greater than the
// pivot are placed after it.
func Partition(array []int, low, high int) int {
	pivot := array[high]

	i := low - 1

	for j := low; j <= high-1; j++ {
		if array[j] < pivot {
			i++
			// swap array[i] with array[j]
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[high] = array[high], array[i+1]
	fmt.Println(colorMagenta, "\tPivot:\t", pivot, colorReset)
	return i + 1
}

// The QuickSort function recursively sorts an array of integers using the QuickSort algorithm.
func QuickSort(array []int, low, high int) {
	var pivot int

	if high > low {
		pivot = Partition(array, low, high)
		QuickSort(array, low, pivot-1)
		QuickSort(array, pivot+1, high)
	}
}

// The TreeSort function uses a Binary Search Tree to sort an array of integers.
func TreeSort(array []int) {
	// Create a Binary Search Tree (BST) using the array elements
	bst := &trees.BinaryTree{}
	for _, v := range array {
		bst.InsertBST(v)
	}
	// Traverse the BST using inorder, thus resulting in a sorted array
	bst.InOrder()
}

// The CountingSort function sorts an array of integers in ascending order using the counting sort
// algorithm.
func CountingSort(inputArray []int, shouldPrint bool) []int {
	outputArray := make([]int, len(inputArray))
	max := inputArray[0]
	// Find Max element in inputArray
	for _, v := range inputArray {
		if v > max {
			max = v
		}
	}

	// Create countArray with size max+1
	countArray := make([]int, max+1)

	// Count all the distinct elements from input inputArray and store them in countArray
	// For example, let us assume the inputArray = inputArray := []int{2, 5, 3, 0, 2, 3, 0, 3}
	// Thus, count of 5 = 1, hence countArray[5] = 1, similary countArray[2] = 2
	for i := 0; i < len(inputArray); i++ {
		countArray[inputArray[i]]++
	}

	// Calculate Prefix Sum of countArray
	for i := 1; i <= max; i++ {
		countArray[i] += countArray[i-1]
	}

	// Creating outputArray from countArray
	for i := len(inputArray) - 1; i >= 0; i-- {
		outputArray[countArray[inputArray[i]]-1] = inputArray[i]
		countArray[inputArray[i]]--
	}

	if shouldPrint {
		fmt.Println(colorYellow + "\tPrinting the sorted array:" + colorReset)
		PrintArray(outputArray)
	}

	return outputArray
}

// The BucketSort function sorts an array of integers using the bucket sort algorithm.
func BucketSort(array []int) {
	n := len(array)
	// Create buckets of size len(array)
	buckets := make([][]int, n)
	// Initialize buckets
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}

	/*
	 Put array elements in different buckets
	 Here, we divide the array element by n which gives us the bucketIndex
	 For example:
	 If n = 10, and array has: 64, 61, 63
	 For the above element, the index will always be 6 and
	 buckets[6] is an array i.e. buckets[6] = [] which will then contain 64, 61, 63
	 Then bucket[6] is sorted using SelectionSort
	 This process is repeated for all other buckets.
	 At the end all the buckets are merged, thus, resulting in a sorted array.
	*/
	for i := 0; i < n; i++ {
		bucketIndex := array[i] / n
		// Push element into the bucket
		buckets[bucketIndex] = append(buckets[bucketIndex], array[i])
	}

	fmt.Println(colorMagenta, "\tCreated buckets are:")
	for i := 0; i < n && len(buckets[i]) > 0; i++ {
		fmt.Println(colorYellow, "\t\t", buckets[i])
	}

	fmt.Println(colorMagenta, "\tSorted Individual buckets using SelectionSort")
	// Sort individual buckets
	for i := 0; i < n && len(buckets[i]) > 0; i++ {
		SelectionSort(buckets[i], true)
		fmt.Println(colorYellow, "\t\t", buckets[i])
	}
	fmt.Println(colorReset)

	// Concatenate all the buckets
	index := 0
	for i := 0; i < n; i++ {
		for j := 0; j < len(buckets[i]) && index < n; j++ {
			array[index] = buckets[i][j]
			index++
		}
	}
}

func RadixSort(array []int) {
	n := len(array)
	max := 0

	outputArray := make([]int, n)

	// Find the maximum number in the array
	for i := 0; i < n; i++ {
		if array[i] > max {
			max = array[i]
		}
	}

	numberOfDigits := 0
	// Find the number of digits in the maximum number
	for max > 0 {
		max /= 10
		numberOfDigits++
	}

	for i := 0; i < numberOfDigits; i++ {
		lastDigit := make([]int, n)
		for j := 0; j < n; j++ {
			lastDigit[j] = int(float64(array[j])/math.Pow10(i)) % 10
		}
		sortedByDigits := CountingSort(lastDigit, false)
		// Update array as per the unit digits sorted
		for k1, v1 := range lastDigit {
			for k2, v2 := range sortedByDigits {
				if v1 == v2 {
					outputArray[k2] = array[k1]
					sortedByDigits[k2] = -1
					break
				}
			}
		}
		// Assign outputArray to original array
		copy(array, outputArray)
	}
}
