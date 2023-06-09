package main

import (
	"fmt"

	"example.com/doublylists"
	"example.com/lists"
	"example.com/sets"
	"example.com/tuples"
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

	fmt.Print("\n")
	fmt.Println("*******************Sets*******************")
	setA := sets.Set{}
	setB := sets.Set{}
	setC := sets.Set{}

	setA.New()
	setB.New()
	setC.New()

	fmt.Println("Adds element in Set A, B & C")
	setA.AddElement(1)
	setA.AddElement(2)
	setA.AddElement(3)
	setA.AddElement(4)

	fmt.Println("\nTry to add existing element")
	setA.AddElement(4)

	setA.AddElement(5)
	fmt.Println("Set A: ", setA)

	setB.AddElement(6)
	setB.AddElement(7)
	setB.AddElement(8)
	setB.AddElement(9)
	setB.AddElement(10)
	fmt.Println("Set B: ", setB)

	setC.AddElement(11)
	setC.AddElement(12)
	setC.AddElement(13)
	setC.AddElement(14)
	setC.AddElement(15)
	fmt.Println("Set C: ", setC)

	fmt.Println("\nChecks if Set contains element")
	fmt.Println("\tsetA.ContainsElement(4)", setA.ContainsElement(4))
	fmt.Println("\tsetA.ContainsElement(10)", setA.ContainsElement(10))

	fmt.Println("\nDeletes 1, 7, 14 from SetA, SetB & SetC")
	setA.DeleteElement(1)
	setB.DeleteElement(7)
	setC.DeleteElement(14)

	fmt.Println("\nTry to delete non-existing item")
	setC.DeleteElement(5)

	fmt.Println("\nPrint Sets after deletion")
	fmt.Println("\tSet A: ", setA)
	fmt.Println("\tSet B: ", setB)
	fmt.Println("\tSet C: ", setC)

	fmt.Println("\nAdd 77 to SetA, SetB & 4 to SetC")
	setA.AddElement(77)
	setB.AddElement(77)
	setC.AddElement(4)

	fmt.Println("\nPrint Sets after adding 77 & 4 to the sets")
	fmt.Println("\tSet A: ", setA)
	fmt.Println("\tSet B: ", setB)
	fmt.Println("\tSet C: ", setC)

	fmt.Println("\nFinds the intersection of SetA & SetB AND SetA & SetC")
	fmt.Println("\tSetA Intersection SetB: ", setA.Intersection(&setB))
	fmt.Println("\tSetA Intersection SetC: ", setA.Intersection(&setC))

	fmt.Println("\nFinds the union of SetA & SetB AND SetA & SetC")
	fmt.Println("\tSetA Union SetB: ", setA.Union(&setB))
	fmt.Println("\tSetA Union SetC: ", setA.Union(&setC))

	fmt.Println("\n*******************Tuples*******************")
	fmt.Println("\tSquare of 4: ", tuples.Square(4))
	fmt.Println("\tCube of 4: ", tuples.Cube(4))
	xToThePower4 := tuples.XToThePower4(4, tuples.Cube(4))
	fmt.Println("\t4^4: ", xToThePower4)

}
