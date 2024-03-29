package heaps

import "fmt"

const (
	QUEUE_SIZE   = 10
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
)

// The `Heap` type represents a heap data structure with a fixed capacity and the ability to store
// integers in either a min-heap or max-heap configuration.
// @property {[]int} elements - The "elements" property is a slice of integers that stores the actual elements of the heap.
// @property {int} count - The count property represents the number of elements currently present in the heap.
// @property {int} capacity - The capacity property represents the maximum number of elements that the heap can hold.
// @property {int} heapType - The heapType property is used to determine whether the heap is a MinHeap or a MaxHeap.
type Heap struct {
	elements []int
	count    int // Number of elements in the heap
	capacity int // Capacity of heap, size of heap
	heapType int // MinHeap = 0, MaxHeap = 1
}

// The `Create` function is used to initialize a new instance of the `Heap` struct with the given parameters.
func (h *Heap) Create(elements []int, capacity, heapType int) {
	h.elements = make([]int, capacity)
	h.capacity = capacity
	h.count = 0
	h.heapType = heapType
}

// The `ParentOfNode` function is that calculates the index of the parent node for a given node in the heap.
func ParentOfNode(h *Heap, location int) int {
	if location <= 0 || location >= h.count {
		return -1
	}
	return (location - 1) / 2
}

// The function returns the index of the left child of a node in a heap.
func LeftChildOfNode(h *Heap, location int) int {
	left := (2 * location) + 1
	if left >= h.count {
		return -1
	}
	return left
}

// The function returns the index of the right child of a node in a heap.
func RightChildOfNode(h *Heap, location int) int {
	right := (2 * location) + 2
	if right >= h.count {
		return -1
	}
	return right
}

// The MaximumElement function returns the maximum element in a heap.
func MaximumElement(h *Heap) int {
	if h.count == 0 {
		return -1
	}
	return h.elements[0]
}

// The PercolateDown function is used to maintain the heap property by recursively swapping elements to
// ensure that the maximum element is at the root.
func PercolateDown(h *Heap, location int) {
	// Max-Heap
	if h.heapType == 1 {
		max := 0
		leftChild := LeftChildOfNode(h, location)
		rightChild := RightChildOfNode(h, location)

		if leftChild == -1 && rightChild == -1 {
			return
		}

		if leftChild != -1 && h.elements[leftChild] > h.elements[location] {
			max = leftChild
		} else {
			max = location
		}

		if rightChild != -1 && h.elements[rightChild] > h.elements[max] {
			max = rightChild
		}

		if max != location {
			// Swap the elements
			h.elements[location], h.elements[max] = h.elements[max], h.elements[location]
		}
		PercolateDown(h, max)
	} else {
		// Min-Heap
		min := -1
		leftChild := LeftChildOfNode(h, location)
		rightChild := RightChildOfNode(h, location)

		if leftChild == -1 && rightChild == -1 {
			return
		}

		if leftChild != -1 && h.elements[leftChild] < h.elements[location] {
			min = leftChild
		} else {
			min = location
		}

		if rightChild != -1 && h.elements[rightChild] < h.elements[min] {
			min = rightChild
		}

		if min != location {
			h.elements[min], h.elements[location] = h.elements[location], h.elements[min]
		}
		PercolateDown(h, min)
	}
}

// The IsEmpty function checks if a heap is empty.
func IsEmpty(h *Heap) bool {
	return h.count == 0
}

// The Delete function removes the maximum element from a heap and returns it.
func Delete(h *Heap) int {
	if h.count == 0 {
		return -1
	}

	// Max-Heap
	if h.heapType == 1 {
		max := h.elements[0]
		h.elements[0] = h.elements[h.count-1]
		h.elements = h.elements[:h.count-1]
		h.count -= 1
		PercolateDown(h, 0)
		return max
	} else {
		// Min-Heap
		min := h.elements[0]
		h.elements[0] = h.elements[h.count-1]
		h.elements = h.elements[:h.count-1]
		h.count -= 1
		PercolateDown(h, 0)
		return min
	}
}

// The Resize function doubles the capacity of a heap by creating a new array and copying the elements
// from the old array.
func Resize(h *Heap) {
	oldArray := h.elements
	h.elements = make([]int, h.capacity*2)
	copy(h.elements, oldArray)
}

// The Insert function inserts a new element into a heap and ensures that the heap property is maintained.
func Insert(h *Heap, data int) {
	if h.count == h.capacity {
		Resize(h)
	}

	h.count += 1
	i := h.count - 1

	// Max-Heap
	if h.heapType == 1 {
		// Percolate up
		for i > 0 && data > h.elements[(i-1)/2] {
			h.elements[i] = h.elements[(i-1)/2]
			i = (i - 1) / 2
		}
		h.elements[i] = data

	} else {
		// Min-Heap
		// Percolate up
		for i >= 0 && data < h.elements[(i-1)/2] {
			h.elements[i] = h.elements[(i-1)/2]
			i = (i - 1) / 2
		}
		h.elements[i] = data
	}
}

