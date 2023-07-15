package main

import (
	"fmt"
	"strings"
	"time"

	"example.com/exercises"
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
	fmt.Println("FIXED SIZE STACK IMPLEMENTATION")
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
	fmt.Println("DYNAMIC STACK IMPLEMENTATION")
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
	dStack.Print(false)
	fmt.Println("\t"+colorCyan+"Prints the top element:\t", dStack.TopElement(), colorReset)

	fmt.Println("\t"+colorYellow+"Pop should remove 50 from the stack", colorReset)
	dStack.Pop()
	fmt.Println("\t"+colorCyan+"Prints the top element:\t", dStack.TopElement(), colorReset)
	dStack.Print(false)

	fmt.Println("\t"+colorYellow+"Should underflow", colorReset)
	dStack.Pop()
	dStack.Pop()
	dStack.Pop()
	dStack.Pop()
	dStack.Pop()

	printDottedLine()

}

func executeLinkedListStackOperations() {
	fmt.Println("LINKED LIST STACK IMPLEMENTATION")
	llStack := stacks.LinkedListStack{}

	startTime := time.Now()
	// Check if Stack is empty
	fmt.Println("\t"+colorYellow+"Initially checks if stack is empty:\t ", llStack.IsEmpty(), colorReset)
	endTime := time.Now()
	fmt.Println("\t"+colorCyan+"Time elapsed for stack.IsEmpty()\t", endTime.Sub(startTime))
	fmt.Println(colorReset)

	fmt.Println("\t"+colorYellow+"Push 100, 200, 300, 400, & 500 into the stack", colorReset)
	llStack.Push(100)
	llStack.Push(200)
	llStack.Push(300)
	llStack.Push(400)
	llStack.Push(500)
	fmt.Println("\t"+colorYellow+"Print Stack elements", colorReset)
	llStack.Print()
	fmt.Println("\t"+colorCyan+"Prints the top element:\t", llStack.TopElement(), colorReset)

	fmt.Println("\t"+colorYellow+"Pop should remove 500 from the stack", colorReset)
	llStack.Pop()
	fmt.Println("\t"+colorCyan+"Prints the top element:\t", llStack.TopElement(), colorReset)
	llStack.Print()

	fmt.Println("\t"+colorYellow+"Should underflow", colorReset)
	llStack.Pop()
	llStack.Pop()
	llStack.Pop()
	llStack.Pop()
	llStack.Pop()

	printDottedLine()

}

func usingStacksForCheckingBalancingOfSymbols(input string) {
	dStack := stacks.DynamicStack{}
	dStack.CreateNew()
	fmt.Println(colorYellow + "\tInput = \t" + input + colorReset)
	result := exercises.IsSymbolBalanced(&dStack, input)
	if result {
		fmt.Println(colorGreen + "\t SYMBOLS ARE BALANCED" + colorReset)
	} else {
		fmt.Println(colorRed + "\t SYMBOLS ARE NOT BALANCED" + colorReset)
	}
	printDottedLine()
}

func usingStacksForInfixToPostfixExpression(input string) {
	dStack := stacks.DynamicStack{}
	dStack.CreateNew()
	fmt.Println(colorYellow + "\tInput = \t" + input + colorReset)
	result := exercises.InfixToPostfix(&dStack, input)
	fmt.Println(colorGreen + "\t POSTFIX EXPRESSION = \t" + strings.ReplaceAll(result, " ", "") + colorReset)
	printDottedLine()
}

func usingStackForPostfixEvaluation(input string) {
	dStack := stacks.DynamicStack{}
	dStack.CreateNew()
	fmt.Println(colorYellow + "\tInput = \t" + input + colorReset)
	result := exercises.PostfixEvaluation(&dStack, input)
	fmt.Println(colorGreen)
	fmt.Printf("\t POSTFIX EVALUATION RESULT = %d\t", result)
	fmt.Println(colorReset)
	printDottedLine()
}

func executeExercises() {
	fmt.Println("STACK EXERCISES")
	fmt.Println("Problem 1: Using stacks for checking balancing of symbols")
	usingStacksForCheckingBalancingOfSymbols("({[]}) ()")
	usingStacksForCheckingBalancingOfSymbols("()(()[()")
	fmt.Println("Problem 2: Infix to Postfix using stack")
	usingStacksForInfixToPostfixExpression("A * B - (C + D) + E")
	fmt.Println("Problem 3: Postfix evaluation using stacks")
	usingStackForPostfixEvaluation("123*+5-")
}
func main() {
	// executeFixedSizeStackOperations()
	// executeDynamicStackOperations()
	// executeLinkedListStackOperations()
	executeExercises()
}
