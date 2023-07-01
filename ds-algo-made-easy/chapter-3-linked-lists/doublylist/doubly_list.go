package doublylist

import "fmt"

const (
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
)

// The DLListNode type represents a node in a doubly linked list.
// @property {int} data - The "data" property of the DLListNode struct represents the value or data
// stored in the node. It can be of type int in this case, but it can be of any other data type
// depending on the requirements of your program.
// @property nextNode - nextNode is a pointer to the next node in a doubly linked list. It points to
// the node that comes after the current node.
// @property previousNode - The previousNode property is a pointer to the previous node in a doubly
// linked list. It allows for traversal of the list in both directions, forward and backward.
type DLListNode struct {
	data         int
	nextNode     *DLListNode
	previousNode *DLListNode
}

// The DoublyLinkedList type represents a doubly linked list.
// @property headNode - The headNode is a pointer to the first node in the doubly linked list.
type DoublyLinkedList struct {
	headNode *DLListNode
}

// The `LastNode` function is used to find and return the last node in a doubly linked list.
func (doublyLinkedList *DoublyLinkedList) LastNode() *DLListNode {
	lastNode := &DLListNode{}

	for node := doublyLinkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil && node.previousNode != nil {
			lastNode = node
			break
		}
	}
	return lastNode
}

// The `PrintReverseList` function is used to print the elements of a doubly linked list in reverse
// order.
func (doublyLinkedList *DoublyLinkedList) PrintReverseList() {
	lastNode := doublyLinkedList.LastNode()
	for lNode := lastNode; lNode != nil; lNode = lNode.previousNode {
		fmt.Printf(colorMagenta+"\t%d\t", lNode.data)
	}
	fmt.Printf("\n")
	fmt.Println(colorReset)
}

// The `ListLength` function is used to calculate the length of a doubly linked list. It iterates
// through each node in the list, starting from the head node, and increments a counter variable
// `count` for each node encountered. It also prints the data of each node as it traverses the list.
// Finally, it returns the value of `count`, which represents the length of the list.
func (doublyLinkedList *DoublyLinkedList) ListLength() int {
	count := 0
	headNode := doublyLinkedList.headNode

	for node := headNode; node != nil; node = node.nextNode {
		count++
		fmt.Printf(colorMagenta+"\t%d\t", node.data)
	}
	fmt.Printf("\n")
	fmt.Println(colorReset)
	return count
}

// The `Insert` function is used to insert a new node with a given data value at a specified position
// in a doubly linked list.
func (doublyLinkedList *DoublyLinkedList) Insert(data, position int) {
	headNode := doublyLinkedList.headNode
	newNode := &DLListNode{}
	newNode.data = data

	if position < 1 {
		panic("Invalid position")
	}

	// Insert at the beginning
	if position == 1 {
		newNode.nextNode = headNode
		newNode.previousNode = headNode
		doublyLinkedList.headNode = newNode
	} else {
		k := 0
		var positionNode *DLListNode

		for node := headNode; node != nil; node = node.nextNode {
			k++
			if k < position {
				positionNode = node
			}
		}
		nodeToInsert := newNode
		// Insert in the middle
		if positionNode.nextNode != nil {
			nodeToInsert.nextNode = positionNode.nextNode
			positionNode.nextNode.previousNode = nodeToInsert
			nodeToInsert.previousNode = positionNode
			positionNode.nextNode = nodeToInsert
		} else {
			// Insert at the end
			positionNode.nextNode = nodeToInsert
			nodeToInsert.previousNode = positionNode
			nodeToInsert.nextNode = nil
		}
	}
}

// The `Delete` function is used to delete a node at a specified position in a doubly linked list.
func (doublyLinkedList *DoublyLinkedList) Delete(position int) {
	headNode := doublyLinkedList.headNode
	var positionNode *DLListNode

	if position < 1 {
		panic("Invalid position")
	}

	// Delete the first node
	if position == 1 {
		doublyLinkedList.headNode = headNode.nextNode
		headNode.nextNode.previousNode = nil
	} else {
		k := 0
		for node := headNode; node != nil; node = node.nextNode {
			k++
			if k < position {
				positionNode = node
			}
		}

		nodeToDelete := positionNode.nextNode
		// Delete from the middle
		if positionNode.nextNode.nextNode != nil {
			positionNode.nextNode = nodeToDelete.nextNode
			nodeToDelete.nextNode.previousNode = positionNode
			nodeToDelete = nil
		} else {
			// Delete the last node
			positionNode.nextNode = nil
			nodeToDelete = nil
		}
	}
}
