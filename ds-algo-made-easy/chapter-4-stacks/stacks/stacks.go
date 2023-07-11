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
	return stack.top == stack.capacity
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
	stack.top += 1
	if stack.IsFull() {
		fmt.Println(colorRed + "\tOverflow: Stack is full" + colorReset)
		return
	}
	stack.elements[stack.top] = data
}

// The `Pop()` function is used to remove and return the top element from the stack.
func (stack *StackFixedArray) Pop() int {
	stack.top -= 1
	if stack.IsEmpty() {
		fmt.Println(colorRed + "\tUndeflow: Stack is empty" + colorReset)
		return 0
	}
	poppedElement := stack.elements[stack.top]
	stack.elements[stack.top] = 0
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

func (stack *StackFixedArray) Print() {
	for i := range stack.elements {
		fmt.Printf(colorMagenta+"\t%d\t", stack.elements[i])
	}
	fmt.Printf("\n" + colorReset)
}
