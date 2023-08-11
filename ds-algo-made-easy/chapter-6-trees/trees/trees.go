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

// The below code is defining a method called "RearElement" for a struct type "Queue". This method
// returns the rear element of the queue, which is stored in the "elements" slice at the index
// specified by the "rear" variable.
func (queue *Queue) RearElement() *BinaryTreeNode {
	return queue.elements[queue.rear]
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

// The below code is defining a method called "CreateNew" for a struct type called "Stack" in the Go
// programming language. This method initializes a new stack by setting the top index to -1 and
// creating an empty dynamic slice to store elements of type "*BinaryTreeNode".
func (stack *Stack) CreateNew() {
	stack.top = -1
	stack.elements = make([]*BinaryTreeNode, 0) // Assigns memory to dynamic slice
}

// The below code is defining a method called `IsEmpty()` for a struct type `Stack`. This method checks
// if the `top` field of the `Stack` struct is equal to -1, indicating that the stack is empty. It
// returns a boolean value indicating whether the stack is empty or not.
func (stack *Stack) IsEmpty() bool {
	return stack.top == -1
}

// The below code is defining a method called "Push" for a stack data structure in the Go programming
// language. This method takes a pointer to a BinaryTreeNode as a parameter and adds it to the stack.
// It does this by incrementing the "top" variable of the stack and appending the node to the
// "elements" slice of the stack.
func (stack *Stack) Push(node *BinaryTreeNode) {
	stack.top += 1
	stack.elements = append(stack.elements, node)
}

// The below code is implementing the Pop() function for a stack data structure in the Go programming
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

// The below code is defining a method called "TopElement" for a struct type called "Stack" in the Go
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
func searchElementHelper(treeNode *BinaryTreeNode, element int) *BinaryTreeNode {
	if treeNode != nil {
		data := treeNode.data
		if data == element {
			return treeNode
		} else {
			isFound := searchElementHelper(treeNode.left, element)
			if isFound != nil {
				return isFound
			} else {
				return searchElementHelper(treeNode.right, element)
			}
		}
	}
	return nil
}

// The `SearchElement` method of the `BinaryTree` struct is used to search for a specific element in
// the binary tree.
func (binaryTree *BinaryTree) SearchElement(element int) *BinaryTreeNode {
	if binaryTree.root != nil {
		return searchElementHelper(binaryTree.root, element)
	}
	return nil
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

// The below code is implementing a method called `SizeOfBinaryTreeUsingLevelOrder()` for a binary tree
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

// The below code is implementing a method called `LevelOrderInReverse()` for a binary tree data
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

// The below code is defining a method called `DeleteTree` for a `BinaryTree` struct in the Go
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

// The below code is implementing a method called "HeightOfBinaryTree" for a binary tree data
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

// The below code is implementing a non-recursive method to calculate the height of a binary tree. It
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

// The below code is calculating the height of a binary tree using the level order traversal algorithm.
// It uses a queue to perform the traversal and keeps track of the current level. The height of the
// binary tree is determined by the number of levels traversed.
func (binaryTree *BinaryTree) HeightOfBTUsingLevelOrder() int {
	queue := &Queue{}
	queue.CreateNew()
	level := 0
	if binaryTree.root != nil {
		queue.EnQueue(binaryTree.root)
		queue.EnQueue(nil) // Completion of root level

		for !queue.IsEmpty() {
			root := queue.DeQueue()

			if root == nil {
				if !queue.IsEmpty() {
					queue.EnQueue(nil)
				}
				level++
			} else {
				if root.left != nil {
					queue.EnQueue(root.left)
				}
				if root.right != nil {
					queue.EnQueue(root.right)
				}
			}
		}
	}
	return level
}

// The below code is implementing a method called `DeepestNodeUsingLevelOrder()` for a binary tree data
// structure. This method finds and returns the deepest node in the binary tree using a level order
// traversal approach.
func (binaryTree *BinaryTree) DeepestNodeUsingLevelOrder() *BinaryTreeNode {
	queue := &Queue{}
	queue.CreateNew()

	nodeToReturn := &BinaryTreeNode{}

	if binaryTree.root != nil {
		queue.EnQueue(binaryTree.root)

		for !queue.IsEmpty() {
			node := queue.DeQueue()

			if node.left != nil {
				queue.EnQueue(node.left)
			}

			if node.right != nil {
				queue.EnQueue(node.right)
			}

			nodeToReturn = node
		}
	}
	return nodeToReturn
}

// The below code is implementing a method called `NumberOfLeafNodesUsingLevelOrder()` for a binary
// tree data structure. This method calculates and returns the number of leaf nodes in the binary tree
// using a level order traversal approach.
func (binaryTree *BinaryTree) NumberOfLeafNodesUsingLevelOrder() int {
	queue := &Queue{}
	queue.CreateNew()

	count := 0

	if binaryTree.root != nil {
		queue.EnQueue(binaryTree.root)

		for !queue.IsEmpty() {
			node := queue.DeQueue()

			if node.left == nil && node.right == nil {
				count++
			}

			if node.left != nil {
				queue.EnQueue(node.left)
			}

			if node.right != nil {
				queue.EnQueue(node.right)
			}
		}
	}
	return count
}

// The below code is implementing a method called NumberOfFullNodesUsingLevelOrder() for a binary tree
// data structure. This method calculates and returns the number of full nodes in the binary tree using
// a level order traversal approach.
func (binaryTree *BinaryTree) NumberOfFullNodesUsingLevelOrder() int {
	queue := &Queue{}
	queue.CreateNew()

	count := 0

	if binaryTree.root != nil {
		queue.EnQueue(binaryTree.root)

		for !queue.IsEmpty() {
			node := queue.DeQueue()

			if node.left != nil && node.right != nil {
				count++
			}

			if node.left != nil {
				queue.EnQueue(node.left)
			}

			if node.right != nil {
				queue.EnQueue(node.right)
			}
		}
	}
	return count
}

// The below code is implementing a method called `NumberOfHalfNodesUsingLevelOrder()` for a binary
// tree data structure. This method calculates and returns the number of half nodes in the binary tree
// using a level order traversal approach.
func (binaryTree *BinaryTree) NumberOfHalfNodesUsingLevelOrder() int {
	count := 0
	if binaryTree.root != nil {
		queue := &Queue{}
		queue.CreateNew()

		queue.EnQueue(binaryTree.root)

		for !queue.IsEmpty() {
			node := queue.DeQueue()

			if node.left == nil || node.right == nil {
				if node.left != nil || node.right != nil {
					count++
				}
			}

			if node.left != nil {
				queue.EnQueue(node.left)
			}

			if node.right != nil {
				queue.EnQueue(node.right)
			}
		}
	}
	return count
}

// The below code is implementing a method called `AreStructurallyIdentical` for a binary tree data
// structure in the Go programming language. This method checks if two binary trees are structurally
// identical, meaning they have the same structure (same number of nodes and same arrangement of nodes)
// regardless of the data stored in the nodes.
func (firstTree *BinaryTree) AreStructurallyIdentical(secondTree *BinaryTree) bool {
	/*
		Algorithm:
		- If both trees are NULL then return true.
		- If both trees are not NULL, then compare data and recursively check left and right subtree structures.
	*/

	areEqual := true

	if firstTree.root == nil && secondTree.root == nil {
		areEqual = true
	}

	queue := &Queue{}

	firstQueue := &Queue{}
	firstQueue.CreateNew()

	secondQueue := &Queue{}
	secondQueue.CreateNew()

	if firstTree.root != nil {
		queue.EnQueue(firstTree.root)

		for !queue.IsEmpty() {
			element := queue.DeQueue()
			firstQueue.EnQueue(element)
			if element.left != nil {
				queue.EnQueue(element.left)
			}

			if element.right != nil {
				queue.EnQueue(element.right)
			}
		}
	}

	if secondTree.root != nil {
		queue.EnQueue(secondTree.root)

		for !queue.IsEmpty() {
			element := queue.DeQueue()
			secondQueue.EnQueue(element)
			if element.left != nil {
				queue.EnQueue(element.left)
			}

			if element.right != nil {
				queue.EnQueue(element.right)
			}
		}
	}

	if firstQueue.IsEmpty() && !secondQueue.IsEmpty() || secondQueue.IsEmpty() && !firstQueue.IsEmpty() {
		areEqual = false
	}

	fmt.Println(colorMagenta)
	// Check the Elements
	for !firstQueue.IsEmpty() && !secondQueue.IsEmpty() {
		firstNode := firstQueue.DeQueue()
		secondNode := secondQueue.DeQueue()

		if firstQueue.IsEmpty() && !secondQueue.IsEmpty() || secondQueue.IsEmpty() && !firstQueue.IsEmpty() {
			areEqual = false
		}

		fmt.Printf("\t%v\t%v\n", firstNode, secondNode)

		if firstNode.data != secondNode.data {
			areEqual = false
			break
		}
	}
	fmt.Println(colorReset)

	// Print the last element in case of non-empty trees being non-equal
	if !firstQueue.IsEmpty() {
		fmt.Println("\tFirst queue is non-empty")
		fmt.Println("\t", colorMagenta, firstQueue.DeQueue(), colorReset)
	}

	if !secondQueue.IsEmpty() {
		fmt.Println("\tSecond queue is non-empty")
		fmt.Println("\t", colorMagenta, secondQueue.DeQueue(), colorReset)
	}

	return areEqual
}

// The function checks if two binary trees are structurally identical by recursively comparing their
// nodes.
func areStructurallyIdenticalRecursiveHelper(root1, root2 *BinaryTreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	// else traverse both right and left tree recursively
	return (root1.data == root2.data && areStructurallyIdenticalRecursiveHelper(root1.left, root2.left) && areStructurallyIdenticalRecursiveHelper(root1.right, root2.right))
}

// The below code is defining a method called "AreStructurallyIdenticalRecursive" for a BinaryTree
// struct in the Go programming language. This method takes another BinaryTree as a parameter and
// returns a boolean value indicating whether the two binary trees are structurally identical.
func (firstTree *BinaryTree) AreStructurallyIdenticalRecursive(secondTree *BinaryTree) bool {
	return areStructurallyIdenticalRecursiveHelper(firstTree.root, secondTree.root)
}

// The function `diameterHelper` calculates the diameter of a binary tree and updates the maximum diameter found so far.
func diameterHelper(root *BinaryTreeNode, ptr *int) int {
	if root == nil {
		return 0
	}
	left := diameterHelper(root.left, ptr)
	right := diameterHelper(root.right, ptr)

	if left+right > *ptr {
		*ptr = left + right
	}
	return max(left, right) + 1
}

// The function calculates the height of a binary tree recursively.
func heightRecursive(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(heightRecursive(root.left), heightRecursive(root.right))
}

// The function `diameterRecursive` calculates the diameter of a binary tree using a recursive
// approach.
func diameterRecursive(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	return max(diameterRecursive(root.left), diameterRecursive(root.right))
}

// The below code is implementing a method called "Diameter" for a binary tree data structure.
// The method calculates the diameter of the binary tree, which is defined as the number of nodes on the longest path between any two leaf nodes in the tree.
func (binaryTree *BinaryTree) Diameter() int {
	/*
		Algorithm:
		1. Calculate diameter of left subtree and right subtree recursively.
		2. Among these two values, send maximum value along with current level (+1)
	*/
	d := 0
	maxOf := max(heightRecursive(binaryTree.root), diameterRecursive(binaryTree.root))
	fmt.Println(colorMagenta, "\tDiameter using alternate approach = : ", maxOf, colorReset)
	return diameterHelper(binaryTree.root, &d)
}

// The below code is implementing a method called `LevelWithMaximumSumUsingLevelOrder()` for a binary
// tree data structure. This method calculates the level in the binary tree that has the maximum sum of
// node values. It uses the level order traversal technique to visit each node in the tree level by
// level. It keeps track of the current sum of node values for each level and updates the maximum sum
// if a higher sum is found. Finally, it returns the maximum sum.
func (binaryTree *BinaryTree) LevelWithMaximumSumUsingLevelOrder() (int, int) {
	if binaryTree.root == nil {
		return 0, 0
	}
	level := 0
	currentSum := 0
	maxSum := 0

	queue := &Queue{}
	queue.CreateNew()

	queue.EnQueue(binaryTree.root)
	queue.EnQueue(nil) // Marking the end of level

	for !queue.IsEmpty() {
		root := queue.DeQueue()
		// Level logic
		if root == nil {
			if !queue.IsEmpty() {
				queue.EnQueue(nil)
			}
			if currentSum > maxSum {
				maxSum = currentSum
			}
			level++
			currentSum = 0
		} else {
			currentSum += root.data
			if root.left != nil {
				queue.EnQueue(root.left)
			}

			if root.right != nil {
				queue.EnQueue(root.right)
			}
		}
	}
	return maxSum, level
}

// The function recursively prints all root-to-leaf paths in a binary tree.
func allRootToLeafPathsHelper(root *BinaryTreeNode, paths []int) {
	if root == nil {
		return
	}
	paths = append(paths, root.data)
	// Below condition implies we have hit the leaf node
	if root.left == nil && root.right == nil {
		// Print the paths array
		fmt.Println(colorMagenta)
		for i, v := range paths {
			if i != len(paths)-1 {
				fmt.Printf("\t%d\t->", v)
			} else {
				fmt.Printf("\t%d", v)
			}
		}
		fmt.Println(colorReset)
	} else {
		allRootToLeafPathsHelper(root.left, paths)
		allRootToLeafPathsHelper(root.right, paths)
	}
}

/*
Algorithm: To print out all root-to-leaf-paths
If we have a binary tree having nodes: 1, 2, 3, 4, 5, 6 & 7

	     1
	   /   \
	  2     3
	 / \   / \
	4   5 6   7

then all root-to-leaf-paths are:

	> 1 - 2 - 4
	> 1 - 2 - 5
	> 1 - 3 - 6
	> 1 - 3 - 7

How we accomplish this? We will use a recursive approach.
Base Condition:

	If root == nil then return
	If root.left == nil && root.right == nil, we have hit the end of path, print the path
	If root.left != nil & root.right != nil, solve them recursively
*/
// The below code is defining a method called "AllRootToLeafPaths" for a binary tree data structure.
// This method is used to find all the paths from the root to the leaf nodes of the binary tree. It
// initializes an empty slice called "paths" to store the paths. Then, it calls a helper function
// called "allRootToLeafPathsHelper" passing the root node of the binary tree and the empty "paths"
// slice as arguments.
func (binaryTree *BinaryTree) AllRootToLeafPaths() {
	paths := make([]int, 0)
	allRootToLeafPathsHelper(binaryTree.root, paths)
}

// The function checks if there exists a path in a binary tree from root to leaf nodes that adds up to
// a given sum.
func existenceOfPathWithGivenSumHelper(root *BinaryTreeNode, sum int) bool {
	if root == nil {
		return sum == 0
	} else {
		remainingSum := sum - root.data
		if (root.left != nil && root.right != nil) || (root.left == nil && root.right == nil) {
			return existenceOfPathWithGivenSumHelper(root.left, remainingSum) || existenceOfPathWithGivenSumHelper(root.right, remainingSum)
		} else if root.left != nil {
			return existenceOfPathWithGivenSumHelper(root.left, remainingSum)
		} else {
			return existenceOfPathWithGivenSumHelper(root.right, remainingSum)
		}
	}
}

// The below code is defining a method called "ExistenceOfPathWithGivenSum" for a binary tree data
// structure. This method takes an integer parameter called "sum" and returns a boolean value.
func (binaryTree *BinaryTree) ExistenceOfPathWithGivenSum(sum int) bool {
	return existenceOfPathWithGivenSumHelper(binaryTree.root, sum)
}

// The addHelper function recursively calculates the sum of all the nodes in a binary tree.
func addHelper(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	} else {
		return (root.data + addHelper(root.left) + addHelper(root.right))
	}
}

// The below code is defining a method called "Add" for a binary tree data structure. This method takes
// no arguments and returns an integer. It calls a helper function called "addHelper" passing in the
// root node of the binary tree. The purpose of this method is to add up the values of all nodes in the
// binary tree and return the sum.
func (binaryTree *BinaryTree) Add() int {
	return addHelper(binaryTree.root)
}

// The below code is implementing a method called `AddUsingLevelOrder` for a binary tree. This method
// calculates the sum of all the nodes in the binary tree using a level order traversal. It uses a
// queue data structure to perform the traversal. The method starts by checking if the root of the
// binary tree is nil, and if so, it returns 0. Otherwise, it creates a new queue and enqueues the root
// node. It then initializes a variable `sum` to 0.
func (binaryTree *BinaryTree) AddUsingLevelOrder() int {
	root := binaryTree.root

	if root == nil {
		return 0
	}

	queue := &Queue{}
	queue.CreateNew()

	queue.EnQueue(root)

	sum := 0

	for !queue.IsEmpty() {
		node := queue.DeQueue()
		sum += node.data
		if node.left != nil {
			queue.EnQueue(node.left)
		}
		if node.right != nil {
			queue.EnQueue(node.right)
		}
	}
	return sum
}

// The mirrorHelper function recursively swaps the left and right child of each node in a binary tree.
func mirrorHelper(root *BinaryTreeNode) *BinaryTreeNode {
	if root == nil {
		return nil
	}
	mirrorHelper(root.left)
	mirrorHelper(root.right)
	// Swap the right and left child
	root.right, root.left = root.left, root.right
	return root
}

// The below code is defining a method called "Mirror" for a binary tree data structure. This method
// returns a pointer to a BinaryTreeNode, which represents the root of the mirrored binary tree. The
// method calls a helper function called "mirrorHelper" and passes the root of the original binary tree
// as an argument. The "mirrorHelper" function recursively mirrors the binary tree by swapping the left
// and right child nodes of each node in the tree.
func (binaryTree *BinaryTree) Mirror() *BinaryTreeNode {
	return mirrorHelper(binaryTree.root)
}

// The function checks if two binary trees are mirrors of each other.
func areMirrorsHelper(root1, root2 *BinaryTreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.data != root2.data {
		return false
	}

	return areMirrorsHelper(root1.left, root2.right) && areMirrorsHelper(root1.right, root2.left)
}

// The below code is defining a method called "AreMirrors"
// This method takes another BinaryTree as input and returns a boolean value.
// It calls a helper function called "areMirrorsHelper" with the root nodes of the two BinaryTrees as arguments to check if the two trees are mirrors of each other.
func (firstTree *BinaryTree) AreMirrors(secondTree *BinaryTree) bool {
	return areMirrorsHelper(firstTree.root, secondTree.root)
}

// The function lcaHelper finds the lowest common ancestor of two nodes in a binary tree.
func lcaHelper(root, node1, node2 *BinaryTreeNode) *BinaryTreeNode {
	if root == nil || root == node1 || root == node2 {
		return root
	}
	left := lcaHelper(root.left, node1, node2)
	right := lcaHelper(root.right, node1, node2)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	} else {
		return right
	}

}

