package main

import (
	"fmt"
	"time"

	"example.com/stacks"
)

const (
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorReset  = "\033[0m"
	colorCyan   = "\033[36m"
	colorGreen  = "\033[32m"
	colorBlue   = "\033[34m"
)

func printDottedLine() {
	fmt.Println(colorBlue + "---------------------------------------------------------------------------------------------------------------------------------------------------------------" + colorReset)
}

func executeFixedSizeStackOperations() {
	fmt.Println("FIXED SIZE STACK")
	stack := stacks.StackFixedArray{}

	// Create a New Stack
	stack.CreateNew()

	startTime := time.Now()
	// Check if Stack is empty
	fmt.Println("\t"+colorYellow+"Initially checks if stack is empty:\t ", stack.IsEmpty(), colorReset)
	endTime := time.Now()
	fmt.Println("\t"+colorCyan+"Time elapsed for stack.IsEmpty()\t", endTime.Sub(startTime))
	fmt.Println(colorReset)

	startTime = time.Now()
	// Check if Stack is full
	fmt.Println("\t"+colorYellow+"Initially checks if stack is full:\t ", stack.IsFull(), colorReset)
	endTime = time.Now()
	fmt.Println("\t"+colorCyan+"Time elapsed for stack.IsFull()\t", endTime.Sub(startTime))
	fmt.Println(colorReset)

	fmt.Println("\t"+colorYellow+"Push 1, 2, 3, 4, & 5 into the stack", colorReset)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	fmt.Println("\t"+colorYellow+"Print Stack elements", colorReset)
	stack.Print()
	fmt.Println("\t"+colorCyan+"Prints the top element:\t", stack.TopElement(), colorReset)

	fmt.Println("\t"+colorYellow+"Pop should remove 5 from the stack", colorReset)
	stack.Pop()
	fmt.Println("\t"+colorCyan+"Prints the top element:\t", stack.TopElement(), colorReset)
	stack.Print()

	fmt.Println("\t"+colorYellow+"Should overflow", colorReset)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Print()
	stack.Push(7)

	fmt.Println("\t"+colorYellow+"Should underflow", colorReset)
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()

	printDottedLine()
}

func executeDynamicStackOperations() {
	fmt.Println("DYNAMIC STACK USING SLICES")
	dStack := stacks.DynamicStack{}

	// Create a New Stack
	dStack.CreateNew()

	startTime := time.Now()
	// Check if Stack is empty
	fmt.Println("\t"+colorYellow+"Initially checks if stack is empty:\t ", dStack.IsEmpty(), colorReset)
	endTime := time.Now()
	fmt.Println("\t"+colorCyan+"Time elapsed for stack.IsEmpty()\t", endTime.Sub(startTime))
	fmt.Println(colorReset)

	fmt.Println("\t"+colorYellow+"Push 10, 20, 30, 40, & 50 into the stack", colorReset)
	dStack.Push(10)
	dStack.Push(20)
	dStack.Push(30)
	dStack.Push(40)
	dStack.Push(50)
	fmt.Println("\t"+colorYellow+"Print Stack elements", colorReset)
	dStack.Print()
	fmt.Println("\t"+colorCyan+"Prints the top element:\t", dStack.TopElement(), colorReset)

	fmt.Println("\t"+colorYellow+"Pop should remove 50 from the stack", colorReset)
	dStack.Pop()
	fmt.Println("\t"+colorCyan+"Prints the top element:\t", dStack.TopElement(), colorReset)
	dStack.Print()

	fmt.Println("\t"+colorYellow+"Should underflow", colorReset)
	dStack.Pop()
	dStack.Pop()
	dStack.Pop()
	dStack.Pop()
	dStack.Pop()

	printDottedLine()

}

func main() {
	executeFixedSizeStackOperations()
	executeDynamicStackOperations()
}