// The BuildHeap function builds a heap by assigning elements from an array to the heap and then
// percolating down from the first leaf node.
func BuildHeap(heap *Heap, array []int, n int) {
	/*
		Algorithm:
		1. When n > heap capacity, resize the heap
		2. Assign heap[i] = array[i]
		3. Start percolating down from the first leaf node
	*/
	if n > heap.capacity {
		Resize(heap)
	}

	fmt.Println("len of heap.elements: ", len(heap.elements))
	heap.elements = append(heap.elements, array...)
	fmt.Println("len of heap.elements: ", len(heap.elements))
	heap.count = n
	for j := (n - 1) / 2; j >= 0; j-- {
		PercolateDown(heap, j)
	}
}

// The `Print` function is a method of the `Heap` struct. It is used to print the elements of the heap
// in a formatted way. It iterates over the `elements` slice of the heap and prints each element
// followed by a tab character. After printing all the elements, it prints a new line character to
// separate the output.
func (h *Heap) Print() {
	fmt.Println(colorMagenta)
	fmt.Printf("\t")
	for _, v := range h.elements {
		fmt.Printf("%d\t", v)
	}
	fmt.Println(colorReset)
}

// The function MaximumInMinHeap finds the maximum element in a min heap.
func MaximumInMinHeap(h *Heap) int {
	max := -1
	for i := (h.count + 1) / 2; i < h.count; i++ {
		if h.elements[i] > max {
			max = h.elements[i]
		}
	}
	return max
}

// The `DeleteArbitrary` function is a method of the `Heap` struct. It is used to delete an arbitrary
// element from the heap.
func (heap *Heap) DeleteArbitrary(data int) int {
	element := -1
	for i, v := range heap.elements {
		if v == data {
			element = v
			heap.elements[i] = heap.elements[heap.count-1]
			heap.elements = heap.elements[:heap.count-1]
			heap.count -= 1
			PercolateDown(heap, i)
		}
	}
	return element
}

// The function `AllElementsLessThanK` takes a heap and a value `k` as input, and returns a list of all
// elements in the heap that are less than `k`.
func AllElementsLessThanK(heap *Heap, k int) []int {
	result := []int{}
	var traverse func(index int)

	traverse = func(index int) {
		if index >= len(heap.elements) || index < 0 {
			return
		}

		if index >= 0 && heap.elements[index] < k {
			result = append(result, heap.elements[index])
		}

		traverse(LeftChildOfNode(heap, index))
		traverse(RightChildOfNode(heap, index))
	}
	traverse(0) // Start traversal from root
	return result
}

// The function `AllElementsGreaterThanK` takes a heap and an integer `k` as input and returns a slice
// containing all elements in the heap that are greater than `k`.
func AllElementsGreaterThanK(heap *Heap, k int) []int {
	result := []int{}
	var traverse func(index int)

	traverse = func(index int) {
		if index >= len(heap.elements) || index < 0 {
			return
		}

		if index >= 0 && heap.elements[index] > k {
			result = append(result, heap.elements[index])
		}

		traverse(LeftChildOfNode(heap, index))
		traverse(RightChildOfNode(heap, index))
	}
	traverse(0) // Start traversal from root
	return result
}

// The RecursiveAllElementsLessThanK function recursively prints all elements in a heap that are less
// than a given value.
func RecursiveAllElementsLessThanK(heap *Heap, k, index int) {
	if index >= len(heap.elements) || index < 0 {
		return
	}
	fmt.Print(colorMagenta)
	if index >= 0 && heap.elements[index] < k {
		fmt.Printf("\t%d", heap.elements[index])
	}
	fmt.Print(colorReset)
	RecursiveAllElementsLessThanK(heap, k, LeftChildOfNode(heap, index))
	RecursiveAllElementsLessThanK(heap, k, RightChildOfNode(heap, index))
}

// The RecursiveAllElementsGreaterThanK function recursively prints all elements in a heap that are
// greater than a given value.
func RecursiveAllElementsGreaterThanK(heap *Heap, k, index int) {
	if index >= len(heap.elements) || index < 0 {
		return
	}

	fmt.Print(colorMagenta)
	if index >= 0 && heap.elements[index] > k {
		fmt.Printf("\t%d", heap.elements[index])
	}
	fmt.Print(colorReset)
	RecursiveAllElementsGreaterThanK(heap, k, LeftChildOfNode(heap, index))
	RecursiveAllElementsGreaterThanK(heap, k, RightChildOfNode(heap, index))
}

func MergeMaxHeaps(h1, h2 *Heap) {
	h2.elements = append(h2.elements, h1.elements...)
	h2.Print()
	length := len(h2.elements)
	fmt.Println((length - 1) / 2)
	PercolateDown(h2, (length-1)/2)
}