// The below code is defining a method called LCA (Lowest Common Ancestor) for a BinaryTree.
// This method takes in two BinaryTreeNode pointers, node1 and node2, and returns a pointer
// to the lowest common ancestor of these two nodes in the binary tree.
// The method calls a helper function called lcaHelper, passing in the root of the binary tree and the two nodes.
func (binaryTree *BinaryTree) LCA(node1, node2 *BinaryTreeNode) *BinaryTreeNode {
	return lcaHelper(binaryTree.root, node1, node2)
}

// The function builds a binary tree using pre-order and in-order traversal arrays.
func buildBinaryTreeHelper(preOrder, inOrder []int, inOrderStart, inOrderEnd int) *BinaryTreeNode {
	preOrderIndex := 0
	newNode := &BinaryTreeNode{}

	if inOrderStart > inOrderEnd {
		return nil
	}

	newNode.data = preOrder[preOrderIndex]
	preOrderIndex++

	if inOrderStart == inOrderEnd {
		fmt.Println(newNode)
		return newNode
	}

	var inOrderIndex int

	for i, v := range inOrder {
		if newNode.data == v {
			inOrderIndex = i
			break
		}
	}

	newNode.left = buildBinaryTreeHelper(preOrder, inOrder, inOrderStart, inOrderIndex-1)
	newNode.right = buildBinaryTreeHelper(preOrder, inOrder, inOrderIndex+1, inOrderEnd)

	fmt.Println(newNode)
	return newNode
}

