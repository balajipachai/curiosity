package searching

import "fmt"

const (
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
	colorYellow  = "\033[33m"
)

/*
What is Searching?
	- Searching is the process of finding an item with specified properties from a collection of items.

Why do we need searching?
	- To retrieve the stored information proficiently we need very efficient searching algorithms.
*/

// The PrintArray function prints each element of an integer array in a formatted manner.
func PrintArray(array []int) {
	fmt.Println(colorMagenta)
	for _, value := range array {
		fmt.Printf("\t%d", value)
	}
	fmt.Print("\n")
	fmt.Println(colorReset)
}

func UnorderedLinearSearch(array []int, data int) int {
	for key, value := range array {
		if value == data {
			return key
		}
	}
	return -1
}

func SortedOrOrderedLinearSearch(array []int, data int) int {
	for key, value := range array {
		if value == data {
			return key
		} else if value > data {
			return -1
		}
	}
	return -1
}

func BinarySearchIterative(array []int, data int) int {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := low + ((high - low) / 2) // To avoid overflow
		if array[mid] == data {
			return mid
		} else if array[mid] < data {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func binarySearchRecursiveHelper(array []int, low, high, data int) int {
	mid := low + ((high - low) / 2)
	if low > high {
		return -1
	}
	if array[mid] == data {
		return mid
	} else if array[mid] < data {
		return binarySearchRecursiveHelper(array, mid+1, high, data)
	} else {
		return binarySearchRecursiveHelper(array, low, mid-1, data)
	}
}

func BinarySearchRecursive(array []int, data int) int {
	return binarySearchRecursiveHelper(array, 0, (len(array) - 1), data)
}

func InterpolationSearch(array []int, data int) int {
	low := 0
	high := len(array) - 1
	for low <= high && data >= array[low] && data <= array[high] {
		mid := low + (((data - array[low]) * (high - low)) / (array[high] - array[low]))
		if array[mid] == data {
			return mid + 1
		}
		if data < array[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
