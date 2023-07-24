package trees

import (
	"fmt"
	"strings"
)

const (
	QUEUE_SIZE   = 10
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
)

// The BinaryTreeNode type represents a node in a binary tree with an integer data value and references
// to its left and right child nodes.
// @property {int} data - The "data" property of the BinaryTreeNode struct represents the value stored
// in the node. It can be of any data type, but in this case, it is specified as an integer (int).
// @property left - The "left" property of a BinaryTreeNode struct is a pointer to another
// BinaryTreeNode struct, representing the left child of the current node.
// @property right - The "right" property of a BinaryTreeNode struct is a pointer to another
// BinaryTreeNode struct, representing the right child of the current node.
type BinaryTreeNode struct {
	data  int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

// The BinaryTree type represents a binary tree data structure.
// @property root - The root property is a pointer to the root node of the binary tree.
type BinaryTree struct {
	root *BinaryTreeNode
}

// QUEUES FOR TREES
// The Queue type represents a queue data structure for storing BinaryTreeNode elements.
// @property {int} front - The front property represents the index of the front element in the queue.
// It is the position of the element that will be dequeued next.
// @property {int} rear - The "rear" property in the Queue struct represents the index of the last
// element in the queue.
// @property {[]*BinaryTreeNode} elements - The `elements` property is a slice of `BinaryTreeNode`
// pointers. It is used to store the elements of the queue.
type Queue struct {
	front    int
	rear     int
	elements []*BinaryTreeNode
}

// The `CreateNew()` function is a method of the `Queue` struct. It is used to initialize a new queue
// by setting the `front` and `rear` variables to -1, indicating an empty queue. It also creates a new
// slice of `BinaryTreeNode` pointers with a length of 0, which will be used to store the elements of
// the queue.
func (queue *Queue) CreateNew() {
	queue.front = -1
	queue.rear = -1
	queue.elements = make([]*BinaryTreeNode, 0)
}

// The `IsEmpty()` function is a method of the `Queue` struct. It checks if the queue is empty by
// comparing the `front` variable of the queue to -1. If the `front` is -1, it means that the queue is
// empty and the function returns `true`. Otherwise, it returns `false`.
func (queue *Queue) IsEmpty() bool {
	return queue.front == -1
}

// The `EnQueue()` function is used to add a `BinaryTreeNode` to the queue.
func (queue *Queue) EnQueue(binaryTreeNode *BinaryTreeNode) {
	queue.elements = append(queue.elements, binaryTreeNode)
	queue.rear = len(queue.elements) - 1

	if queue.front == -1 {
		queue.front = queue.rear
	}
}

// The `DeQueue()` function is used to remove and return the front element from the queue.
func (queue *Queue) DeQueue() *BinaryTreeNode {
	if queue.IsEmpty() {
		fmt.Println(colorRed + "\tUndeflow: Queue is empty" + colorReset)
		return nil
	}
	data := queue.elements[queue.front]
	queue.elements = queue.elements[queue.front+1:]
	if queue.front == queue.rear {
		queue.front = -1
		queue.rear = -1
	} else {
		queue.front = 0
		queue.rear -= 1
	}
	return data
}

// The `FrontElement()` function is a method of the `Queue` struct. It is used to retrieve the front
// element of the queue without removing it.
func (queue *Queue) FrontElement() *BinaryTreeNode {
	return queue.elements[queue.front]
}

// The `EmptyQueue()` function is a method of the `Queue` struct. It is used to remove all elements
// from the queue until it becomes empty.
func (queue *Queue) EmptyQueue() {
	for !queue.IsEmpty() {
		queue.DeQueue()
	}
}

// The `LevelOrderTraversal()` function is a method of the `BinaryTree` struct. It performs a
// level-order traversal of the binary tree, which means it visits each node in the tree level by
// level, from left to right.
func (binaryTree *BinaryTree) LevelOrderTraversal() {

	// Create Queue
	// The code `queue := &Queue{}` creates a new instance of the `Queue` struct and assigns it to the
	// variable `queue`. The `&` symbol is used to take the address of the newly created struct, so that
	// `queue` is a pointer to the `Queue` struct.
	queue := &Queue{}
	queue.CreateNew()

	node := &BinaryTreeNode{}
	root := binaryTree.root

	if root == nil {
		return
	}

	queue.EnQueue(root)

	for !queue.IsEmpty() {
		node = queue.DeQueue()
		fmt.Printf(colorMagenta+"\t%d\t", node.data)
		if node.left != nil {
			queue.EnQueue(node.left)
		}

		if node.right != nil {
			queue.EnQueue(node.right)
		}
	}
	fmt.Printf("\n")
}

// The `Insert` function is a method of the `BinaryTree` struct. It is used to insert a new node with
// the given data into the binary tree.
func (binaryTree *BinaryTree) Insert(data int) {
	newNode := &BinaryTreeNode{}
	newNode.data = data

	temp := &BinaryTreeNode{}
	root := binaryTree.root

	if root == nil {
		binaryTree.root = newNode
		return
	}

	queue := &Queue{}
	queue.CreateNew()

	queue.EnQueue(root)

	for !queue.IsEmpty() {
		temp = queue.DeQueue()
		if temp.left != nil {
			queue.EnQueue(temp.left)
		} else {
			temp.left = newNode
			return
		}

		if temp.right != nil {
			queue.EnQueue(temp.right)
		} else {
			temp.right = newNode
			return
		}
	}
}

// The preOrderHelper function performs a pre-order traversal of a binary tree and prints the data of
// each node.
func preOrderHelper(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	fmt.Printf(colorMagenta+"\t%d\t", node.data)
	preOrderHelper(node.left)
	preOrderHelper(node.right)
}

// The `PreOrder()` method of the `BinaryTree` struct is used to perform a pre-order traversal of the
// binary tree and print the data of each node.
func (binaryTree *BinaryTree) PreOrder() {
	preOrderHelper(binaryTree.root)
	fmt.Printf("\n")
}

// The function recursively traverses a binary tree in in-order fashion and prints the data of each
// node.
func inOrderHelper(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	inOrderHelper(node.left)
	fmt.Printf(colorMagenta+"\t%d\t", node.data)
	inOrderHelper((node.right))
}

// The `InOrder()` method of the `BinaryTree` struct is used to perform an in-order traversal of the
// binary tree and print the data of each node. It calls the `inOrderHelper()` function, passing the
// root node of the binary tree as an argument. The `inOrderHelper()` function recursively traverses
// the binary tree in in-order fashion, visiting the left subtree, then the current node, and finally
// the right subtree. It prints the data of each node in a formatted manner.
func (binaryTree *BinaryTree) InOrder() {
	inOrderHelper(binaryTree.root)
	fmt.Printf("\n")
}

// The function performs a post-order traversal of a binary tree and prints the data of each node.
func postOrderHelper(node *BinaryTreeNode) {
	if node == nil {
		return
	}
	postOrderHelper(node.left)
	postOrderHelper(node.right)
	fmt.Printf(colorMagenta+"\t%d\t", node.data)
}

// The `PostOrder()` method of the `BinaryTree` struct is used to perform a post-order traversal of the
// binary tree and print the data of each node. It calls the `postOrderHelper()` function, passing the
// root node of the binary tree as an argument. The `postOrderHelper()` function recursively traverses
// the binary tree in post-order fashion, visiting the left subtree, then the right subtree, and
// finally the current node. It prints the data of each node in a formatted manner.
func (binaryTree *BinaryTree) PostOrder() {
	postOrderHelper(binaryTree.root)
	fmt.Printf("\n")
}

// The function prints a binary tree with indentation and prefixes to represent the tree structure.
func printTreeWithIndent(node *BinaryTreeNode, level int, prefix string) {
	if node == nil {
		return
	}

	indent := strings.Repeat(" ", level*2)
	fmt.Printf("\t%s%s%d\n", indent, prefix, node.data)

	if node.left != nil || node.right != nil {
		printTreeWithIndent(node.left, level+1, "├── ")
		printTreeWithIndent(node.right, level+1, "└── ")
	}
}

// The `PrintTree()` method of the `BinaryTree` struct is used to print the binary tree in a tree-like
// structure. It calls the `printTreeWithIndent()` function, passing the root node of the binary tree,
// the initial level (0), and an empty prefix as arguments.
func (binaryTree *BinaryTree) PrintTree() {
	printTreeWithIndent(binaryTree.root, 0, "")

}
