package queue

import "fmt"

const (
	QUEUE_SIZE   = 10
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
)

// ------------------------------------------ FIXED SIZE QUEUE IMPLEMENTATION ---------------------------------------------------

// The type QueueFixedArray represents a fixed-size queue implemented using an array.
// @property {int} front - The front property represents the index of the first element in the queue.
// @property {int} rear - The "rear" property in the QueueFixedArray struct represents the index of the
// last element in the queue.
// @property {int} capacity - The capacity property represents the maximum number of elements that the
// queue can hold. It determines the size of the underlying fixed-size array used to store the
// elements.
// @property elements - The `elements` property is an array that stores the elements of the queue. It
// has a fixed size of `QUEUE_SIZE`, which is a constant value.
type QueueFixedArray struct {
	front    int
	rear     int
	capacity int
	elements [QUEUE_SIZE]int
}

// The `CreateNew()` function is a method of the `QueueFixedArray` struct. It is used to initialize a
// new instance of the queue.
func (queue *QueueFixedArray) CreateNew() {
	queue.front = -1
	queue.rear = -1
	queue.capacity = QUEUE_SIZE
}

// The `IsFull()` function is a method of the `QueueFixedArray` struct. It is used to check if the
// queue is full or not.
func (queue *QueueFixedArray) IsFull() bool {
	return ((queue.rear + 1) % queue.capacity) == queue.front
}

// The `IsEmpty()` function is a method of the `QueueFixedArray` struct. It is used to check if the
// queue is empty or not.
func (queue *QueueFixedArray) IsEmpty() bool {
	return queue.front == -1
}

// The `EnQueue` function is used to add an element to the queue. It takes an integer `data` as a
// parameter.
func (queue *QueueFixedArray) EnQueue(data int) {
	if queue.IsFull() {
		fmt.Println(colorRed + "\tOverflow: Queue is full" + colorReset)
		return
	}
	queue.rear = (queue.rear + 1) % queue.capacity
	queue.elements[queue.rear] = data
	if queue.front == -1 {
		queue.front = queue.rear
	}
}

// The `DeQueue()` function is a method of the `QueueFixedArray` struct. It is used to remove and
// return the element at the front of the queue.
func (queue *QueueFixedArray) DeQueue() int {
	data := 0
	if queue.IsEmpty() {
		fmt.Println(colorRed + "\tUndeflow: Queue is empty" + colorReset)
		return 0
	}
	data = queue.elements[queue.front]
	queue.elements[queue.front] = 0
	if queue.front == queue.rear {
		queue.front = -1
		queue.rear = -1
	} else {
		queue.front = (queue.front + 1) % queue.capacity
	}
	return data
}

// The `FrontElement()` function is a method of the `QueueFixedArray` struct. It is used to retrieve
// the element at the front of the queue without removing it.
func (queue *QueueFixedArray) FrontElement() int {
	return queue.elements[queue.front]
}

// The `RearElement()` function is a method of the `QueueFixedArray` struct. It is used to retrieve the
// element at the rear of the queue without removing it.
func (queue *QueueFixedArray) RearElement() int {
	return queue.elements[queue.rear]
}

// The `Print()` function is a method of the `QueueFixedArray` struct. It is used to print the elements
// of the queue.
func (queue *QueueFixedArray) Print() {
	for i := range queue.elements {
		fmt.Printf(colorMagenta+"\t%d\t", queue.elements[i])
	}
	fmt.Printf("\n" + colorReset)
}

// ------------------------------------------ DYNAMIC SIZE QUEUE IMPLEMENTATION ---------------------------------------------------

// The `QueueDynamicArray` type represents a queue implemented using a dynamic array.
// @property {int} front - The front property represents the index of the front element in the queue.
// @property {int} rear - The "rear" property represents the index of the last element in the queue.
// @property {[]int} elements - The `elements` property is an array that stores the elements of the queue.
type QueueDynamicArray struct {
	front    int
	rear     int
	elements []int
}

// The `CreateNew()` function is a method of the `QueueDynamicArray` struct. It is used to initialize a
// new instance of the queue implemented using a dynamic array.
func (dQueue *QueueDynamicArray) CreateNew() {
	dQueue.front = -1
	dQueue.rear = -1
	dQueue.elements = make([]int, 0)
}

// The `IsEmpty()` function is a method of the `QueueDynamicArray` struct. It is used to check if the
// queue implemented using a dynamic array is empty or not. It returns a boolean value indicating
// whether the front index of the queue is -1 or not. If the front index is -1, it means that the queue
// is empty.
func (dQueue *QueueDynamicArray) IsEmpty() bool {
	return dQueue.front == -1
}

// The `EnQueue` function is a method of the `QueueDynamicArray` struct. It is used to add an element
// to the queue implemented using a dynamic array.
func (dQueue *QueueDynamicArray) EnQueue(data int) {
	// No need to check IsFull as it's a dynamic array
	dQueue.elements = append(dQueue.elements, data)
	dQueue.rear = len(dQueue.elements) - 1

	if dQueue.front == -1 {
		dQueue.front = dQueue.rear
	}
}

