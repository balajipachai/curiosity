package main

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

func main() {
	tree := &BinarySearchTree{}
	tree.InsertElement(8)
	tree.InsertElement(3)
	tree.InsertElement(10)
	tree.InsertElement(1)
	tree.InsertElement(6)
	tree.String()
}
