package main

import (
	"fmt"

	"example.com/doublylists"
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

	fmt.Print("\n")
	fmt.Println("*******************Doubly Linked List*******************")
	dLinkedList := doublylists.LinkedList{}

	dLinkedList.AddToHead(10)
	dLinkedList.AddToHead(20)
	dLinkedList.AddToHead(30)
	dLinkedList.AddToHead(40)
	dLinkedList.AddToHead(50)

	dLinkedList.IterateList()

	fmt.Println("\nDoubly List's LastNode:")
	fmt.Println(dLinkedList.LastNode())

	fmt.Println("Adds 100 to end of the list:")
	dLinkedList.AddToEnd(100)
	fmt.Println("After adding 100 to end of the list, printing the List's LastNode:")
	fmt.Println(dLinkedList.LastNode())
	fmt.Println("Chekcs and prints a node with value 50:")
	fmt.Println(dLinkedList.NodeWithValue(50))
	fmt.Println("Adds 90 after node having value 30:")
	dLinkedList.AddAfter(30, 90)
	fmt.Println("Prints the list after adding 90 after 30:")
	dLinkedList.IterateList()
}
