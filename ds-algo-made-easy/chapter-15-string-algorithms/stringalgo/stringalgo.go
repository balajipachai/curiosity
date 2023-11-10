package stringalgo

import (
	"fmt"
	"math"
	"time"
)

const (
	colorCyan    = "\033[36m"
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
	colorYellow  = "\033[33m"
	NO_OF_CHARS  = 256
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

/*
T - is the text containing the pattern P
P - is the pattern that we are looking for inside T
*/
// The function StringMatchBruteForce implements the brute force algorithm to find the first occurrence
// of a pattern P in a text T and returns the starting index of the match or -1 if no match is found.
func StringMatchBruteForce(T []int, P []int) int {
	start := time.Now()
	// Since the length of T is n, we have n – m + 1 possible choices for comparisons
	n := len(T)
	m := len(P)

	for i := 0; i <= n-m; i++ {
		j := 0
		for (j < m) && (P[j] == T[i+j]) {
			j += 1
			if j == m {
				return i
			}
		}
	}
	timeElapsed := time.Since(start)
	fmt.Println(colorCyan, "\t\tThe `StringMatchBruteForce` took ", timeElapsed, colorReset)
	return -1
}

/*
Algorithm:
-	Initially calculate the hash value of the pattern.
-	Start iterating from the starting of the string:
  - Calculate the hash value of the current substring having length m.
  - If the hash value of the current substring and the pattern are same check if the substring is same as the pattern.
  - If they are same, store the starting index as a valid answer. Otherwise, continue for the next substrings.

-	Return the starting indices as the required answer.
*/
// The Robin-Karp algorithm is used to find a pattern in a given text by calculating the hash values of
// the pattern and sliding window of the text.
func RobinKarp(T []int, P []int) int {
	start := time.Now()
	// Since the length of T is n, we have n – m + 1 possible choices for comparisons
	n := len(T)
	m := len(P)
	base := 256
	p := math.MaxInt

	patternHash := 0
	windowOfTextHash := 0

	// Calculate h = pow(base, m-1) % p
	h := int(math.Pow(float64(base), float64(m-1))) % p

	// Calcuate pattern hash and the first window of text
	for i := 0; i < m; i++ {
		patternHash = (base*patternHash + P[i]) % p
		windowOfTextHash = (base*windowOfTextHash + T[i]) % p
	}

	var j int = 0

	for i := 0; i <= n-m; i++ {
		if patternHash == windowOfTextHash {
			for j < m {
				if P[j] != T[i+j] {
					break
				}
				j += 1
			}
			if j == m {
				fmt.Println(colorCyan, "\tPattern found at index ", i, colorReset)
			}
		}

		// Calculate hash for the next window of text
		// Remove leading digit, add trailing digit
		if i < n-m {
			windowOfTextHash = (base*(windowOfTextHash-T[i]*h) + T[i+m]) % p
		}

		if windowOfTextHash < 0 {
			windowOfTextHash += p
		}
	}
	timeElapsed := time.Since(start)
	fmt.Println(colorCyan, "\t\tThe `RobinKarp` algorithm took ", timeElapsed, colorReset)
	return -1
}

// Global Variable, whihc is used as a prefix table for KMP algorithm
var F = make([]int, 0)

// The KMPPrefixTable function calculates the prefix table for a given pattern P.
func KMPPrefixTable(P []int) {
	m := len(P)
	F = make([]int, m)
	F[0] = 0
	var i, j int = 1, 0
	for i < m {
		if P[i] == P[j] {
			F[i] = j + 1
			i++
			j++
		} else if j > 0 {
			j = F[j-1]
		} else {
			F[i] = 0
			i++
		}
	}
}

// The KMP function implements the Knuth-Morris-Pratt algorithm to find the first occurrence of a
// pattern P in a text T and returns the starting index of the pattern in the text, or -1 if the
// pattern is not found.
func KMP(T []int, P []int) int {
	start := time.Now()
	i, j := 0, 0
	KMPPrefixTable(P)
	n := len(T)
	m := len(P)

	for i < n {
		if T[i] == P[j] {
			if j == m-1 {
				timeElapsed := time.Since(start)
				fmt.Println(colorCyan, "\t\tThe `KMP` algorithm took ", timeElapsed, colorReset)
				return i - j
			} else {
				i++
				j++
			}
		} else if j > 0 {
			j = F[j-1]
		} else {
			i++
		}
	}
	timeElapsed := time.Since(start)
	fmt.Println(colorCyan, "\t\tThe `KMP` algorithm took ", timeElapsed, colorReset)
	return -1
}

// The function badCharacterHeuristic assigns the index of the last occurrence of each character in the
// pattern P to the corresponding index in the badCharacter array.
func badCharacterHeuristic(P []int, badCharacter [NO_OF_CHARS]int) {
	m := len(P)

	for i := 0; i < NO_OF_CHARS; i++ {
		badCharacter[i] = -1
	}

	for i := 0; i < m; i++ {
		badCharacter[P[i]] = i
	}
}

// The Boyer-Moore algorithm is implemented in Go to search for a pattern in a given text.
func BoyerMoore(T []int, P []int) {
	start := time.Now()
	m := len(P)
	n := len(T)

	badCharacter := make([]int, NO_OF_CHARS)
	badCharacterHeuristic(P, [256]int(badCharacter))

	s := 0 // No. of shifts

	// Since there are n-m possible comparisons
	for s <= n-m {
		j := m - 1 // As comparison of the pattern begins from right to left
		for j >= 0 && (P[j] == T[s+j]) {
			j -= 1
		}

		if j < 0 {
			fmt.Println(colorCyan, "\t\tPattern occurs at shift = ", s, colorReset)
			if (s + m) < n {
				s += m - badCharacter[T[s+m]]
			} else {
				s += 1
			}
		} else {
			k := j - badCharacter[T[s+j]]
			if k > 0 {
				s += k
			} else {
				s += 1
			}
		}
	}
	timeElapsed := time.Since(start)
	fmt.Println(colorCyan, "\t\tThe `Boyer-Moore` algorithm took ", timeElapsed, colorReset)
}

// TRIE DATASTRUCTURE
type TrieNode struct {
	isEndOfWord bool
	childrens   [26]*TrieNode
}

// The function `GetTrieNode` returns a new instance of a TrieNode with its properties initialized.
func GetTrieNode() *TrieNode {
	newNode := &TrieNode{}
	newNode.isEndOfWord = false
	for i := 0; i < 26; i++ {
		newNode.childrens[i] = nil
	}
	return newNode
}

// The `Insert` function is used to insert a word into a Trie data structure. It takes a slice of
// characters as input and iterates through each character. For each character, it calculates the index
// by subtracting the ASCII value of 'a' from the character. It then checks if the child node at that
// index is nil. If it is nil, it creates a new TrieNode and assigns it to the child node at that
// index. It then moves the iterator to the child node at that index. Finally, it sets the
// `isEndOfWord` property of the last node to true, indicating that a word ends at that node.
func (root *TrieNode) Insert(characters []int) {
	n := len(characters)
	iterator := root
	for i := 0; i < n; i++ {
		index := characters[i] - 'a'
		if iterator.childrens[index] == nil {
			iterator.childrens[index] = GetTrieNode()
		}
		iterator = iterator.childrens[index]
	}
	iterator.isEndOfWord = true
}

// The `Search` function is used to search for a word in a Trie data structure. It takes a slice of
// characters as input and iterates through each character. For each character, it calculates the index
// by subtracting the ASCII value of 'a' from the character. It then checks if the child node at that
// index is nil. If it is nil, it means that the word does not exist in the Trie and the function
// returns false. If the child node is not nil, it moves the iterator to the child node at that index.
// Finally, it checks if the `isEndOfWord` property of the last node is true, indicating that a word
// ends at that node. If it is true, it means that the word exists in the Trie and the function returns
// true.
func (root *TrieNode) Search(characters []int) bool {
	n := len(characters)
	iterator := root
	for i := 0; i < n; i++ {
		index := characters[i] - 'a'
		if iterator.childrens[index] == nil {
			return false
		}
		iterator = iterator.childrens[index]
	}
	return iterator.isEndOfWord
}

// The function recursively prints all the words stored in a Trie data structure.
func Print(root *TrieNode, characters [26]int, level int) {
	iterator := root
	if iterator.isEndOfWord {
		PrintArray(characters[:], true)
	}

	for i := 0; i < 26; i++ {
		if iterator.childrens[i] != nil {
			characters[level] = i + 'a'
			Print(root.childrens[i], characters, level+1)
		}
	}
}
