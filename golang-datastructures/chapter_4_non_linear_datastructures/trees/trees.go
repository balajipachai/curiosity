package trees

import "fmt"

// The `TreeNode` type represents a node in a binary tree with an integer value and pointers to its
// left and right child nodes.
// @property {int} value - The value property of the TreeNode struct represents the value or data
// stored in the node. It can be of any data type depending on the requirements of the program. In this
// case, it is an integer.
// @property leftNode - `leftNode` is a pointer to the left child node of a binary tree node. In a
// binary tree, each node can have at most two children, a left child and a right child. The `leftNode`
// property of a node points to its left child node, which can be another
// @property rightNode - `rightNode` is a property of the `TreeNode` struct that represents the right
// child node of the current node. In a binary tree, each node can have at most two child nodes, a left
// child and a right child. The `rightNode` property points to the right child node of
type TreeNode struct {
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

// The BinarySearchTree type represents a binary search tree data structure with a root node.
// @property rootNode - The rootNode property is a pointer to the root node of a binary search tree. A
// binary search tree is a data structure where each node has at most two children, and the left child
// is always less than the parent, while the right child is always greater than the parent. The root
// node is the
type BinarySearchTree struct {
	rootNode *TreeNode
}

// This function is inserting a new element into a binary search tree. It creates a new `TreeNode` with
// the given `value` and sets its `leftNode` and `rightNode` pointers to `nil`. If the tree is empty
// (i.e., `rootNode` is `nil`), it sets the `rootNode` to the new node. Otherwise, it calls the
// `insertTreeNode` function to insert the new node into the appropriate position in the tree.
func (tree *BinarySearchTree) InsertElement(value int) {
	newNode := &TreeNode{value: value, leftNode: nil, rightNode: nil}
	if tree.rootNode == nil {
		tree.rootNode = newNode
	} else {
		insertTreeNode(tree.rootNode, newNode)
	}
}

// The function inserts a new node into a binary search tree based on its value.
func insertTreeNode(rootNode *TreeNode, newNode *TreeNode) {
	// if newNode.value is less than rootNode, then insert to left
	if newNode.value < rootNode.value {
		if rootNode.leftNode == nil {
			rootNode.leftNode = newNode
		} else {
			insertTreeNode(rootNode.leftNode, newNode)
		}
	}
	// if newNode.value is greater than rootNode, then insert to right
	if newNode.value > rootNode.value {
		if rootNode.rightNode == nil {
			rootNode.rightNode = newNode
		} else {
			insertTreeNode(rootNode.rightNode, newNode)
		}
	}
}

// `func (tree *BinarySearchTree) InOrderTraverseTree(function func(int))` is a method of the
// `BinarySearchTree` struct that performs an in-order traversal of the binary search tree. It takes a
// function as an argument, which is called on each node's value in the tree. The `inOrderTraverseTree`
// function is called with the root node of the tree and the provided function. The
// `inOrderTraverseTree` function recursively traverses the left subtree, calls the provided function
// on the current node's value, and then recursively traverses the right subtree. This results in the
// nodes being visited in ascending order of their values.
func (tree *BinarySearchTree) InOrderTraverseTree(function func(int)) {
	inOrderTraverseTree(tree.rootNode, function)
}

// The function performs an in-order traversal of a binary tree and applies a given function to each
// node's value.
func inOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		inOrderTraverseTree(treeNode.leftNode, function)
		function(treeNode.value)
		inOrderTraverseTree(treeNode.rightNode, function)
	}
}

// `func (tree *BinarySearchTree) PreOrderTraverseTree(function func(int))` is a method of the
// `BinarySearchTree` struct that performs a pre-order traversal of the binary search tree. It takes a
// function as an argument, which is called on each node's value in the tree. The
// `preOrderTraverseTree` function is called with the root node of the tree and the provided function.
// The `preOrderTraverseTree` function recursively traverses the left subtree, then the right subtree,
// and finally calls the provided function on the current node's value. This results in the nodes being
// visited in a pre-order traversal order.
func (tree *BinarySearchTree) PreOrderTraverseTree(function func(int)) {
	preOrderTraverseTree(tree.rootNode, function)
}

// The function performs a pre-order traversal of a binary tree and applies a given function to each
// node's value.
func preOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		function(treeNode.value)
		preOrderTraverseTree(treeNode.leftNode, function)
		preOrderTraverseTree(treeNode.rightNode, function)
	}
}

// `func (tree *BinarySearchTree) PostOrderTraverseTree(function func(int))` is a method of the
// `BinarySearchTree` struct that performs a post-order traversal of the binary search tree. It takes a
// function as an argument, which is called on each node's value in the tree. The
// `postOrderTraverseTree` function is called with the root node of the tree and the provided function.
// The `postOrderTraverseTree` function recursively traverses the left subtree, then the right subtree,
// and finally calls the provided function on the current node's value. This results in the nodes being
// visited in a post-order traversal order.
func (tree *BinarySearchTree) PostOrderTraverseTree(function func(int)) {
	postOrderTraverseTree(tree.rootNode, function)
}