// The below code is defining a method called "BuildBinaryTree"
// This method takes in two slices of integers, "preOrder" and "inOrder",
// as well as two integers "preOrderIndex" and "inOrderIndex".
// It then calls a helper function called "buildBinaryTreeHelper" with these parameters and returns the result.
func (binaryTree *BinaryTree) BuildBinaryTree(preOrder, inOrder []int, preOrderIndex, inOrderIndex int) *BinaryTreeNode {
	return buildBinaryTreeHelper(preOrder, inOrder, preOrderIndex, inOrderIndex)
}

// The function calculates the vertical sum of a binary tree by recursively traversing the tree and
// updating a hash table with the sum of nodes at each horizontal distance.
func verticalSumInBinaryTreeHelper(root *BinaryTreeNode, horizontalDistance int, hashTable map[int]int) {
	if root == nil {
		return
	}
	verticalSumInBinaryTreeHelper(root.left, horizontalDistance-1, hashTable)
	hashTable[horizontalDistance] += root.data
	verticalSumInBinaryTreeHelper(root.right, horizontalDistance+1, hashTable)
}

// The below code is defining a method called "VerticalSumInBinaryTree"
// This method takes two parameters: "horizontalDistance" of type int and "hashTable" of type map[int]int.
func (binaryTree *BinaryTree) VerticalSumInBinaryTree(horizontalDistance int, hashTable map[int]int) {
	verticalSumInBinaryTreeHelper(binaryTree.root, horizontalDistance, hashTable)
}

