package stacks

import "fmt"

const (
	STACK_SIZE   = 10
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
)

// The `StackFixedArray` type represents a stack data structure with a fixed capacity of 10 elements.
// @property {int} top - The "top" property represents the index of the top element in the stack.
// It keeps track of the position where the next element will be inserted or removed from the stack.
// @property {int} capacity - The capacity property represents the maximum number of elements that the
// stack can hold. In this case, the stack has a capacity of 10, meaning it can hold up to 10 elements.
// @property elements - The `elements` property is an array of integers with a fixed size of 10.
type StackFixedArray struct {
	top      int
	capacity int
	elements [STACK_SIZE]int
}

// The `CreateNew()` function is initializing a new instance of the `StackFixedArray` struct. It sets
// the `top` property to -1, indicating that the stack is empty, and sets the `capacity` property to
// 10, indicating that the stack can hold up to 10 elements.
func (stack *StackFixedArray) CreateNew() {
	stack.top = -1
	stack.capacity = 10
}

// The `IsFull()` function is checking whether the stack is full or not. It does this by comparing the
// `top` property of the stack (which represents the index of the top element) with the `capacity`
// property (which represents the maximum number of elements the stack can hold). If the `top` property
// is equal to the `capacity` property, it means that the stack is full and the function returns
// `true`. Otherwise, it returns `false`.
func (stack *StackFixedArray) IsFull() bool {
	return stack.top == (stack.capacity - 1)
}

// The `IsEmpty()` function is checking whether the stack is empty or not. It does this by comparing
// the `top` property of the stack (which represents the index of the top element) with -1. If the
// `top` property is equal to -1, it means that the stack is empty and the function returns `true`.
// Otherwise, it returns `false`.
func (stack *StackFixedArray) IsEmpty() bool {
	return stack.top == -1
}

// The `Push` function is used to add an element to the top of the stack. It takes an integer `data` as
// a parameter.
func (stack *StackFixedArray) Push(data int) {
	if stack.IsFull() {
		fmt.Println(colorRed + "\tOverflow: Stack is full" + colorReset)
		return
	}
	stack.top += 1
	stack.elements[stack.top] = data
}

// The `Pop()` function is used to remove and return the top element from the stack.
func (stack *StackFixedArray) Pop() int {
	if stack.IsEmpty() {
		fmt.Println(colorRed + "\tUndeflow: Stack is empty" + colorReset)
		return 0
	}
	poppedElement := stack.elements[stack.top]
	stack.elements[stack.top] = 0
	stack.top -= 1
	return poppedElement
}

// The `TopElement()` function is returning the value of the top element in the stack without removing
// it. It does this by accessing the `elements` array of the stack and retrieving the element at the
// index specified by the `top` property.
func (stack *StackFixedArray) TopElement() int {
	return stack.elements[stack.top]
}

// The `Capacity()` function is a method of the `StackFixedArray` struct. It returns the value of the
// `capacity` property of the stack, which represents the maximum number of elements that the stack can
// hold. This function allows you to retrieve the capacity of the stack.
func (stack *StackFixedArray) Capacity() int {
	return stack.capacity
}

// The `Print()` function is a method of the `StackFixedArray` struct. It is used to print the elements
// of the stack in a formatted way.
func (stack *StackFixedArray) Print() {
	for i := range stack.elements {
		fmt.Printf(colorMagenta+"\t%d\t", stack.elements[i])
	}
	fmt.Printf("\n" + colorReset)
}

// ====================================================================================================================

// The DynamicStack type is a stack data structure that can dynamically resize its underlying array to
// accommodate more elements.
// @property {int} top - The "top" property represents the index of the top element in the stack. It
// keeps track of the position where the next element will be added or removed from the stack.
// @property {[]int} elements - The `elements` property is a slice of integers that represents the
// elements stored in the dynamic stack.
type DynamicStack struct {
	top      int
	elements []int
}

// The `CreateNew()` function is a method of the `DynamicStack` struct. It is used to initialize a new
// instance of the `DynamicStack` struct.
func (dynamicStack *DynamicStack) CreateNew() {
	dynamicStack.top = -1
	dynamicStack.elements = make([]int, 0) // Assigns memory to dynamic slice
}

