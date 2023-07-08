package singlylist

import "fmt"

const (
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
)

var HashTable = make(map[int]*ListNode)

// The above code defines a struct type ListNode with an integer data field and a pointer to the next
// node.
// @property {int} data - The data property is an integer that represents the value stored in a node of
// a linked list. It can be any integer value depending on the requirements of the program.
// @property NextNode - NextNode is a pointer to the next node in a linked list. It is used to traverse
// the list and access the data stored in each node. If the current node is the last node in the list,
// NextNode will be set to nil to indicate the end of the list.
type ListNode struct {
	Data     int
	NextNode *ListNode
}

// The SinglyLinkedList type represents a linked list data structure with a head node.
// @property HeadNode - HeadNode is a pointer to the first node in a singly linked list. It points to
// the node that contains the first element of the list. If the list is empty, HeadNode will be nil.
type SinglyLinkedList struct {
	HeadNode *ListNode
}

// This function is calculating the length of a singly linked list. It starts by getting the head node
// of the list from the `SinglyLinkedList` struct. It then iterates through the list using a `for`
// loop, starting from the head node and moving to the next node until it reaches the end of the list
// (when `node` is `nil`). For each node, it increments a `count` variable and prints the data value of
// the node using `fmt.Printf()`. Finally, it returns the `count` variable, which represents the length
// of the list.
func (singlyLinkedList *SinglyLinkedList) ListLength() int {
	HeadNode := singlyLinkedList.HeadNode
	count := 0
	for node := HeadNode; node != nil; node = node.NextNode {
		count++
		fmt.Printf(colorMagenta+"\t%d\t", node.Data)
	}
	fmt.Printf("\n")
	fmt.Println(colorReset)
	return count
}

// The `Insert` function is a method of the `SinglyLinkedList` struct that inserts a new node with a
// given data value at a specified position in the linked list.
func (singlyLinkedList *SinglyLinkedList) Insert(data, position int) {
	HeadNode := singlyLinkedList.HeadNode

	newNode := &ListNode{}
	newNode.Data = data

	if position < 1 {
		panic("Invalid position")
	}

	// position == 1, insert node at the beginning
	if position == 1 {
		newNode.NextNode = HeadNode
		singlyLinkedList.HeadNode = newNode
	} else {
		count := 0
		var positionNode *ListNode
		for node := HeadNode; node != nil; node = node.NextNode {
			count++
			if count < position {
				positionNode = node
			}
		}
		//Intermediary node
		if positionNode.NextNode != nil {
			newNode.NextNode = positionNode.NextNode
			positionNode.NextNode = newNode
		} else {
			// Last node
			positionNode.NextNode = newNode
		}

	}
}

func (singlyLinkedList *SinglyLinkedList) InsertAndMaintainHashTable(data, position int) {
	HeadNode := singlyLinkedList.HeadNode

	newNode := &ListNode{}
	newNode.Data = data

	if position < 1 {
		panic("Invalid position")
	}

	// position == 1, insert node at the beginning
	if position == 1 {
		newNode.NextNode = HeadNode
		singlyLinkedList.HeadNode = newNode
		HashTable[position] = newNode
	} else {
		count := 0
		var positionNode *ListNode
		for node := HeadNode; node != nil; node = node.NextNode {
			count++
			if count < position {
				positionNode = node
			}
		}
		//Intermediary node
		if positionNode.NextNode != nil {
			newNode.NextNode = positionNode.NextNode
			positionNode.NextNode = newNode
		} else {
			// Last node
			positionNode.NextNode = newNode
		}
		HashTable[position] = newNode

	}
}

// The `Delete` function is a method of the `SinglyLinkedList` struct that deletes a node at a
// specified position in the linked list. It takes an integer `position` as input, which represents the
// position of the node to be deleted. The function starts by getting the head node of the list from
// the `SinglyLinkedList` struct. It then checks if the `position` is less than 1, and if so, it panics
// with an error message. If `position` is 1, it means the first node needs to be deleted, so it
// updates the `HeadNode` pointer to point to the second node in the list. If `position` is greater
// than 1, it iterates through the list using a `for` loop, starting from the head node and moving to
// the next node until it reaches the node at the specified position (when `count` equals
// `position-1`). It then updates the `NextNode` pointer of the previous node to point to the node
// after the one being deleted, effectively removing the node from the list. If the node being deleted
// is the last node in the list, the `NextNode` pointer is set to nil
func (singlyLinkedList *SinglyLinkedList) Delete(position int) {
	HeadNode := singlyLinkedList.HeadNode

	listLength := singlyLinkedList.ListLength()

	if position < 1 || position > listLength {
		panic("Invalid position")
	}

	// position == 1, insert node at the beginning
	if position == 1 {
		singlyLinkedList.HeadNode = HeadNode.NextNode
	} else {
		count := 0
		var positionNode *ListNode
		for node := HeadNode; node != nil; node = node.NextNode {
			count++
			if count < position {
				positionNode = node
			}
		}
		// Intermediary node
		if positionNode.NextNode.NextNode != nil {
			positionNode.NextNode = positionNode.NextNode.NextNode
		} else {
			// Last node
			positionNode.NextNode = nil
		}
	}
}

func CreateListWithALoop() SinglyLinkedList {
	singlyLinkedList := SinglyLinkedList{}
	singlyLinkedList.HeadNode = nil
	firstNode := ListNode{
		Data:     10,
		NextNode: nil,
	}
	secondNode := &ListNode{
		Data:     20,
		NextNode: nil,
	}
	thirdNode := &ListNode{
		Data:     30,
		NextNode: nil,
	}
	fourthNode := &ListNode{
		Data:     40,
		NextNode: nil,
	}
	fifthNode := &ListNode{
		Data:     50,
		NextNode: nil,
	}

	// Linking
	singlyLinkedList.HeadNode = &firstNode
	firstNode.NextNode = secondNode
	secondNode.NextNode = thirdNode
	thirdNode.NextNode = fourthNode
	fourthNode.NextNode = fifthNode
	fifthNode.NextNode = thirdNode

	return singlyLinkedList

}
