package main

import (
	"fmt"

	"example.com/doublylists"
	"example.com/lists"
	"example.com/queue"
	"example.com/sets"
	"example.com/stack"
	"example.com/tuples"
)

// The function demonstrates the implementation and usage of singly and doubly linked lists in Go.
func executeListOperations() {
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
}

// The function demonstrates the use of sets in Go by performing various set operations such as adding,
// deleting, checking for element existence, finding intersection and union of sets.
func executeSetOperations() {
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
}

// The function executes tuple operations and prints the results.
func executeTupleOperations() {
	fmt.Println("\n*******************Tuples*******************")
	fmt.Println("\tSquare of 4: ", tuples.Square(4))
	fmt.Println("\tCube of 4: ", tuples.Cube(4))
	xToThePower4 := tuples.XToThePower4(4, tuples.Cube(4))
	fmt.Println("\t4^4: ", xToThePower4)
}

// The function creates orders and adds them to a queue, then deletes one order from the queue and
// prints the updated queue.
func executeQueueOperations() {
	fmt.Println("\n*******************Queue*******************")
	priority1 := 10
	quantity1 := 100
	product1 := "Utensils"
	customerName1 := "Balaji"

	priority2 := 20
	quantity2 := 200
	product2 := "Chocolates"
	customerName2 := "Anjali"

	priority3 := 15
	quantity3 := 150
	product3 := "Moong Dal Halwa"
	customerName3 := "Ajit"

	order1 := &queue.Order{}
	order1.New(priority1, quantity1, product1, customerName1)

	order2 := &queue.Order{}
	order2.New(priority2, quantity2, product2, customerName2)

	order3 := &queue.Order{}
	order3.New(priority3, quantity3, product3, customerName3)

	fmt.Println("\n*******************Printing Order1, Order2 & Order 3*******************")
	fmt.Println("\t", order1)
	fmt.Println("\t", order2)
	fmt.Println("\t", order3)

	fmt.Println("\nAdd Order1, Order2, Order3 to the Queue")
	queue := make(queue.Queue, 0)

	queue.Add(order1)
	queue.Add(order2)
	queue.Add(order3)
	fmt.Println("\n*******************Printing Queue*******************")
	for _, order := range queue {
		fmt.Println("\t", order)
	}

	fmt.Println("\n*******************Delete Order3*******************")
	queue.Delete(order3)

	fmt.Println("\n*******************Printing Queue After Deletion*******************")
	for _, order := range queue {
		fmt.Println("\t", order)
	}
}

// The function executes stack operations such as pushing and popping elements from a stack and prints
// the results.
func executeStackOperations() {
	fmt.Println("\n*******************Stack*******************")
	s := &stack.Stack{}
	s.New()
	fmt.Println("\nPrinting the new stack:\n\t", s)
	fmt.Println("\n Adds 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 in the stack")
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Push(7)
	s.Push(8)
	s.Push(9)
	s.Push(10)
	fmt.Println("\nPrinting the Stack elements:\t", s.Element)
	fmt.Println()
	fmt.Println("\nPrints the Top of the stack element:\t", s.Top())
	fmt.Println("\nRemoves last element from the stack:")
	s.Pop()
	fmt.Println("\nPrinting the Stack elements after Pop:\t", s.Element)
	fmt.Println("\nPrints the Top of the stack element:\t", s.Top())

}

// The main function executes various operations on lists, sets, tuples, and queues.
func main() {
	executeListOperations()
	executeSetOperations()
	executeTupleOperations()
	executeQueueOperations()
	executeStackOperations()
}