// This function performs a post-order traversal of a binary tree and applies a given function to each
// node's value.
func postOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		postOrderTraverseTree(treeNode.leftNode, function)
		postOrderTraverseTree(treeNode.rightNode, function)
		function(treeNode.value)
	}
}

// `func (tree *BinarySearchTree) MinNode() *int` is a method of the `BinarySearchTree` struct that
// returns a pointer to the minimum value node in the binary search tree. It starts at the root node of
// the tree and traverses the left subtree until it reaches a leaf node (i.e., a node with no left
// child). The value of this node is then returned as a pointer to an integer. If the tree is empty
// (i.e., `rootNode` is `nil`), it returns a `nil` pointer.
func (tree *BinarySearchTree) MinNode() *int {
	treeNode := tree.rootNode
	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.leftNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.leftNode
	}
}

// `func (tree *BinarySearchTree) MaxNode() *int` is a method of the `BinarySearchTree` struct that
// returns a pointer to the maximum value node in the binary search tree. It starts at the root node of
// the tree and traverses the right subtree until it reaches a leaf node (i.e., a node with no right
// child). The value of this node is then returned as a pointer to an integer. If the tree is empty
// (i.e., `rootNode` is `nil`), it returns a `nil` pointer.
func (tree *BinarySearchTree) MaxNode() *int {
	treeNode := tree.rootNode

	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.rightNode == nil {
			return &treeNode.value
		}
		treeNode = treeNode.rightNode
	}
}

// `func (tree *BinarySearchTree) SearchNode(value int) bool` is a method of the `BinarySearchTree`
// struct that searches for a node with the given `value` in the binary search tree. It starts at the
// root node of the tree and recursively traverses the tree based on the value of the current node. If
// the value of the current node is less than the search value, it continues the search in the right
// subtree. If the value of the current node is greater than the search value, it continues the search
// in the left subtree. If the value of the current node is equal to the search value, it returns
// `true`. If the search value is not found in the tree, it returns `false`.
func (tree *BinarySearchTree) SearchNode(value int) bool {
	treeNode := tree.rootNode
	if treeNode == nil {
		return false
	}
	return searchNode(treeNode, value)
}

// The function searches for a node with a specific value in a binary search tree.
func searchNode(treeNode *TreeNode, value int) bool {
	if treeNode == nil {
		return false
	}
	if value < treeNode.value {
		return searchNode(treeNode.leftNode, value)
	}
	if value > treeNode.value {
		return searchNode(treeNode.rightNode, value)
	}
	return true
}

// The below code is defining a method called `RemoveNode` for a binary search tree data structure in
// the Go programming language. The method takes an integer value as input and attempts to remove the
// node with that value from the tree. It does this by calling a recursive function called `removeNode`
// on the root node of the tree.
func (tree *BinarySearchTree) RemoveNode(value int) {
	treeNode := tree.rootNode
	removeNode(treeNode, value)
}

// The function removes a node with a given value from a binary search tree.
func removeNode(treeNode *TreeNode, value int) *TreeNode {
	if treeNode == nil {
		return nil
	}
	if value < treeNode.value {
		treeNode.leftNode = removeNode(treeNode.leftNode, value)
		return treeNode
	}
	if value > treeNode.value {
		treeNode.rightNode = removeNode(treeNode.rightNode, value)
		return treeNode
	}
	// value == treeNode.value
	if treeNode.leftNode == nil && treeNode.rightNode == nil {
		treeNode = nil
		return nil
	}
	if treeNode.leftNode == nil {
		treeNode = treeNode.rightNode
		return treeNode
	}
	if treeNode.rightNode == nil {
		treeNode = treeNode.leftNode
		return treeNode
	}

	leftmostrightNode := &TreeNode{}
	leftmostrightNode = treeNode.rightNode
	for {
		// find smallest value on the right side
		if leftmostrightNode != nil && leftmostrightNode.leftNode != nil {
			leftmostrightNode = leftmostrightNode.leftNode
		} else {
			break
		}
	}
	treeNode.value = leftmostrightNode.value
	treeNode.rightNode = removeNode(treeNode.rightNode, treeNode.value)
	return treeNode
}

// The function "PrintTraversedValue" prints an integer value with a tab separator.
func PrintTraversedValue(value int) {
	fmt.Printf("%d\t", value)
}

// The `String()` function is a method of the `BinarySearchTree` struct that prints out the binary
// search tree in a readable format. It calls the `stringify()` function to traverse the tree in an
// in-order traversal and print out each node's value. The `fmt.Println()` statements before and after
// the `stringify()` function call are used to print out a separator before and after the tree's
// contents to make it easier to read.
func (tree *BinarySearchTree) String() {
	fmt.Println("************************************************")
	stringify(tree.rootNode, 0)
	fmt.Println("************************************************")
}

// The function recursively prints the values of a binary tree in a specific format.
func stringify(treeNode *TreeNode, level int) {
	if treeNode != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += " "
		}
		format += "*** > "
		level++
		stringify(treeNode.leftNode, level)
		fmt.Printf(format+"%d\n", treeNode.value)
		stringify(treeNode.rightNode, level)
	}
}