/* GENERIC TREES (N-ary Trees)

A generic tree can be represented as a binary tree using the below representation

type struct NaryTreeNode {
	data int
	firstChild *NaryTreeNode
	nextSibling *NaryTreeNode
}

During construction we need to specify whether the node has firstChild or nextSibling or both and then the rest operations of binary tree
are applicable on `NaryTree` imagine left => firstChild & right => nextSibling
*/

// Given a parent array P, where P[i] indicates the parent of ith node in the tree (assume parent of root node is indicated with –1).
// Give an algorithm for finding the height or depth of the tree.
func GenericTreeDepth(parentArray []int) int {
	currentDepth := -1
	maxDepth := -1
	for i := range parentArray {
		currentDepth = 0
		j := i

		for parentArray[j] != -1 {
			currentDepth++
			j = parentArray[j]
		}

		if currentDepth > maxDepth {
			maxDepth = currentDepth
		}
	}
	return maxDepth
}

// The IsIsomorphicHelper function checks if two binary trees are isomorphic by recursively comparing
// their left and right subtrees.
func IsIsomorphicHelper(root1, root2 *BinaryTreeNode) bool {
	// Return true for Empty Tree
	if root1 == nil && root2 == nil {
		return true
	}

	// Return false when either of the node is nil
	if (root1 == nil && root2 != nil) || (root1 != nil && root2 == nil) {
		return false
	}
	return IsIsomorphicHelper(root1.left, root2.left) && IsIsomorphicHelper(root1.right, root2.right)
}

