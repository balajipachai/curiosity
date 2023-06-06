package lists

import "fmt"

// The below code defines a Node struct with an integer property and a pointer to the next Node.
// @property {int} property - The "property" field is an integer variable that represents the value or
// data stored in a node of a linked list. It can be any integer value depending on the requirements of
// the program.
// @property nextNode - nextNode is a pointer to the next node in a linked list data structure. It
// allows us to traverse the list by pointing to the next node in the sequence.
type Node struct {
	property int
	nextNode *Node
}

// The LinkedList type represents a linked list data structure with a head node.
// @property headNode - headNode is a pointer to the first node in a linked list. It is used to keep
// track of the beginning of the list and allows for traversal of the list by following the next
// pointers of each node.
type LinkedList struct {
	headNode *Node
}

// This function is adding a new node to the beginning of the linked list. It creates a new node with
// the given property value, sets its nextNode pointer to the current head node of the linked list, and
// then sets the headNode pointer of the linked list to point to the new node, effectively making it
// the new head node.
func (linkedList *LinkedList) AddToHead(property int) {
	node := Node{}
	node.property = property
	node.nextNode = linkedList.headNode
	linkedList.headNode = &node
}

// The `IterateList()` function is a method of the `LinkedList` struct that prints out the values of
// each node in the linked list. It starts at the head node of the linked list and iterates through
// each node by following the `nextNode` pointers until it reaches the end of the list (i.e., a node
// with a `nil` `nextNode` pointer). For each node, it prints out the value of its `property` field
// using `fmt.Printf()`.
func (linkedList *LinkedList) IterateList() {
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Printf("%d\t", node.property)
	}
}

// The `LastNode()` function is a method of the `LinkedList` struct that returns a pointer to the last
// node in the linked list. It does this by iterating through the linked list starting from the head
// node and following the `nextNode` pointers until it reaches the end of the list (i.e., a node with a
// `nil` `nextNode` pointer). Once it reaches the last node, it returns a pointer to that node. If the
// linked list is empty, it returns a pointer to a new empty `Node` struct.
func (linkedList *LinkedList) LastNode() *Node {
	lastNode := &Node{}
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		// node.nextNode == nil, implies that it is the lastNode
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd(property int) is a method of the `LinkedList` struct that
// adds a new node with the given `property` value to the end of the linked list. It creates a new
// `Node` struct with the given `property` value and a `nil` `nextNode` pointer. It then calls the
// `LastNode()` method to get a pointer to the last node in the linked list. If the linked list is not
// empty (i.e., `lastNode` is not `nil`), it sets the `nextNode` pointer of the last node to point to
// the new node, effectively adding it to the end of the list. If the linked list is empty, the new
// node becomes the head node of the linked list.
func (linkedList *LinkedList) AddToEnd(property int) {
	newNode := Node{}
	newNode.property = property
	newNode.nextNode = nil
	lastNode := linkedList.LastNode()

	if lastNode != nil {
		lastNode.nextNode = &newNode
	}
}

// NodeWithValue(property int) *Node is a method of the `LinkedList` struct that searches the linked list
// for a node with a `property` value that matches the given `property` argument.
// It does this by iterating through the linked list starting from the head node
// and following the `nextNode` pointers until it finds a node with a `property` value that matches the
// given `property` argument. Once it finds a matching node, it returns a pointer to that node. If no
// node with a matching `property` value is found, it returns a pointer to a new empty `Node` struct.
// This method can be used to find a specific node in the linked list so that a new node can be added
// after it using the `AddAfter()` method.
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

// AddAfter(nodeProperty, property int) is a method of the `LinkedList`
// struct that adds a new node with the given `property` value after the node in the linked list
// with a `property` value that matches the given `nodeProperty` argument.
func (linkedList *LinkedList) AddAfter(nodeProperty, property int) {
	newNode := Node{}
	newNode.property = property
	newNode.nextNode = nil
	nodeWithValue := linkedList.NodeWithValue(nodeProperty)

	if nodeWithValue.nextNode != nil {
		newNode.nextNode = nodeWithValue.nextNode
	}
	nodeWithValue.nextNode = &newNode
}
