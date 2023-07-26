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

// -----------------------------------------------------------------------------------

// The Stack type is a data structure that stores BinaryTreeNode elements and keeps track of the top
// element.
// @property {int} top - The "top" property represents the index of the top element in the stack. It
// keeps track of the position of the last element that was added to the stack.
// @property {[]*BinaryTreeNode} elements - The `elements` property is a slice of pointers to
// `BinaryTreeNode` objects.
type Stack struct {
	top      int
	elements []*BinaryTreeNode
}

// The above code is defining a method called "CreateNew" for a struct type called "Stack" in the Go
// programming language. This method initializes a new stack by setting the top index to -1 and
// creating an empty dynamic slice to store elements of type "*BinaryTreeNode".
func (stack *Stack) CreateNew() {
	stack.top = -1
	stack.elements = make([]*BinaryTreeNode, 0) // Assigns memory to dynamic slice
}

// The above code is defining a method called `IsEmpty()` for a struct type `Stack`. This method checks
// if the `top` field of the `Stack` struct is equal to -1, indicating that the stack is empty. It
// returns a boolean value indicating whether the stack is empty or not.
func (stack *Stack) IsEmpty() bool {
	return stack.top == -1
}

// The above code is defining a method called "Push" for a stack data structure in the Go programming
// language. This method takes a pointer to a BinaryTreeNode as a parameter and adds it to the stack.
// It does this by incrementing the "top" variable of the stack and appending the node to the
// "elements" slice of the stack.
func (stack *Stack) Push(node *BinaryTreeNode) {
	stack.top += 1
	stack.elements = append(stack.elements, node)
}

// The above code is implementing the Pop() function for a stack data structure in the Go programming
// language.
func (stack *Stack) Pop() *BinaryTreeNode {
	if stack.IsEmpty() {
		fmt.Println(colorRed + "\tUndeflow: Stack is empty" + colorReset)
		return nil
	}
	poppedElement := stack.elements[stack.top]
	stack.elements = stack.elements[:stack.top]
	stack.top -= 1
	return poppedElement
}

// The above code is defining a method called "TopElement" for a struct type called "Stack" in the Go
// programming language. This method returns the top element of the stack, which is stored in the
// "elements" slice at the index specified by the "top" variable.
func (stack *Stack) TopElement() *BinaryTreeNode {
	return stack.elements[stack.top]
}

// -----------------------------------------------------------------------------------

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

// The function recursively finds the maximum element in a binary tree.
func maxElementHelper(treeNode *BinaryTreeNode) int {
	var max int
	if treeNode != nil {
		rootVal := treeNode.data
		left := maxElementHelper(treeNode.left)
		right := maxElementHelper(treeNode.right)

		if left > right {
			max = left
		} else {
			max = right
		}

		if rootVal > max {
			max = rootVal
		}

		return max
	}
	return max
}

// The `FindMaxElementUsingRecursion()` method of the `BinaryTree` struct is used to find the maximum element in the
// binary tree.
func (binaryTree *BinaryTree) FindMaxElementUsingRecursion() int {
	if binaryTree.root != nil {
		return maxElementHelper(binaryTree.root)
	}
	return 0
}

// The `FindMaxElementUsingLevelOrder()` method of the `BinaryTree` struct is used to find the maximum
// element in the binary tree using a level-order traversal.
func (binaryTree *BinaryTree) FindMaxElementUsingLevelOrder() int {

	max := 0

	queue := &Queue{}
	queue.CreateNew()

	if binaryTree.root != nil {
		queue.EnQueue(binaryTree.root)

		for !queue.IsEmpty() {
			node := queue.DeQueue()

			if node.data > max {
				max = node.data
			}

			if node.left != nil {
				queue.EnQueue(node.left)
			}

			if node.right != nil {
				queue.EnQueue(node.right)
			}
		}
		return max
	}
	return max
}

// The searchElementHelper function recursively searches for an element in a binary tree.
func searchElementHelper(treeNode *BinaryTreeNode, element int) bool {
	if treeNode != nil {
		data := treeNode.data
		if data == element {
			return true
		} else {
			isFound := searchElementHelper(treeNode.left, element)
			if isFound {
				return true
			} else {
				return searchElementHelper(treeNode.right, element)
			}
		}
	}
	return false
}

// The `SearchElement` method of the `BinaryTree` struct is used to search for a specific element in
// the binary tree.
func (binaryTree *BinaryTree) SearchElement(element int) bool {
	if binaryTree.root != nil {
		return searchElementHelper(binaryTree.root, element)
	}
	return false
}

// The `SearchElementUsingLevelOrder` method of the `BinaryTree` struct is used to search for a
// specific element in the binary tree using a level-order traversal.
func (binaryTree *BinaryTree) SearchElementUsingLevelOrder(element int) bool {
	queue := &Queue{}
	queue.CreateNew()

	if binaryTree.root != nil {
		queue.EnQueue(binaryTree.root)

		for !queue.IsEmpty() {
			node := queue.DeQueue()
			if node.data == element {
				return true
			}
			if node.left != nil {
				queue.EnQueue(node.left)
			}
			if node.right != nil {
				queue.EnQueue(node.right)
			}
		}
	}
	return false
}

