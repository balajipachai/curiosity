package main

import (
	"fmt"

	"example.com/doublylist"
	"example.com/singlylist"
)

const (
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorReset  = "\033[0m"
	colorCyan   = "\033[36m"
)

// This function performs various operations on a singly linked list, including inserting and deleting
// elements and printing the list length.
func executeSinglyLinkedListOperations() {
	singlyLinkedList := singlylist.SinglyLinkedList{}
	fmt.Println(colorCyan + "\nInsert 10, 20, 30, 40 & 50 in the list\n" + colorReset)

	// Insert
	singlyLinkedList.Insert(10, 1)
	// The `LastNode` function is used to find and return the last node in a doubly linked list.
	singlyLinkedList.Insert(20, 2)
	singlyLinkedList.Insert(30, 3)
	singlyLinkedList.Insert(40, 4)
	singlyLinkedList.Insert(50, 5)

	fmt.Println(colorCyan + "\nPrint the list elements:" + colorReset)
	fmt.Printf(colorYellow+"\nListLength: %d\t\n", singlyLinkedList.ListLength())
	fmt.Println(colorReset)

	fmt.Println(colorRed + "\nDelete 10 from the list" + colorReset)
	singlyLinkedList.Delete(1)
	fmt.Println(colorCyan + "\nPrint the list elements after delete:" + colorReset)
	fmt.Printf(colorYellow+"\nListLength: %d\t\n", singlyLinkedList.ListLength())
	fmt.Println(colorReset)

	fmt.Println(colorRed + "\nDelete 50 from the list" + colorReset)
	singlyLinkedList.Delete(4)
	fmt.Println(colorCyan + "\nPrint the list elements after delete:" + colorReset)
	fmt.Printf(colorYellow+"\nListLength: %d\t\n", singlyLinkedList.ListLength())
	fmt.Println(colorReset)

	fmt.Println(colorRed + "\nDelete 30 from the list" + colorReset)
	singlyLinkedList.Delete(2)
	fmt.Println(colorCyan + "\nPrint the list elements after delete:" + colorReset)
	fmt.Printf(colorYellow+"\nListLength: %d\t\n", singlyLinkedList.ListLength())
	fmt.Println(colorReset)
}

func executeDoublyLinkedListOperations() {
	doublyLinkedList := doublylist.DoublyLinkedList{}

	fmt.Println(colorCyan + "\nInsert 60, 70, 80, 90 & 100 in the list\n" + colorReset)

	// Insert one after the other
	doublyLinkedList.Insert(60, 1)
	doublyLinkedList.Insert(70, 2)
	doublyLinkedList.Insert(80, 3)
	doublyLinkedList.Insert(90, 4)
	doublyLinkedList.Insert(100, 5)
	// Insert in the middle
	doublyLinkedList.Insert(65, 2)

	fmt.Println(colorCyan + "\nPrint the list elements in REVERSE order:" + colorReset)
	doublyLinkedList.PrintReverseList()
	fmt.Println(colorReset)

	fmt.Println(colorCyan + "\nPrint the list elements:" + colorReset)
	fmt.Printf(colorYellow+"\nListLength: %d\t\n", doublyLinkedList.ListLength())
	fmt.Println(colorReset)

	fmt.Println(colorRed + "\nDelete 60 from the list" + colorReset)
	doublyLinkedList.Delete(1)
	fmt.Println(colorCyan + "\nPrint the list elements after delete:" + colorReset)
	fmt.Printf(colorYellow+"\nListLength: %d\t\n", doublyLinkedList.ListLength())
	fmt.Println(colorReset)

	fmt.Println(colorCyan + "\nPrint the list elements in REVERSE order:" + colorReset)
	doublyLinkedList.PrintReverseList()
	fmt.Println(colorReset)

	fmt.Println(colorRed + "\nDelete 80 from the list" + colorReset)
	doublyLinkedList.Delete(3)
	fmt.Println(colorCyan + "\nPrint the list elements after delete:" + colorReset)
	fmt.Printf(colorYellow+"\nListLength: %d\t\n", doublyLinkedList.ListLength())
	fmt.Println(colorReset)

	fmt.Println(colorCyan + "\nPrint the list elements in REVERSE order:" + colorReset)
	doublyLinkedList.PrintReverseList()
	fmt.Println(colorReset)

	fmt.Println(colorRed + "\nDelete 100 from the list" + colorReset)
	doublyLinkedList.Delete(4)
	fmt.Println(colorCyan + "\nPrint the list elements after delete:" + colorReset)
	fmt.Printf(colorYellow+"\nListLength: %d\t\n", doublyLinkedList.ListLength())
	fmt.Println(colorReset)

	fmt.Println(colorCyan + "\nPrint the list elements in REVERSE order:" + colorReset)
	doublyLinkedList.PrintReverseList()
	fmt.Println(colorReset)
}

// The main function calls and executes the singly linked list operations.
func main() {
	executeSinglyLinkedListOperations()
	executeDoublyLinkedListOperations()
}
