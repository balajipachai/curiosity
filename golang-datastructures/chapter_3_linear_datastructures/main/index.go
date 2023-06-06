package main

import (
	"fmt"

	"example.com/lists"
)

// The main function creates a singly linked list, adds nodes to it, and performs various operations on it
// such as finding the last node, adding a node to the end, finding a node with a specific value, and
// adding a node after a specific node.
func main() {
	fmt.Println("*******************Singly Linked List*******************")
	linkedList := lists.LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(2)
	linkedList.AddToHead(3)
	linkedList.AddToHead(4)
	linkedList.AddToHead(5)

	linkedList.IterateList()

	fmt.Println("\nList's LastNode:")
	fmt.Println(linkedList.LastNode())

	fmt.Println("Adds 10 to end of the list:")
	linkedList.AddToEnd(10)
	fmt.Println("After adding 10 to end of the list, printing the List's LastNode:")
	fmt.Println(linkedList.LastNode())
	fmt.Println("Chekcs and prints a node with value 2:")
	fmt.Println(linkedList.NodeWithValue(2))
	fmt.Println("Adds 9 after node having value 3:")
	linkedList.AddAfter(3, 9)
	fmt.Println("Prints the list after adding 9 after 3:")
	linkedList.IterateList()
}
