package heaps

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
// @property {int} capacity - The capacity property represents the maximum number of elements that the
// heap can hold.
// @property {int} heapType - The heapType property is used to determine whether the heap is a MinHeap or a MaxHeap.
type Heap struct {
	elements []int
	count    int // Number of elements in the heap
	capacity int // Capacity of heap, size of heap
	heapType int // MinHeap = 0, MaxHeap = 1
}

// The `Create` function is a method of the `Heap` struct.
// It is used to initialize a new instance of the `Heap` struct with the given parameters.
func (h *Heap) Create(elements []int, capacity, heapType int) {
	h.elements = make([]int, capacity)
	h.capacity = capacity
	h.count = 0
	h.heapType = heapType
}

// The `ParentOfNode` function is a method of the `Heap` struct.
// It calculates the index of the parent node for a given node in the heap.
func ParentOfNode(h *Heap, location int) int {
	if location <= 0 || location >= h.count {
		return -1
	}
	return (location - 1) / 2
}

// The function returns the index of the left child of a node in a heap.
func LeftChildOfNode(h *Heap, location int) int {
	if location >= h.count {
		return -1
	}
	return (2 * location) + 1
}

// The function returns the index of the right child of a node in a heap.
func RightChildOfNode(h *Heap, location int) int {
	if location >= h.count {
		return -1
	}
	return (2 * location) + 2
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
}

// The Delete function removes the maximum element from a heap and returns it.
func Delete(h *Heap) int {
	if h.count == 0 {
		return -1
	}

	max := h.elements[0]
	h.elements[0] = h.elements[h.count-1]
	h.elements = h.elements[:h.count-1]
	h.count -= 1
	PercolateDown(h, 0)
	return max
}

// The Resize function doubles the capacity of a heap by creating a new array and copying the elements
// from the old array.
func Resize(h *Heap) {
	oldArray := h.elements
	h.elements = make([]int, h.capacity*2)
	copy(h.elements, oldArray)
}

// The Insert function inserts a new element into a heap and ensures that the heap property is
// maintained.
func Insert(h *Heap, data int) {
	if h.count == h.capacity {
		Resize(h)
	}

	h.count += 1
	i := h.count - 1

	// Percolate up
	for i >= 0 && data > h.elements[(i-1)/2] {
		h.elements[i] = h.elements[(i-1)/2]
		i = (i - 1) / 2
	}
	h.elements[i] = data
}
