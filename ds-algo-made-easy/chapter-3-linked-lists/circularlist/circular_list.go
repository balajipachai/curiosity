package circularlist

import "fmt"

const (
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
)

// The CLListNode type represents a node in a circular linked list, with an integer data field and a
// reference to the next node.
// @property {int} data - The "data" property of the CLListNode struct is an integer that represents
// the value stored in the node.
// @property nextNode - nextNode is a pointer to the next node in a circular linked list. It points to
// the memory address of the next CLListNode object in the list.
type CLListNode struct {
	data     int
	nextNode *CLListNode
}

// The CircularLinkedList type represents a circular linked list.
// @property headNode - The headNode is a pointer to the first node in the circular linked list.
type CirculaLinkedList struct {
	headNode *CLListNode
}

// The `ListLength` function is used to calculate the length of the circular linked list and print its
// elements.
func (circulaLinkedList *CirculaLinkedList) ListLength() int {
	count := 0
	headNode := circulaLinkedList.headNode
	lastNode := circulaLinkedList.LastNode()
	for node := headNode; node != nil; node = node.nextNode {
		count++
		fmt.Printf(colorMagenta+"\t%d\t", node.data)
		if node == lastNode {
			break
		}
	}
	fmt.Printf("\n")
	fmt.Println(colorReset)
	return count
}

// The `LastNode` function is used to find and return the last node in the circular linked list.
func (circulaLinkedList *CirculaLinkedList) LastNode() *CLListNode {
	var lastNode *CLListNode
	for node := circulaLinkedList.headNode; node != nil; node = node.nextNode {
		if circulaLinkedList.headNode == node.nextNode {
			lastNode = node
			break
		}
	}
	return lastNode
}

// The `Insert` function is used to insert a new node into the circular linked list at a specified
// position.
func (circulaLinkedList *CirculaLinkedList) Insert(data, position int) {
	headNode := circulaLinkedList.headNode
	nodeToInsert := &CLListNode{}
	nodeToInsert.data = data

	if position < 1 {
		panic("Invalid position")
	}

	// Insert in the beginning
	if position == 1 {
		nodeToInsert.nextNode = headNode
		circulaLinkedList.headNode = nodeToInsert
		circulaLinkedList.headNode.nextNode = nodeToInsert
	} else {
		k := 0
		var positionNode *CLListNode
		lastNode := circulaLinkedList.LastNode()

		for node := headNode; node != nil; node = node.nextNode {
			k++
			if k < position {
				positionNode = node
			}
			if node == lastNode {
				break
			}
		}
		// Insert at the end
		if positionNode.nextNode == headNode {
			nodeToInsert.nextNode = headNode
			positionNode.nextNode = nodeToInsert
		} else {
			// Insert in the middle
			nodeToInsert.nextNode = positionNode.nextNode
			positionNode.nextNode = nodeToInsert
		}
	}
}

// The `Delete` function is used to delete a node from the circular linked list at a specified
// position.
func (circulaLinkedList *CirculaLinkedList) Delete(position int) {
	headNode := circulaLinkedList.headNode

	if position < 1 {
		panic("Invalid position")
	}

	// Delete the first node
	if position == 1 {
		lastNode := circulaLinkedList.LastNode()
		circulaLinkedList.headNode = headNode.nextNode
		lastNode.nextNode = circulaLinkedList.headNode
	} else {
		k := 0
		var positionNode *CLListNode
		lastNode := circulaLinkedList.LastNode()

		for node := headNode; node != nil; node = node.nextNode {
			k++
			if k < position {
				positionNode = node
			}
			if node == lastNode {
				break
			}
		}
		nodeToDelete := positionNode.nextNode
		// Delete the last node
		if positionNode.nextNode.nextNode == headNode {
			positionNode.nextNode = nodeToDelete.nextNode
			nodeToDelete.nextNode = nil
		} else {
			positionNode.nextNode = nodeToDelete.nextNode
			nodeToDelete.nextNode = nil
		}
	}
}