// The `IsEmpty()` function is a method of the `DynamicStack` struct. It is used to check whether the
// dynamic stack is empty or not.
func (dynamicStack *DynamicStack) IsEmpty() bool {
	return dynamicStack.top == -1
}

// The `Push` function is used to add an element to the top of the dynamicStack. It takes an integer `data` as
// a parameter.
func (dynamicStack *DynamicStack) Push(data int) {
	dynamicStack.top += 1
	dynamicStack.elements = append(dynamicStack.elements, data)
}

// The `Pop()` function is used to remove and return the top element from the dynamicStack.
func (dynamicStack *DynamicStack) Pop() int {
	if dynamicStack.IsEmpty() {
		fmt.Println(colorRed + "\tUndeflow: Stack is empty" + colorReset)
		return 0
	}
	poppedElement := dynamicStack.elements[dynamicStack.top]
	dynamicStack.elements = dynamicStack.elements[:dynamicStack.top]
	dynamicStack.top -= 1
	return poppedElement
}

// The `TopElement()` function is returning the value of the top element in the dynamicStack without removing
// it. It does this by accessing the `elements` array of the dynamicStack and retrieving the element at the
// index specified by the `top` property.
func (dynamicStack *DynamicStack) TopElement() int {
	return dynamicStack.elements[dynamicStack.top]
}

// The `Print()` function is a method of the `StackFixedArray` struct. It is used to print the elements
// of the stack in a formatted way.
func (dynamicStack *DynamicStack) Print(printAsChar bool) {
	for i := range dynamicStack.elements {
		if printAsChar {
			fmt.Printf(colorMagenta+"\t%c\t", dynamicStack.elements[i])
		} else {
			fmt.Printf(colorMagenta+"\t%d\t", dynamicStack.elements[i])
		}
	}
	fmt.Printf("\n" + colorReset)
}

// ====================================================================================================================
// The LinkedListStack type represents a stack implemented using a linked list.
// @property {int} data - The "data" property represents the value stored in each node of the linked
// list stack. It can be of any data type, but in this case, it is specified as an integer (int).
// @property nextNode - nextNode is a pointer to the next node in the linked list. It is of type
// *LinkedListStack, which means it points to another node of the same type.
type LinkedListStackNode struct {
	data     int
	nextNode *LinkedListStackNode
}

// The LinkedListStack type represents a stack data structure implemented using a linked list.
// @property headNode - The headNode property is a pointer to the first node in the linked list stack.
type LinkedListStack struct {
	headNode *LinkedListStackNode
}

// The `IsEmpty()` function is a method of the `LinkedListStack` struct. It is used to check whether
// the linked list stack is empty or not.
func (linkedListStack *LinkedListStack) IsEmpty() bool {
	return linkedListStack.headNode == nil
}

// The `Push` function in the `LinkedListStack` struct is used to add an element to the top of the
// linked list stack.
func (linkedListStack *LinkedListStack) Push(data int) {
	headNode := linkedListStack.headNode

	newNode := &LinkedListStackNode{}
	newNode.data = data

	newNode.nextNode = headNode
	linkedListStack.headNode = newNode
}

// The `Pop()` function in the `LinkedListStack` struct is used to remove and return the top element
// from the linked list stack.
func (linkedListStack *LinkedListStack) Pop() int {
	if linkedListStack.IsEmpty() {
		fmt.Println(colorRed + "\tUndeflow: Stack is empty" + colorReset)
		return 0
	}
	headNode := linkedListStack.headNode
	data := headNode.data
	if headNode.nextNode != nil {
		nextNode := headNode.nextNode
		headNode.nextNode = nil
		linkedListStack.headNode = nextNode
	} else {
		// there is only 1 node
		linkedListStack.headNode = nil
	}
	return data
}

// The `Print()` function is a method of the `LinkedListStack` struct. It is used to print the elements
// of the linked list stack in a formatted way.
func (linkedListStack *LinkedListStack) Print() {
	for node := linkedListStack.headNode; node != nil; node = node.nextNode {
		fmt.Printf(colorMagenta+"\t%d\t", node.data)
	}
	fmt.Printf("\n" + colorReset)
}

// The `TopElement()` function is a method of the `LinkedListStack` struct. It is used to retrieve the
// value of the top element in the linked list stack without removing it.
func (linkedListStack *LinkedListStack) TopElement() int {
	return linkedListStack.headNode.data
}