// Two binary trees root1 and root2 are isomorphic if they have the same structure.
// The values of the nodes does not affect whether two trees are isomorphic or not.
func (tree1 *BinaryTree) IsIsomorphic(tree2 *BinaryTree) bool {
	return IsIsomorphicHelper(tree1.root, tree2.root)
}

// The function isQuasiIsomorphicHelper checks if two binary trees are quasi-isomorphic, meaning that
// they have the same structure but the values of the nodes can be different.
func isQuasiIsomorphicHelper(root1, root2 *BinaryTreeNode) bool {
	// Return true for Empty Tree
	if root1 == nil && root2 == nil {
		return true
	}

	// Return false when either of the node is nil
	if (root1 == nil && root2 != nil) || (root1 != nil && root2 == nil) {
		return false
	}

	return ((isQuasiIsomorphicHelper(root1.left, root2.left)) && (isQuasiIsomorphicHelper(root1.right, root2.right))) || ((isQuasiIsomorphicHelper(root1.left, root2.right)) && (isQuasiIsomorphicHelper(root1.right, root2.left)))

}

// Two trees root1 and root2 are quasi-isomorphic if root1 can be transformed into root2
// by swapping the left and right children of some of the nodes of root1.
// Data in the nodes are not important in determining quasi-isomorphism; only the shape is important.
func (tree1 *BinaryTree) IsQuasiIsomorphic(tree2 *BinaryTree) bool {
	return isQuasiIsomorphicHelper(tree1.root, tree2.root)
}

