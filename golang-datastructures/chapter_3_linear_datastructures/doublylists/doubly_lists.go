package doublylists

import "fmt"

// The type Node represents a node in a doubly linked list with a property of type int and pointers to
// the next and previous nodes.
// @property {int} property - The "property" field in the Node struct is an integer that represents the
// value or data stored in the node.
// @property nextNode - nextNode is a pointer to the next node in a linked list. It allows us to
// traverse the list by following the pointer to the next node. If the current node is the last node in
// the list, the nextNode pointer will be nil.
// @property previousNode - `previousNode` is a pointer to the previous node in a doubly linked list.
// It allows for efficient traversal in both forward and backward directions.
type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}

// The LinkedList type represents a linked list data structure with a head node.
// @property headNode - headNode is a pointer to the first node in a linked list. It is used to keep
// track of the beginning of the list and allows for easy traversal of the list.
type LinkedList struct {
	headNode *Node
}

// This function adds a new node to the beginning of a doubly linked list. It takes an integer value as
// input and creates a new node with that value. If the list is not empty, the new node is linked to
// the current head node and the previous node of the current head node is set to the new node.
// Finally, the head node of the list is updated to point to the new node.
func (linkedList *LinkedList) AddToHead(property int) {
	node := &Node{}
	node.property = property
	node.nextNode = nil
	node.previousNode = nil

	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
		linkedList.headNode.previousNode = node
	}
	linkedList.headNode = node
}

// `IterateList()` is a method of the `LinkedList` struct that iterates through the linked list and
// prints out the value of each node in the list. It starts at the head node of the list and continues
// until it reaches the end of the list (i.e., the `nextNode` pointer of the current node is `nil`).
// For each node, it prints out the value of the `property` field using `fmt.Printf()`. The `\t`
// character is used to add a tab between each value printed.
func (linkedList *LinkedList) IterateList() {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Printf("%d\t", node.property)
	}
}

// `LastNode()` is a method of the `LinkedList` struct that returns a pointer to the last node in the
// linked list. It does this by iterating through the list starting from the head node and checking if
// the `nextNode` pointer of the current node is `nil`. If it is, then the current node is the last
// node in the list, so it is assigned to the `lastNode` variable. Once the end of the list is reached,
// `lastNode` is returned.
func (linkedList *LinkedList) LastNode() *Node {
	lastNode := &Node{}
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// `AddToEnd()` is a method of the `LinkedList` struct that adds a new node to the end of the linked
// list. It takes an integer value as input and creates a new node with that value. If the list is not
// empty, it finds the last node in the list using the `LastNode()` method and links the new node to
// it. The `previousNode` pointer of the new node is set to the last node, and the `nextNode` pointer
// of the last node is set to the new node. If the list is empty, the new node becomes the head node of
// the list.
func (linkedList *LinkedList) AddToEnd(property int) {
	node := Node{}
	node.property = property
	node.nextNode = nil
	node.previousNode = nil

	lastNode := linkedList.LastNode()

	if lastNode != nil {
		lastNode.nextNode = &node
		node.previousNode = lastNode
	}
}

// `func (linkedList *LinkedList) NodeWithValue(property int) *Node` is a method of the `LinkedList`
// struct that searches the linked list for a node with a specific value. It takes an integer value as
// input and iterates through the list starting from the head node. For each node, it checks if the
// `property` field of the node matches the input value. If it does, the method returns a pointer to
// that node. If no node with the input value is found, the method returns a pointer to an empty `Node`
// struct.
func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	nodeWithValue := &Node{}
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWithValue = node
			break
		}
	}
	return nodeWithValue
}

// `func (linkedList *LinkedList) AddAfter(nodeProperty, property int)` is a method of the `LinkedList`
// struct that adds a new node with a given value after a node with a specific value. It takes two
// integer values as input: `nodeProperty` represents the value of the node after which the new node
// should be added, and `property` represents the value of the new node.
func (linkedList *LinkedList) AddAfter(nodeProperty, property int) {
	newNode := &Node{}
	newNode.nextNode = nil
	newNode.previousNode = nil
	newNode.property = property

	nodeWithValue := linkedList.NodeWithValue(nodeProperty)
	lastNode := linkedList.LastNode()

	// This implies we are trying to add at the end of the list
	if lastNode.property == nodeWithValue.property {
		lastNode.nextNode = newNode
		newNode.previousNode = lastNode
	}

	// This implies we are trying to add the node in the middle
	if nodeWithValue != nil {
		newNode.nextNode = nodeWithValue.nextNode
		nodeWithValue.nextNode = newNode
		newNode.previousNode = nodeWithValue
	}
}

// `func (linkedList *LinkedList) NodeBetweenValues(firstProperty, secondProperty int) *Node` is a
// method of the `LinkedList` struct that searches the linked list for a node that is between two nodes
// with specific values. It takes two integer values as input: `firstProperty` represents the value of
// the node before the desired node, and `secondProperty` represents the value of the node after the
// desired node.
func (linkedList *LinkedList) NodeBetweenValues(firstProperty, secondProperty int) *Node {
	nodeBetweenValues := &Node{}
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty && node.nextNode.property == secondProperty {
				nodeBetweenValues = node
				break
			}
		}
	}
	return nodeBetweenValues
}
