package main

import (
	"fmt"
	"time"

	"example.com/circularlist"
	"example.com/doublylist"
	"example.com/problemAndsolutions"
	"example.com/singlylist"
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

// This function performs various operations on a singly linked list, including inserting and deleting
// elements and printing the list length.
func executeSinglyLinkedListOperations() {
	singlyLinkedList := singlylist.SinglyLinkedList{}
	fmt.Println(colorCyan + "\nInsert 10, 20, 30, 40 & 50 in the list\n" + colorReset)

	// Insert
	singlyLinkedList.Insert(10, 1)
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

func executeCircularLinkedListOperations() {
	circularLinkedList := circularlist.CirculaLinkedList{}

	fmt.Println(colorCyan + "\nInsert 200, 300, 400, 500, 600 & 250 in the list\n" + colorReset)

	// Insert one after the other
	circularLinkedList.Insert(200, 1)
	circularLinkedList.Insert(300, 2)
	circularLinkedList.Insert(400, 3)
	circularLinkedList.Insert(500, 4)
	circularLinkedList.Insert(600, 5)
	// Insert in the middle
	circularLinkedList.Insert(250, 2)

	fmt.Println(colorCyan+"LastNode is: ", circularLinkedList.LastNode())
	fmt.Println(colorReset)

	fmt.Println(colorCyan + "\nPrint the list elements:" + colorReset)
	fmt.Printf(colorYellow+"\nListLength: %d\t\n", circularLinkedList.ListLength())
	fmt.Println(colorReset)

	fmt.Println(colorRed + "\nDelete 200 from the list" + colorReset)
	circularLinkedList.Delete(1)
	fmt.Println(colorCyan + "\nPrint the list elements after delete:" + colorReset)
	fmt.Printf(colorYellow+"\nListLength: %d\t\n", circularLinkedList.ListLength())
	fmt.Println(colorReset)

	fmt.Println(colorRed + "\nDelete 400 from the list" + colorReset)
	circularLinkedList.Delete(3)
	fmt.Println(colorCyan + "\nPrint the list elements after delete:" + colorReset)
	fmt.Printf(colorYellow+"\nListLength: %d\t\n", circularLinkedList.ListLength())
	fmt.Println(colorReset)

	fmt.Println(colorRed + "\nDelete 600 from the list" + colorReset)
	circularLinkedList.Delete(4)
	fmt.Println(colorCyan + "\nPrint the list elements after delete:" + colorReset)
	fmt.Printf(colorYellow+"\nListLength: %d\t\n", circularLinkedList.ListLength())
	fmt.Println(colorReset)

	fmt.Println(colorCyan+"LastNode is: ", circularLinkedList.LastNode())
	fmt.Println(colorReset)
}

func executeproblemAndsolutions() {
	singlyLinkedList := singlylist.SinglyLinkedList{}

	// Insert & Maintain a Hash Table
	singlyLinkedList.InsertAndMaintainHashTable(10, 1)
	singlyLinkedList.InsertAndMaintainHashTable(20, 2)
	singlyLinkedList.InsertAndMaintainHashTable(30, 3)
	singlyLinkedList.InsertAndMaintainHashTable(40, 4)
	singlyLinkedList.InsertAndMaintainHashTable(50, 5)

	startTime := time.Now()
	fmt.Printf(colorYellow+"\tFindNthNodeFromEnd where n = 2: %v\t\n", problemAndsolutions.FindNthNodeFromEnd(&singlyLinkedList, 2))
	fmt.Println(colorReset)
	endTime := time.Now()
	fmt.Println(colorGreen+"Time elapsed for FindNthNodeFromEnd where n = 2\t", endTime.Sub(startTime))
	fmt.Println(colorReset)

	startTime = time.Now()
	fmt.Printf(colorYellow+"\tFindNthNodeFromEndUsingHashTable where n = 3: %v\t\n", problemAndsolutions.FindNthNodeFromEndUsingHashTable(&singlyLinkedList, 3))
	fmt.Println(colorReset)
	endTime = time.Now()
	fmt.Println(colorGreen+"Time elapsed for FindNthNodeFromEndUsingHashTable where n = 3\t", endTime.Sub(startTime))
	fmt.Println(colorReset)

	startTime = time.Now()
	fmt.Printf(colorYellow+"\tFindNthNodeInOneScan where n = 1: %v\t\n", problemAndsolutions.FindNthNodeInOneScan(&singlyLinkedList, 1))
	fmt.Println(colorReset)
	endTime = time.Now()
	fmt.Println(colorGreen+"Time elapsed for tFindNthNodeInOneScan where n = 1\t", endTime.Sub(startTime))
	fmt.Println(colorReset)

	startTime = time.Now()
	fmt.Printf(colorYellow+"\tFindNthNodeInOneScan where n = 7: %v\t\n", problemAndsolutions.FindNthNodeInOneScan(&singlyLinkedList, 7))
	fmt.Println(colorReset)
	endTime = time.Now()
	fmt.Println(colorGreen+"Time elapsed for tFindNthNodeInOneScan where n = 7\t", endTime.Sub(startTime))
	fmt.Println(colorReset)
}

// The main function calls and executes the singly linked list operations.
func main() {
	fmt.Println("SINGLY LINKED LIST")
	executeSinglyLinkedListOperations()
	printDottedLine()
	fmt.Println("DOUBLY LINKED LIST")
	executeDoublyLinkedListOperations()
	printDottedLine()
	fmt.Println("CIRCULAR LINKED LIST")
	executeCircularLinkedListOperations()
	printDottedLine()
	fmt.Println("EXERCISE PROBLEMS")
	executeproblemAndsolutions()
	printDottedLine()
}