// The function calculates the size of a binary tree by recursively counting the number of nodes in the
// left and right subtrees.
func sizeOfBTHelper(node *BinaryTreeNode) int {
	if node != nil {
		return sizeOfBTHelper(node.left) + 1 + sizeOfBTHelper(node.right)
	}
	return 0
}

// The `SizeOfBinaryTree()` method of the `BinaryTree` struct is used to calculate the size of the
// binary tree. It calls the `sizeOfBTHelper()` function, passing the root node of the binary tree as
// an argument.
func (binaryTree *BinaryTree) SizeOfBinaryTree() int {
	if binaryTree.root != nil {
		return sizeOfBTHelper(binaryTree.root)
	}
	return 0
}

// The above code is implementing a method called `SizeOfBinaryTreeUsingLevelOrder()` for a binary tree
// data structure. This method calculates the size of the binary tree using a level order traversal
// approach.
func (binaryTree *BinaryTree) SizeOfBinaryTreeUsingLevelOrder() int {
	if binaryTree.root != nil {
		size := 0
		queue := &Queue{}
		queue.CreateNew()

		queue.EnQueue(binaryTree.root)

		for !queue.IsEmpty() {
			node := queue.DeQueue()
			size++
			if node.left != nil {
				queue.EnQueue(node.left)
			}
			if node.right != nil {
				queue.EnQueue(node.right)
			}
		}
		return size
	}
	return 0
}

// The above code is implementing a method called `LevelOrderInReverse()` for a binary tree data
// structure. This method performs a level order traversal of the binary tree and prints the nodes in
// reverse order.
func (binaryTree *BinaryTree) LevelOrderInReverse() {
	if binaryTree.root != nil {
		stack := &Stack{}
		stack.CreateNew()

		queue := &Queue{}
		queue.CreateNew()

		queue.EnQueue(binaryTree.root)

		for !queue.IsEmpty() {
			node := queue.DeQueue()
			stack.Push(node)

			if node.left != nil {
				queue.EnQueue(node.left)
			}

			if node.right != nil {
				queue.EnQueue(node.right)
			}
		}

		// Start Printing the LevelOrder in Reverse
		for !stack.IsEmpty() {
			poppedNode := stack.Pop()
			fmt.Printf(colorMagenta+"\t%d\t", poppedNode.data)
		}
		fmt.Println(colorReset)
	}
}

// The deleteTreeHelper function recursively deletes all nodes in a binary tree.
func deleteTreeHelper(treeNode *BinaryTreeNode) {
	if treeNode != nil {
		deleteTreeHelper(treeNode.left)
		deleteTreeHelper(treeNode.right)
		fmt.Println("\t"+colorRed+"Deleting node...", treeNode, colorReset)
		treeNode = nil
	}
}

// The above code is defining a method called `DeleteTree` for a `BinaryTree` struct in the Go
// programming language. This method is used to delete the entire binary tree. It first checks if the
// root of the binary tree is not nil, and if it is not, it calls a helper function called
// `deleteTreeHelper` to delete the tree.
func (binaryTree *BinaryTree) DeleteTree() {
	if binaryTree.root != nil {
		deleteTreeHelper(binaryTree.root)
	}
}

// The max function returns the larger of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// The function calculates the height of a binary tree by recursively traversing the tree and counting
// the number of levels.
func heightOfBTHelper(treeNode *BinaryTreeNode) int {
	var leftHeight, rightHeight int
	if treeNode == nil {
		return 0
	}
	leftHeight = heightOfBTHelper(treeNode.left)
	rightHeight = heightOfBTHelper(treeNode.right)
	return 1 + max(leftHeight, rightHeight)
}

// The above code is implementing a method called "HeightOfBinaryTree" for a binary tree data
// structure. This method calculates the height of the binary tree by recursively calculating the
// height of the left and right subtrees of each node and assigning the height to the node as the
// maximum of the heights of its two children plus 1. If the binary tree is empty (root is nil), the
// method returns 0.
func (binaryTree *BinaryTree) HeithtOfBinaryTree() int {
	/*
		Recursively calculate height of left and right subtrees of a node and assign height to the node as max of the heights of two children plus 1. This is similar to PreOrder tree traversal (and DFS of Graph algorithms).
	*/
	if binaryTree.root != nil {
		return heightOfBTHelper(binaryTree.root)
	}
	return 0
}

// The function calculates the height of a binary tree node by counting the number of levels from the
// given node to the bottom of the tree.
func height(node *BinaryTreeNode, isLeft bool) int {
	count := 0
	for node != nil {
		count++
		if isLeft {
			node = node.left
		} else {
			node = node.right
		}
	}
	return count
}

// The above code is implementing a non-recursive method to calculate the height of a binary tree. It
// first checks if the root of the binary tree is not nil. If it is not nil, it calculates the height
// of the left subtree and the height of the right subtree using the `height` function. The `height`
// function takes a node and a boolean flag indicating whether it is calculating the height of the left
// subtree or the right subtree. It recursively calculates the height by traversing the subtree and
// counting the number of levels. Finally, it returns the maximum height between the left and right
// subtrees
func (binaryTree *BinaryTree) HeightOfBinaryTreeNonRecursive() int {
	root := binaryTree.root

	if root != nil {
		leftHeight := height(root.left, true)
		rightHeight := height(root.right, false)
		return 1 + max(leftHeight, rightHeight)
	}
	return 0
}
