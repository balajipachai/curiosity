package hashing

import (
	"fmt"
	"sort"
	"time"
)

const (
	colorCyan    = "\033[36m"
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
	colorYellow  = "\033[33m"
)

// The PrintArray function prints each element of an integer array in a formatted manner.
func PrintArray(char []int, printAsChar bool) {
	fmt.Println(colorMagenta)
	for _, value := range char {
		if printAsChar {
			fmt.Printf("\t%c", value)
		} else {
			fmt.Printf("\t%d", value)
		}

	}
	fmt.Print("\n")
	fmt.Println(colorReset)
}

// Problem 2 Given an array of characters, give an algorithm for removing the duplicates.
/*
Check if a duplicate character exists, if yes, replace that character with the last character
*/
// The `RemoveDuplicates` function takes an array of integers and returns a new array with all
// duplicate elements removed.
func RemoveDuplicates(char []int) []int {
	start := time.Now()
	n := len(char)
	result := []int{}

	for i := 0; i < n; i++ {
		isDuplicate := false

		for j := 0; j < i; j++ {
			if char[i] == char[j] {
				isDuplicate = true
				break
			}
		}

		if !isDuplicate {
			result = append(result, char[i])
		}
	}
	timeElapsed := time.Since(start)
	fmt.Println(colorCyan, "\t\tThe `RemoveDuplicates` took ", timeElapsed, colorReset)
	return result
}

// The `RemoveDuplicatesWithSorting` function removes duplicate elements from a slice of integers by
// sorting the slice and then comparing adjacent elements.
func RemoveDuplicatesWithSorting(chars []int) []int {
	start := time.Now()
	sort.Ints(chars)
	result := []int{}
	for i := 0; i < len(chars)-1; i++ {
		if chars[i] != chars[i+1] {
			result = append(result, chars[i])
		}
	}
	result = append(result, chars[len(chars)-1])
	timeElapsed := time.Since(start)
	fmt.Println(colorCyan, "\t\tThe `RemoveDuplicatesWithSorting` took ", timeElapsed, colorReset)
	return result
}

// The `RemoveDuplicatesWithHashing` function removes duplicate elements from an array using a hash
// map.
func RemoveDuplicatesWithHashing(chars []int) []int {
	start := time.Now()
	existsInHashMap := make(map[int]bool)
	result := []int{}
	for _, v := range chars {
		if !existsInHashMap[v] {
			result = append(result, v)
		}
	}
	timeElapsed := time.Since(start)
	fmt.Println(colorCyan, "\t\tThe `RemoveDuplicatesWithHashing` took ", timeElapsed, colorReset)
	return result
}

// The `AreArraysSame` function checks if two arrays have the same elements, regardless of their order,
// and returns true if they do.
func AreArraysSame(a []int, b []int) bool {
	start := time.Now()
	valueMap := make(map[int]bool)
	indexMap := make(map[int]bool)
	sameCount := 0
	for _, av := range a {
		for bi, bv := range b {
			if av == bv {
				if !valueMap[bv] && !indexMap[bi] {
					valueMap[bv] = true
					indexMap[bi] = true
					sameCount++
				}
			}
		}
	}
	timeElapsed := time.Since(start)
	fmt.Println(colorCyan, "\t\tThe `AreArraysSame` took ", timeElapsed, colorReset)
	return sameCount == len(a)
}

// The function `AreArraysSameUsingSorting` checks if two arrays are the same by sorting them and
// comparing each element.
func AreArraysSameUsingSorting(a []int, b []int) bool {
	start := time.Now()
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	// Sort a & b
	sort.Ints(a)
	sort.Ints(b)
	// Check if a[i] == b[i]
	for ai := range a {
		if a[ai] != b[ai] {
			return false
		}

	}
	timeElapsed := time.Since(start)
	fmt.Println(colorCyan, "\t\tThe `AreArraysSameUsingSorting` took ", timeElapsed, colorReset)
	return true
}

/*
Algorithm
Construct the hash table with array A elements as keys.
• While inserting the elements, keep track of the number frequency for each number.
That means, if there are duplicates, then increment the counter of that corresponding
key.
• After constructing the hash table for A’s elements, now scan the array B.
• For each occurrence of B’s elements reduce the corresponding counter values.
• At the end, check whether all counters are zero or not.
• If all counters are zero, then both arrays are the same otherwise the arrays are
different.
*/
func AreArraysSameUsingHashTable(a []int, b []int) bool {
	start := time.Now()
	counterHash := make(map[int]int)
	for _, av := range a {
		counterHash[av]++
	}

	for _, bv := range b {
		counterHash[bv]--
	}

	for _, value := range counterHash {
		if value != 0 {
			return false
		}
	}
	timeElapsed := time.Since(start)
	fmt.Println(colorCyan, "\t\tThe `AreArraysSameUsingHashTable` took ", timeElapsed, colorReset)
	return true
}
