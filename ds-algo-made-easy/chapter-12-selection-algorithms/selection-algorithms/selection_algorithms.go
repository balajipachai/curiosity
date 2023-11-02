/*
	Selection algorithm is a set of algorithms for finding the kth smallest/largest element in a list.
	It is also called kth order statistic.
		1. Partition based Selection Algorithm
		2. Linear Selection Algorithm - Median of Medians Algorithm
		3. Finding the Kth smallest element in sorted order
*/

package selectionalgorithms

import (
	"fmt"

	"example.com/sorting"
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

func whyWasteResources(array []int) (max, min int) {
	if len(array) == 0 {
		return 0, 0
	}

	if len(array) == 1 {
		return array[0], array[0]
	}

	if len(array) == 2 {
		if array[0] > array[1] {
			return array[0], array[1]
		} else {
			return array[1], array[0]
		}
	}
	return max, min
}

// Finding Maximum & Minimum of an array using the below different methods
/*
	1. Maximum & Minimum of an array using Sorting
*/
// The MaxMinUsingSorting function takes an array of integers, sorts it in ascending order, and returns
// the minimum and maximum values in the array.
func MaxMinUsingSorting(array []int) (max int, min int) {
	if len(array) > 2 {
		sortedArray := sorting.CountingSort(array, false)
		fmt.Println(colorMagenta, "Printing the array after sorting:", colorReset)
		sorting.PrintArray(sortedArray)
		return sortedArray[len(sortedArray)-1], sortedArray[0]
	}
	return whyWasteResources(array)
}

/*
2. Maximum & Minimum of an array using Linear Search

	Initialize values of min and max as minimum and maximum of the first two elements respectively. Starting from 3rd, compare each element with max and min, and change max and min accordingly (i.e., if the element is smaller than min then change min, else if the element is greater than max then change max, else ignore the element)
*/
// The function MaxMinUsingLinearSearch finds the maximum and minimum values in an array using linear
// search.
func MaxMinUsingLinearSearch(array []int) (max int, min int) {
	if len(array) > 2 {
		min = array[0]
		max = array[1]
		for i := 2; i < len(array); i++ {
			if array[i] > max {
				max = array[i]
			} else if array[i] < min {
				min = array[i]
			}
		}
		return
	}
	return whyWasteResources(array)
}

/*
3. Maximum & Minimum of an array using Tournament Method

	Divide the array into two parts and compare the maximums and minimums of the two parts to get the maximum and the minimum of the whole array.
*/
// The function uses the tournament method to find the maximum and minimum values in an array.
func MaxMinUsingTournamentMethod(array []int, low, high int) (max, min int) {

	if low == high {
		max = array[low]
		min = array[low]
		return
	}

	if high == low+1 {
		if array[low] > array[high] {
			max = array[low]
			min = array[high]
		} else {
			max = array[high]
			min = array[low]
		}
		return
	}

	mid := (low + high) / 2
	max1, min1 := MaxMinUsingTournamentMethod(array, low, mid)
	max2, min2 := MaxMinUsingTournamentMethod(array, mid+1, high)
	if max1 > max2 {
		max = max1
	} else {
		max = max2
	}

	if min1 < min2 {
		min = min1
	} else {
		min = min2
	}
	return
}

/*
4. Maximum & Minimum of an array by comparing in pairs

	If n is odd then initialize min and max as the first element.
	If n is even then initialize min and max as minimum and maximum of the first two elements respectively.
	For the rest of the elements, pick them in pairs and compare their
	maximum and minimum with max and min respectively.
*/
// The function MaxMinByComparingPairs takes an array of integers and returns the maximum and minimum
// values in the array by comparing pairs of elements.
func MaxMinByComparingPairs(array []int) (max, min int) {
	if len(array) > 2 {
		var i int
		// n is even
		if len(array)%2 == 0 {
			if array[0] > array[1] {
				max = array[0]
				min = array[1]
			} else {
				max = array[1]
				min = array[0]
			}
			// Set loop starting index
			i = 2
		} else {
			// n is odd
			min = array[0]
			max = array[0]
			// Set loop starting index
			i = 1
		}
		// Proceed to the next steps as min and max have been initialized
		for i < len(array)-1 {
			if array[i] > array[i+1] {
				if array[i] > max {
					max = array[i]
				}

				if array[i+1] < min {
					min = array[i+1]
				}
			} else {
				if array[i+1] > max {
					max = array[i+1]
				}

				if array[i] < min {
					min = array[i]
				}
			}
			// Increment the loop index by 2
			i += 2
		}
		return
	}
	return whyWasteResources(array)
}
