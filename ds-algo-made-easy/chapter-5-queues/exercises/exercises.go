package exercises

import (
	"fmt"

	"example.com/queue"
	"example.com/stacks"
)

const (
	colorCyan   = "\033[36m"
	colorReset  = "\033[0m"
	colorYellow = "\033[33m"
)

// The function `ReverseQueue` takes a dynamic queue as input and reverses its contents using a stack.
func ReverseQueue(dQueue *queue.QueueDynamicArray) {
	stack := stacks.DynamicStack{}
	stack.CreateNew()
	/*
		Algorithm:
		1. Make use a Stack: S
		2. While the Queue `Q` is not empty, DeQueue `Q` & Push into `S`
		3. While the Stack `S` is not empty, Pop from the Stack & EnQueue `Q`

		Thus, we have reversed the contents of the Queue:
		DeQueing happens from the front, i.e. the first element, we do DeQueue and add it to the Stack, thus, Queue's first element
		becomes Stack's last element and we go on pushing to the stack until the queue is empty. Hence, ensuring the Q's last element ins Stacks's top element.
		As Stack's top element is Q's last element, then we do EnQueue until the Stack is not empty.
		Thus, we have reversed the Queue using Q's Enqueue & DeQueue methods.
	*/
	for !dQueue.IsEmpty() {
		fmt.Println(colorCyan + "\tDeQueue & Push into Stack" + colorReset)
		stack.Push(dQueue.DeQueue())
	}

	for !stack.IsEmpty() {
		fmt.Println(colorCyan + "\tPop from Stack & Enqueue in Queue" + colorReset)
		dQueue.EnQueue(stack.Pop())
	}
}

// The function implements a queue using two stacks in the Go programming language.
func QueueImplementationUsingStack(elements []int) {
	S1 := stacks.DynamicStack{}
	S1.CreateNew()

	S2 := stacks.DynamicStack{}
	S2.CreateNew()

	/*
		Algorithm:
		In order to implement a queue, we must define the EnQueu & DeQueue Operations.
		We can achieve this by making use of two Stacks, S1 & S2
		The EnQueue operation is easy, push into S1
		For DeQueue operation:
			If S2 is not empty then, Pop from it
			If S2 is empty then, Move all elements from S1 to S2 and then Pop the top element

		Using this approach a queue can be implemented using Stacks
	*/
	// EnQueue
	for _, v := range elements {
		S1.Push(v)
	}

	// DeQueue
	if !S2.IsEmpty() {
		fmt.Println(colorCyan, "\t", S2.Pop(), colorReset)
	} else {
		for !S1.IsEmpty() {
			S2.Push(S1.Pop())
		}
		// Print the Queue Elements
		fmt.Println(colorCyan)
		for !S2.IsEmpty() {
			fmt.Printf("\t%d", S2.Pop())
		}
		fmt.Println("\n" + colorReset)
	}
}

// The function transfers elements from a queue to a stack while preserving the order of the elements.
func TransferQElementsToStackPreservingOrder(dQueue *queue.QueueDynamicArray) {
	S1 := stacks.DynamicStack{}
	S1.CreateNew()

	// Push Q elements into the stack, this will have the an element at the top & a1 element at the bottom
	for !dQueue.IsEmpty() {
		S1.Push(dQueue.DeQueue())
	}

	fmt.Println("\t"+colorYellow+"Print Stack's Top element", colorReset, colorCyan, S1.TopElement(), colorReset)
	fmt.Println("\t"+colorYellow+"Print Stack element", colorReset)
	S1.Print(false)

	// Pop all the elements from the stack, and enqueu to the Q
	for !S1.IsEmpty() {
		dQueue.EnQueue(S1.Pop())
	}

	// Again Push to The Stack, then top of the stack will be 1
	for !dQueue.IsEmpty() {
		S1.Push(dQueue.DeQueue())
	}

	fmt.Println("\t"+colorYellow+"Print Stack's Top element", colorReset, colorCyan, S1.TopElement(), colorReset)
	fmt.Println("\t"+colorYellow+"Print Stack element", colorReset)
	S1.Print(false)
}