// The function CheckQuasiIsomorphism checks if two binary trees are quasi-isomorphic.
func CheckQuasiIsomorphism() bool {
	t1Node1 := &BinaryTreeNode{
		data: 1,
	}
	t1Node2 := &BinaryTreeNode{
		data: 2,
	}
	t1Node3 := &BinaryTreeNode{
		data: 3,
	}
	t1Node4 := &BinaryTreeNode{
		data: 4,
	}
	t1Node5 := &BinaryTreeNode{
		data: 5,
	}

	/*
				Tree1
		        1
			   / \
			  3   2
			   \   \
			    5   4

	*/
	// Link the above nodes and create a tree
	t1Node1.left = t1Node2
	t1Node1.right = t1Node3
	t1Node3.right = t1Node4
	t1Node2.right = t1Node5

	t2Node1 := &BinaryTreeNode{
		data: 6,
	}
	t2Node2 := &BinaryTreeNode{
		data: 7,
	}
	t2Node3 := &BinaryTreeNode{
		data: 8,
	}
	t2Node4 := &BinaryTreeNode{
		data: 9,
	}
	t2Node5 := &BinaryTreeNode{
		data: 10,
	}

	/*
			Tree2
		    6
		   / \
		  7   8
		 /     \
		9       10

	*/
	t2Node1.left = t2Node2
	t2Node1.right = t2Node3
	t2Node2.left = t2Node4
	t2Node3.right = t2Node5

	printTreeWithIndent(t1Node1, 0, "")
	printTreeWithIndent(t2Node1, 0, "")

	return isQuasiIsomorphicHelper(t1Node1, t2Node1)

}
