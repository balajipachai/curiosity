package stack

// The code defines a Stack data structure in Go.
// @property {[]int} element - The "element" property is a slice of integers, which will
// be used to store the elements of the stack.
type Stack struct {
	Element []int
}

/*
- Notes: The stack works in Last-In-First-Out (LIFO) manner.
	- The methods on the stack include Push, Pop, & Top
	- Push - It adds an element to the end of the stack.
	- Pop - It removes an element from the top of the stack.
	- Top - It returns the topmost element from the stack.
*/

// The `func (stack *Stack) New()` function is initializing a new instance of the Stack data structure.
// It creates a new slice of integers with a length of 0 using the `make()` function and assigns it to
// the `element` property of the Stack object. This function is called when a new Stack object is
// created to ensure that the `element` property is initialized to an empty slice.
func (stack *Stack) New() *Stack {
	stack.Element = make([]int, 0)
	return stack
}

// The `Top()` function is implementing the `Top` method of the Stack data structure. It returns the
// topmost element from the stack without removing it. It does this by accessing the last element of
// the `element` slice using the `len()` function and subtracting 1 to get the index of the last
// element. It then returns this last element.
func (stack *Stack) Top() int {
	return stack.Element[len(stack.Element)-1]
}

// This function is implementing the Push method of the Stack data structure. It adds an element to the
// end of the stack by appending it to the slice of integers that represents the stack's elements. The
// function takes an integer argument called "element" which is the value to be added to the stack. The
// "stack" parameter is a pointer to the Stack object on which the method is called, and it allows the
// function to modify the object's "element" property.
func (stack *Stack) Push(element int) {
	stack.Element = append(stack.Element, element)
}

// This function is implementing the Pop method of the Stack data structure. It removes the topmost
// element from the stack and returns it. It first checks if the stack is empty, and if it is, it
// returns 0. Otherwise, it retrieves the last element from the stack, removes it from the slice using
// the append function, and returns it.
func (stack *Stack) Pop() int {
	stackElementLength := len(stack.Element)
	// Empty Stack
	if stackElementLength == 0 {
		return 0
	}
	lastElement := stack.Element[stackElementLength-1]
	stack.Element = stack.Element[:stackElementLength-1]
	return lastElement
}
