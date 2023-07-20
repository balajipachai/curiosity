package main

import (
	"fmt"
	"time"

	"example.com/exercises"
	"example.com/queue"
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

// The function executes various operations on a fixed-size queue and prints the results.
func executeFixedSizeQueueOperations() {
	fmt.Println("FIXED SIZE QUEUE IMPLEMENTATION")
	queue := queue.QueueFixedArray{}

	// Create a New Queue
	queue.CreateNew()

	startTime := time.Now()
	// Check if Queue is empty
	fmt.Println("\t"+colorYellow+"Initially checks if queue is empty:\t ", queue.IsEmpty(), colorReset)
	endTime := time.Now()
	fmt.Println("\t"+colorCyan+"Time elapsed for queue.IsEmpty()\t", endTime.Sub(startTime))
	fmt.Println(colorReset)

	startTime = time.Now()
	// Check if Queue is full
	fmt.Println("\t"+colorYellow+"Initially checks if queue is full:\t ", queue.IsFull(), colorReset)
	endTime = time.Now()
	fmt.Println("\t"+colorCyan+"Time elapsed for queue.IsFull()\t", endTime.Sub(startTime))
	fmt.Println(colorReset)

	fmt.Println("\t"+colorYellow+"EnQueue 1, 2, 3, 4, & 5 into the queue", colorReset)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.EnQueue(5)
	fmt.Println("\t"+colorYellow+"Print Queue elements", colorReset)
	queue.Print()
	fmt.Println("\t"+colorCyan+"Prints the front element:\t", queue.FrontElement(), colorReset)
	fmt.Println("\t"+colorCyan+"Prints the rear element:\t", queue.RearElement(), colorReset)

	fmt.Println("\t"+colorYellow+"DeQueue should remove 1 from the queue", colorReset)
	queue.DeQueue()
	fmt.Println("\t"+colorCyan+"Prints the front element:\t", queue.FrontElement(), colorReset)
	fmt.Println("\t" + colorYellow + "Enqueues 7 in the queue" + colorReset)
	queue.EnQueue(7)
	fmt.Println("\t"+colorCyan+"Prints the rear element:\t", queue.RearElement(), colorReset)
	queue.Print()

	fmt.Println("\t"+colorYellow+"Should overflow", colorReset)
	queue.EnQueue(6)
	queue.EnQueue(7)
	queue.EnQueue(8)
	queue.EnQueue(9)
	queue.EnQueue(10)
	queue.Print()
	queue.EnQueue(11)

	fmt.Println("\t"+colorYellow+"Should underflow", colorReset)
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()

	printDottedLine()
}

// The function `executeDynamicSizeQueueOperations()` demonstrates the implementation of a dynamic size
// queue using an array.
func executeDynamicSizeQueueOperations() {
	fmt.Println("DYNAMIC SIZE QUEUE IMPLEMENTATION")
	dQueue := queue.QueueDynamicArray{}

	// Create a New Queue
	dQueue.CreateNew()

	startTime := time.Now()
	// Check if Queue is empty
	fmt.Println("\t"+colorYellow+"Initially checks if dQueue is empty:\t ", dQueue.IsEmpty(), colorReset)
	endTime := time.Now()
	fmt.Println("\t"+colorCyan+"Time elapsed for dQueue.IsEmpty()\t", endTime.Sub(startTime))
	fmt.Println(colorReset)

	fmt.Println("\t"+colorYellow+"EnQueue 1, 2, 3, 4, & 5 into the dQueue", colorReset)
	dQueue.EnQueue(1)
	dQueue.EnQueue(2)
	dQueue.EnQueue(3)
	dQueue.EnQueue(4)
	dQueue.EnQueue(5)
	fmt.Println("\t"+colorYellow+"Print Queue elements", colorReset)
	dQueue.Print()
	fmt.Println("\t"+colorCyan+"Prints the front element:\t", dQueue.FrontElement(), colorReset)
	fmt.Println("\t"+colorCyan+"Prints the rear element:\t", dQueue.RearElement(), colorReset)

	fmt.Println("\t"+colorYellow+"DeQueue should remove 1 from the dQueue", colorReset)
	dQueue.DeQueue()
	fmt.Println("\t"+colorCyan+"Prints the front element:\t", dQueue.FrontElement(), colorReset)
	fmt.Println("\t" + colorYellow + "Enqueues 7 in the dQueue" + colorReset)
	dQueue.EnQueue(7)
	fmt.Println("\t"+colorCyan+"Prints the rear element:\t", dQueue.RearElement(), colorReset)
	fmt.Println("\t"+colorYellow+"Print Queue elements", colorReset)
	dQueue.Print()

	fmt.Println("\t"+colorYellow+"Should underflow", colorReset)
	dQueue.DeQueue()
	dQueue.DeQueue()
	dQueue.DeQueue()
	dQueue.DeQueue()
	dQueue.DeQueue()
	fmt.Println("\t"+colorYellow+"Print Queue elements", colorReset)
	dQueue.Print()
	dQueue.DeQueue()
	printDottedLine()
}

func executeLinkedListQueueOperations() {
	fmt.Println("LINKED LIST STACK IMPLEMENTATION")
	llQueue := queue.LinkedListQueue{}

	startTime := time.Now()
	// Check if Stack is empty
	fmt.Println("\t"+colorYellow+"Initially checks if queue is empty:\t ", llQueue.IsEmpty(), colorReset)
	endTime := time.Now()
	fmt.Println("\t"+colorCyan+"Time elapsed for queue.IsEmpty()\t", endTime.Sub(startTime))
	fmt.Println(colorReset)

	fmt.Println("\t"+colorYellow+"EnQueue 100, 200, 300, 400, & 500 into the queue", colorReset)
	llQueue.EnQueue(100)
	llQueue.EnQueue(200)
	llQueue.EnQueue(300)
	llQueue.EnQueue(400)
	llQueue.EnQueue(500)
	fmt.Println("\t"+colorYellow+"Print Stack elements", colorReset)
	llQueue.Print()
	fmt.Println("\t"+colorCyan+"Prints the front element:\t", llQueue.FrontElement(), colorReset)
	fmt.Println("\t"+colorCyan+"Prints the rear element:\t", llQueue.RearElement(), colorReset)

	fmt.Println("\t"+colorYellow+"DeQueue should remove 100 from the queue", colorReset)
	fmt.Println(colorCyan, "\t", llQueue.DeQueue(), colorReset)
	fmt.Println("\t"+colorCyan+"Prints the front element:\t", llQueue.FrontElement(), colorReset)
	fmt.Println("\t"+colorCyan+"Prints the rear element:\t", llQueue.RearElement(), colorReset)
	llQueue.Print()

	fmt.Println("\t"+colorYellow+"Should underflow (DeQueueing until queue is empty)", colorReset)
	fmt.Println(colorCyan, "\t", llQueue.DeQueue(), colorReset)
	fmt.Println(colorCyan, "\t", llQueue.DeQueue(), colorReset)
	fmt.Println(colorCyan, "\t", llQueue.DeQueue(), colorReset)
	fmt.Println(colorCyan, "\t", llQueue.DeQueue(), colorReset)
	fmt.Println(colorCyan, "\t", llQueue.DeQueue(), colorReset)

	printDottedLine()

}

func problemOne() {
	dQueue := &queue.QueueDynamicArray{}
	// Create a New Queue
	dQueue.CreateNew()
	fmt.Println("\t"+colorYellow+"EnQueue 1, 2, 3, 4, & 5 into the dQueue", colorReset)
	dQueue.EnQueue(1)
	dQueue.EnQueue(2)
	dQueue.EnQueue(3)
	dQueue.EnQueue(4)
	dQueue.EnQueue(5)
	fmt.Println("\t"+colorYellow+"Print Queue elements", colorReset)
	dQueue.Print()
	fmt.Println("\t"+colorYellow+"Reverse Queue elements", colorReset)
	exercises.ReverseQueue(dQueue)
	fmt.Println("\t"+colorYellow+"Print Queue after reverse", colorReset)
	dQueue.Print()

	printDottedLine()
}

func problemTwo() {
	fmt.Println("\t"+colorYellow+"EnQueue 1, 2, 3, 4, & 5 into the dQueue using Stacks S1 & S2", colorReset)
	elements := []int{1, 2, 3, 4, 5}
	fmt.Println("\t"+colorYellow+"Print the Queue elements:", colorReset)
	exercises.QueueImplementationUsingStack(elements)
	printDottedLine()
}

func problemFive() {
	dQueue := &queue.QueueDynamicArray{}
	// Create a New Queue
	dQueue.CreateNew()
	fmt.Println("\t"+colorYellow+"EnQueue 1, 2, 3, 4, & 5 into the dQueue using Stacks S1 & S2", colorReset)
	dQueue.EnQueue(1)
	dQueue.EnQueue(2)
	dQueue.EnQueue(3)
	dQueue.EnQueue(4)
	dQueue.EnQueue(5)
	fmt.Println("\t"+colorYellow+"Print Queue elements", colorReset)
	dQueue.Print()
	exercises.TransferQElementsToStackPreservingOrder(dQueue)
	printDottedLine()
}

func executeExercises() {
	fmt.Println("QUEUE EXERCISES")
	fmt.Println("Problem 1: Give an algorithm for reversing a queue Q. To access the queue, we are only allowed to use the methods of queue ADT.")
	problemOne()
	fmt.Println("Problem 2: How can you implement a queue using two stacks?")
	problemTwo()
	fmt.Println("Problem 5: Transfer Q items to a Stack, preserving the order of elements in the Q")
	problemFive()
}

func main() {
	executeFixedSizeQueueOperations()
	executeDynamicSizeQueueOperations()
	executeLinkedListQueueOperations()
	executeExercises()
}