// The `DeQueue()` function is a method of the `QueueDynamicArray` struct. It is used to remove and
// return the element at the front of the queue implemented using a dynamic array.
func (dQueue *QueueDynamicArray) DeQueue() int {
	if dQueue.IsEmpty() {
		fmt.Println(colorRed + "\tUndeflow: Queue is empty" + colorReset)
		return 0
	}
	data := dQueue.elements[dQueue.front]
	dQueue.elements = dQueue.elements[dQueue.front+1:]
	if dQueue.front == dQueue.rear {
		dQueue.front = -1
		dQueue.rear = -1
	} else {
		dQueue.front = 0
		dQueue.rear -= 1
	}
	return data
}

// The `FrontElement()` function is a method of the `QueueDynamicArray` struct. It is used to retrieve
// the element at the front of the queue implemented using a dynamic array without removing it.
func (dQueue *QueueDynamicArray) FrontElement() int {
	return dQueue.elements[dQueue.front]
}

// The `RearElement()` function is a method of the `QueueDynamicArray` struct. It is used to retrieve
// the element at the rear of the queue implemented using a dynamic array without removing it.
func (dQueue *QueueDynamicArray) RearElement() int {
	return dQueue.elements[dQueue.rear]
}

// The `Print()` function is a method of the `QueueDynamicArray` struct. It is used to print the
// elements of the queue implemented using a dynamic array.
func (dQueue *QueueDynamicArray) Print() {
	for i := range dQueue.elements {
		fmt.Printf(colorMagenta+"\t%d\t", dQueue.elements[i])
	}
	fmt.Printf("\n" + colorReset)
}

// ------------------------------------------ LINKED LIST QUEUE IMPLEMENTATION ---------------------------------------------------

// The LinkedListQueueNode type represents a node in a linked list queue, with an integer data value
// and a reference to the next node.
// @property {int} data - The data property is an integer that represents the value stored in the node
// of the linked list.
// @property nextNode - The `nextNode` property is a pointer to the next node in the linked list. It is
// used to connect the current node to the next node in the queue.
type LinkedListQueueNode struct {
	data     int
	nextNode *LinkedListQueueNode
}

// The LinkedListQueue type represents a queue implemented using a linked list.
// @property headNode - The headNode property is a pointer to the first node in the linked list queue.
type LinkedListQueue struct {
	headNode *LinkedListQueueNode
}

/*
Notes:
Queue implementation using linked list is:
EnQueue is AddToEnd (add the node in the last)
DeQueue is RemoveFromHead (remove the first node)
*/

// The `IsEmpty()` function is a method of the `LinkedListQueue` struct. It is used to check if the
// linked list queue is empty or not. It does this by checking if the `headNode` property of the
// `LinkedListQueue` struct is `nil`. If the `headNode` is `nil`, it means that the queue is empty and
// the function returns `true`. Otherwise, it returns `false`.
func (linkedListQueue *LinkedListQueue) IsEmpty() bool {
	return linkedListQueue.headNode == nil
}

// The `EnQueue()` function is a method of the `LinkedListQueue` struct. It is used to add an element
// to the queue implemented using a linked list.
func (linkedListQueue *LinkedListQueue) EnQueue(data int) {
	newNode := &LinkedListQueueNode{}
	newNode.data = data

	if linkedListQueue.IsEmpty() {
		// Insert the first node
		newNode.nextNode = linkedListQueue.headNode
		linkedListQueue.headNode = newNode
	} else {
		// Traverse the node, till the last one and then insert the new node
		for node := linkedListQueue.headNode; node != nil; node = node.nextNode {
			// This implies node is the lastNode
			if node.nextNode == nil {
				node.nextNode = newNode
				break
			}
		}
	}
}

// The `DeQueue()` function is a method of the `LinkedListQueue` struct. It is used to remove and
// return the element at the front of the queue implemented using a linked list.
func (linkedListQueue *LinkedListQueue) DeQueue() *LinkedListQueueNode {
	if linkedListQueue.IsEmpty() {
		fmt.Println(colorRed + "\tUndeflow: Queue is empty" + colorReset)
		return nil
	}
	nodeToReturn := linkedListQueue.headNode
	// Update head to point to the next node
	linkedListQueue.headNode = nodeToReturn.nextNode
	return nodeToReturn
}

// The `FrontElement()` function is a method of the `LinkedListQueue` struct. It is used to retrieve
// the element at the front of the queue implemented using a linked list without removing it.
func (linkedListQueue *LinkedListQueue) FrontElement() *LinkedListQueueNode {
	if !linkedListQueue.IsEmpty() {
		return linkedListQueue.headNode
	}
	return nil
}

// The `RearElement()` function is a method of the `LinkedListQueue` struct. It is used to retrieve the
// element at the rear of the queue implemented using a linked list without removing it.
func (linkedListQueue *LinkedListQueue) RearElement() *LinkedListQueueNode {
	if !linkedListQueue.IsEmpty() {
		for node := linkedListQueue.headNode; node != nil; node = node.nextNode {
			if node.nextNode == nil {
				return node
			}
		}
	}
	return nil
}

// The `Print()` function is a method of the `LinkedListQueue` struct. It is used to print the elements
// of the queue implemented using a linked list.
func (linkedListQueue *LinkedListQueue) Print() {
	for node := linkedListQueue.headNode; node != nil; node = node.nextNode {
		fmt.Printf(colorMagenta+"\t%d\t", node.data)
	}
	fmt.Printf("\n" + colorReset)
}
