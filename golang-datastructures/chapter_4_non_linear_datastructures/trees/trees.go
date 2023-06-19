package trees

type TreeNode struct {
	value     int
	leftNode  *TreeNode
	rightNode *TreeNode
}

type BinarySearchTree struct {
	rootNode *TreeNode
}

func (tree *BinarySearchTree) InsertElement(value int) {
	newNode := &TreeNode{value: value, leftNode: nil, rightNode: nil}
	if tree.rootNode == nil {
		tree.rootNode = newNode
	} else {
		insertTreeNode(tree.rootNode, newNode)
	}
}

func insertTreeNode(rootNode *TreeNode, newNode *TreeNode) {
	// if newNode.value is less than rootNode, then insert to left

	// if newNode.value is greater than rootNode, then insert to right
}
