package main

import (
	"fmt"

	"example.com/trees"
)

const (
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorReset  = "\033[0m"
	colorCyan   = "\033[36m"
	colorGreen  = "\033[32m"
	colorBlue   = "\033[34m"
)

func printDottedLine() {
	fmt.Println(colorBlue + "---------------------------------------------------------------------------------------------------------------------------------------------------------------" + colorReset)
}

// The function "getATree" creates and returns a binary tree based on the given elements.
func getATree(elements []int) *trees.BinaryTree {
	binaryTree := &trees.BinaryTree{}
	for _, v := range elements {
		binaryTree.Insert(v)
	}
	return binaryTree
}

func treeTraversals() {
	fmt.Println("\t"+colorYellow+"Add 1, 2, 3, 4, 5, 6, 7, 8 & 9 into the Tree", colorReset)
	binaryTree := getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println("\t"+colorYellow+"Printing the Tree:", colorReset)
	binaryTree.PrintTree()
	printDottedLine()
	fmt.Println("\t"+colorYellow+"Printing the PreOrderTraversal", colorReset)
	binaryTree.PreOrder()

	fmt.Println("\t"+colorYellow+"Printing the InOrderTraversal", colorReset)
	binaryTree.InOrder()

	fmt.Println("\t"+colorYellow+"Printing the PostOrderTraversal", colorReset)
	binaryTree.PostOrder()

	fmt.Println("\t"+colorYellow+"Printing the LevelOrderTraversal", colorReset)
	binaryTree.LevelOrderTraversal()
	printDottedLine()
}

func problemOne() {
	fmt.Println("\t"+colorYellow+"Add 1, 2, 3, 4, 5, 6 & 7 into the Tree", colorReset)
	binaryTree := getATree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("\t"+colorYellow+"Printing the Tree:", colorReset)
	binaryTree.PrintTree()
	printDottedLine()
	fmt.Println("\t"+colorYellow+"Printing the Maximum element (Recursively)", colorReset, colorCyan, "\t", binaryTree.FindMaxElementUsingRecursion(), colorReset)

	emptyTree := getATree([]int{})
	fmt.Println("\t"+colorYellow+"Printing the Maximum element (Recursively) (Empty Tree)", colorReset, colorCyan, "\t", emptyTree.FindMaxElementUsingRecursion(), colorReset)

	printDottedLine()
}

func problemTwo() {
	fmt.Println("\t"+colorYellow+"Add 1, 2, 3, 4, 5, 6 & 7 into the Tree", colorReset)
	binaryTree := getATree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("\t"+colorYellow+"Printing the Tree:", colorReset)
	binaryTree.PrintTree()
	printDottedLine()
	fmt.Println("\t"+colorYellow+"Printing the Maximum element using Level Order", colorReset, colorCyan, "\t", binaryTree.FindMaxElementUsingLevelOrder(), colorReset)

	emptyTree := getATree([]int{})
	fmt.Println("\t"+colorYellow+"Printing the Maximum element using Level Order (Empty Tree)", colorReset, colorCyan, "\t", emptyTree.FindMaxElementUsingLevelOrder(), colorReset)

	printDottedLine()
}

func problemThree() {
	fmt.Println("\t"+colorYellow+"Add 1, 2, 3, 4, 5, 6 & 7 into the Tree", colorReset)
	binaryTree := getATree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("\t"+colorYellow+"Printing the Tree:", colorReset)
	binaryTree.PrintTree()
	printDottedLine()
	fmt.Println("\t"+colorYellow+"Searching the element 5 (Recursively)", colorReset, colorCyan, "\t", binaryTree.SearchElement(5), colorReset)

	fmt.Println("\t"+colorYellow+"Searching the element 9 (Recursively)", colorReset, colorCyan, "\t", binaryTree.SearchElement(9), colorReset)

	emptyTree := getATree([]int{})
	fmt.Println("\t"+colorYellow+"Searching the element 9 (Recursively) (Empty Tree)", colorReset, colorCyan, "\t", emptyTree.SearchElement(9), colorReset)

	printDottedLine()
}

func problemFour() {
	fmt.Println("\t"+colorYellow+"Add 1, 2, 3, 4, 5, 6 & 7 into the Tree", colorReset)
	binaryTree := getATree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("\t"+colorYellow+"Printing the Tree:", colorReset)
	binaryTree.PrintTree()
	printDottedLine()
	fmt.Println("\t"+colorYellow+"Searching the element 5 (Level Order)", colorReset, colorCyan, "\t", binaryTree.SearchElementUsingLevelOrder(5), colorReset)

	fmt.Println("\t"+colorYellow+"Searching the element 9 (Level Order)", colorReset, colorCyan, "\t", binaryTree.SearchElementUsingLevelOrder(9), colorReset)

	emptyTree := getATree([]int{})
	fmt.Println("\t"+colorYellow+"Searching the element 9 (Level Order) (Empty Tree)", colorReset, colorCyan, "\t", emptyTree.SearchElementUsingLevelOrder(9), colorReset)

	printDottedLine()
}

func problemFive() {
	fmt.Println("\t"+colorYellow+"Add 1, 2, 3, 4, 5, 6 & 7 into the Tree", colorReset)
	binaryTree := getATree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("\t"+colorYellow+"Printing the Tree:", colorReset)
	binaryTree.PrintTree()
	printDottedLine()
	fmt.Println("\t"+colorYellow+"Printing the LevelOrderTraversal", colorReset)
	binaryTree.LevelOrderTraversal()
	printDottedLine()
}

func problemSix() {
	fmt.Println("\t"+colorYellow+"Add 1, 2, 3, 4, 5, 6 & 7 into the Tree", colorReset)
	binaryTree := getATree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("\t"+colorYellow+"Printing the Tree:", colorReset)
	binaryTree.PrintTree()
	printDottedLine()
	fmt.Println("\t"+colorYellow+"Finding the size (No of nodes) in Binary Tree (Recursively)", colorReset, colorCyan, "\t", binaryTree.SizeOfBinaryTree(), colorReset)

	emptyTree := getATree([]int{})
	fmt.Println("\t"+colorYellow+"Finding the size (No of nodes) in Binary Tree (Recursively) (Empty Tree)", colorReset, colorCyan, "\t", emptyTree.SizeOfBinaryTree(), colorReset)

	printDottedLine()
}

func problemSeven() {
	fmt.Println("\t"+colorYellow+"Add 1, 2, 3, 4, 5, 6 & 7 into the Tree", colorReset)
	binaryTree := getATree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("\t"+colorYellow+"Printing the Tree:", colorReset)
	binaryTree.PrintTree()
	printDottedLine()
	fmt.Println("\t"+colorYellow+"Finding the size (No of nodes) in Binary Tree (Level Order)", colorReset, colorCyan, "\t", binaryTree.SizeOfBinaryTreeUsingLevelOrder(), colorReset)

	emptyTree := getATree([]int{})
	fmt.Println("\t"+colorYellow+"Finding the size (No of nodes) in Binary Tree (Level Order) (Empty Tree)", colorReset, colorCyan, "\t", emptyTree.SizeOfBinaryTreeUsingLevelOrder(), colorReset)

	printDottedLine()
}

func problemEight() {
	fmt.Println("\t"+colorYellow+"Add 1, 2, 3, 4, 5, 6 & 7 into the Tree", colorReset)
	binaryTree := getATree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("\t"+colorYellow+"Printing the Tree:", colorReset)
	binaryTree.PrintTree()
	printDottedLine()
	fmt.Println("\t" + colorYellow + "Print (Level Order) in reverse" + colorReset)
	binaryTree.LevelOrderInReverse()

	emptyTree := getATree([]int{})
	fmt.Println("\t" + colorYellow + "Print (Level Order) in reverse (Empty Tree)" + colorReset)
	emptyTree.LevelOrderInReverse()
	printDottedLine()
}

func problemNine() {
	fmt.Println("\t"+colorYellow+"Add 1, 2, 3, 4, 5, 6 & 7 into the Tree", colorReset)
	binaryTree := getATree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("\t"+colorYellow+"Printing the Tree (Before Delete)", colorReset)
	binaryTree.PrintTree()
	printDottedLine()
	fmt.Println("\t" + colorYellow + "Print (Level Order) traversal (Before Delete)" + colorReset)
	binaryTree.LevelOrderTraversal()
	fmt.Println("\t" + colorRed + "Deleting the tree" + colorReset)
	binaryTree.DeleteTree()
	fmt.Println("\t" + colorYellow + "Print (Level Order) traversal (After Delete)" + colorReset)
	printDottedLine()
}

func problemTen() {
	fmt.Println("\t"+colorYellow+"Add 1, 2, 3, 4, 5, 6 & 7 into the Tree", colorReset)
	// First Tree
	binaryTree := getATree([]int{1, 2, 3, 4, 5, 6, 7})
	binaryTree.LevelOrderTraversal()
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	binaryTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Print height of the tree (BOTH LEFT & RIGHT AT SAME LEVEL) [Recursive]" + colorReset)
	fmt.Println(colorCyan, "\tTree height is = ", binaryTree.HeithtOfBinaryTree(), colorReset)
	printDottedLine()

	// Second Tree
	newTree := getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	newTree.LevelOrderTraversal()
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Print height of the tree (LEFT SUBTREE HEIGHT > RIGHT SUBTREE HEIGHT) [Recursive]" + colorReset)
	fmt.Println(colorCyan, "\tTree height is = ", newTree.HeithtOfBinaryTree(), colorReset)
	printDottedLine()

}

func problemEleven() {
	fmt.Println("\t"+colorYellow+"Add 1, 2, 3, 4, 5, 6 & 7 into the Tree", colorReset)
	// First Tree
	binaryTree := getATree([]int{1, 2, 3, 4, 5, 6, 7})
	binaryTree.LevelOrderTraversal()
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	binaryTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Print height of the tree (BOTH LEFT & RIGHT AT SAME LEVEL) [Non-Recursive]" + colorReset)
	fmt.Println(colorCyan, "\tTree height is = ", binaryTree.HeightOfBinaryTreeNonRecursive(), colorReset)

	fmt.Println("\t" + colorYellow + "Print height of the tree (BOTH LEFT & RIGHT AT SAME LEVEL) [LevelOrder]" + colorReset)
	fmt.Println(colorCyan, "\tTree height is = ", binaryTree.HeightOfBTUsingLevelOrder(), colorReset)
	printDottedLine()

	// Second Tree
	newTree := getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	newTree.LevelOrderTraversal()
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Print height of the tree (LEFT SUBTREE HEIGHT > RIGHT SUBTREE HEIGHT) [Non-Recursive]" + colorReset)
	fmt.Println(colorCyan, "\tTree height is = ", newTree.HeightOfBinaryTreeNonRecursive(), colorReset)

	fmt.Println("\t" + colorYellow + "Print height of the tree (LEFT SUBTREE HEIGHT > RIGHT SUBTREE HEIGHT) [LevelOrder]" + colorReset)
	fmt.Println(colorCyan, "\tTree height is = ", newTree.HeightOfBTUsingLevelOrder(), colorReset)
	printDottedLine()
}

func problemTwelve() {
	newTree := getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	newTree.LevelOrderTraversal()
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Finding deepest node using (LevelOrder)" + colorReset)
	fmt.Println(colorCyan, "\tDeepest node is = ", newTree.DeepestNodeUsingLevelOrder(), colorReset)
	printDottedLine()
}

// TODO WIP
func problemThirteen() {
	newTree := getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	newTree.LevelOrderTraversal()
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Finding deepest node using (LevelOrder)" + colorReset)
	fmt.Println(colorCyan, "\tDeepest node is = ", newTree.DeepestNodeUsingLevelOrder(), colorReset)
	printDottedLine()
}

func problemFourteen() {
	newTree := getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	newTree.LevelOrderTraversal()
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Finding number of leaf nodes using (LevelOrder)" + colorReset)
	fmt.Println(colorCyan, "\tDeepest node is = ", newTree.NumberOfLeafNodesUsingLevelOrder(), colorReset)
	printDottedLine()
}

func problemFifteen() {
	newTree := getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	newTree.LevelOrderTraversal()
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Finding number of full nodes using (LevelOrder)" + colorReset)
	fmt.Println(colorCyan, "\tNumber of Full Nodes is = ", newTree.NumberOfFullNodesUsingLevelOrder(), colorReset)
	printDottedLine()
}

func problemSixteen() {
	newTree := getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Finding number of half nodes using (LevelOrder)" + colorReset)
	fmt.Println(colorCyan, "\tNumber of Half Nodes is = ", newTree.NumberOfHalfNodesUsingLevelOrder(), colorReset)

	newTree = getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	fmt.Println(colorCyan, "\tNumber of Half Nodes is = ", newTree.NumberOfHalfNodesUsingLevelOrder(), colorReset)
	printDottedLine()
}

func problemSeventeen() {
	fmt.Println("\t" + colorRed + "******************CHECK USING LEVEL ORDER******************" + colorReset)
	firstTree := getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	fmt.Println("\t"+colorYellow+"Printing the first Tree", colorReset)
	firstTree.PrintTree()
	secondTree := getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	fmt.Println("\t"+colorYellow+"Printing the second Tree", colorReset)
	secondTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Are 1st & 2nd trees structurally identical using (LevelOrder)" + colorReset)
	fmt.Println(colorCyan, "\tAre equal = ", firstTree.AreStructurallyIdentical(secondTree), colorReset)

	fmt.Println("\t" + colorYellow + "Are 2 empty trees structurally identical using (LevelOrder)" + colorReset)
	firstEmptyTree := getATree([]int{})
	secondEmptyTree := getATree([]int{})
	fmt.Println(colorCyan, "\tAre equal = ", firstEmptyTree.AreStructurallyIdentical(secondEmptyTree), colorReset)

	fmt.Println("\t" + colorYellow + "Are 1st (Non-Empty) & 2nd (Empty) trees structurally identical using (LevelOrder)" + colorReset)
	fmt.Println("\t"+colorYellow+"Printing the first Tree", colorReset)
	firstTree.PrintTree()
	fmt.Println(colorCyan, "\tAre equal = ", firstTree.AreStructurallyIdentical(secondEmptyTree), colorReset)

	fmt.Println("\t" + colorYellow + "Are 1st (Non-Empty) & 2nd (Non-Empty) trees structurally identical using (LevelOrder)" + colorReset)
	fmt.Println("\t"+colorYellow+"Printing the first Tree", colorReset)
	firstTree.PrintTree()
	thirdTree := getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println("\t"+colorYellow+"Printing the second Tree", colorReset)
	thirdTree.PrintTree()
	fmt.Println(colorCyan, "\tAre equal = ", firstTree.AreStructurallyIdentical(thirdTree), colorReset)

	fmt.Println("\t" + colorYellow + "Are 1st (Non-Empty) & 2nd (Non-Empty) trees structurally identical using (LevelOrder)" + colorReset)
	fmt.Println("\t"+colorYellow+"Printing the first Tree", colorReset)
	firstTree.PrintTree()
	fourthTree := getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	fmt.Println("\t"+colorYellow+"Printing the second Tree", colorReset)
	fourthTree.PrintTree()
	fmt.Println(colorCyan, "\tAre equal = ", firstTree.AreStructurallyIdentical(fourthTree), colorReset)
	printDottedLine()
	fmt.Println("\t" + colorRed + "******************CHECK USING RECURSIVE APPROACH******************" + colorReset)
	fmt.Println("\t"+colorYellow+"Printing the first Tree", colorReset)
	firstTree.PrintTree()
	fmt.Println("\t"+colorYellow+"Printing the second Tree", colorReset)
	secondTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Are 1st & 2nd trees structurally identical using (Recursive Approach)" + colorReset)
	fmt.Println(colorCyan, "\tAre equal = ", firstTree.AreStructurallyIdenticalRecursive(secondTree), colorReset)

	fmt.Println("\t" + colorYellow + "Are 2 empty trees structurally identical using (Recursive Approach)" + colorReset)
	fmt.Println(colorCyan, "\tAre equal = ", firstEmptyTree.AreStructurallyIdenticalRecursive(secondEmptyTree), colorReset)

	fmt.Println("\t" + colorYellow + "Are 1st (Non-Empty) & 2nd (Empty) trees structurally identical using (Recursive Approach)" + colorReset)
	fmt.Println("\t"+colorYellow+"Printing the first Tree", colorReset)
	firstTree.PrintTree()
	fmt.Println(colorCyan, "\tAre equal = ", firstTree.AreStructurallyIdenticalRecursive(secondEmptyTree), colorReset)

	fmt.Println("\t" + colorYellow + "Are 1st (Non-Empty) & 2nd (Non-Empty) trees structurally identical using (Recursive Approach)" + colorReset)
	fmt.Println("\t"+colorYellow+"Printing the first Tree", colorReset)
	firstTree.PrintTree()
	fmt.Println("\t"+colorYellow+"Printing the second Tree", colorReset)
	thirdTree.PrintTree()
	fmt.Println(colorCyan, "\tAre equal = ", firstTree.AreStructurallyIdenticalRecursive(thirdTree), colorReset)

	fmt.Println("\t" + colorYellow + "Are 1st (Non-Empty) & 2nd (Non-Empty) trees structurally identical using (Recursive Approach)" + colorReset)
	fmt.Println("\t"+colorYellow+"Printing the first Tree", colorReset)
	firstTree.PrintTree()
	fmt.Println("\t"+colorYellow+"Printing the second Tree", colorReset)
	fourthTree.PrintTree()
	fmt.Println(colorCyan, "\tAre equal = ", firstTree.AreStructurallyIdenticalRecursive(fourthTree), colorReset)
	printDottedLine()
}

func executeExercises() {
	fmt.Println("TREE EXERCISES")
	fmt.Println("Tree Traversals: PreOrder | InOrder | PostOrder")
	treeTraversals()
	fmt.Println("Problem 1: Give an algorithm for finding maximum element in binary tree.")
	problemOne()
	fmt.Println("Problem 2: Give an algorithm for finding maximum element in binary tree (Level Order)")
	problemTwo()
	fmt.Println("Problem 3: Give an algorithm for searching an element in binary tree (Recursively)")
	problemThree()
	fmt.Println("Problem 4: Give an algorithm for searching an element in binary tree (Level Order)")
	problemFour()
	fmt.Println("Problem 5: Give an algorithm for inserting an element into binary tree")
	problemFive()
	fmt.Println("Problem 6: Give an algorithm for finding the size of binary tree (Recursively)")
	problemSix()
	fmt.Println("Problem 7: Give an algorithm for finding the size of binary tree (Level Order)")
	problemSeven()
	fmt.Println("Problem 8: Give an algorithm for printing the level order data in reverse order. For example, the output for the below tree should be: 4 5 6 7 2 3 1")
	problemEight()
	fmt.Println("Problem 9: Give an algorithm for deleting the tree.")
	problemNine()
	fmt.Println("Problem 10: Give an algorithm for finding the height (or depth) of the binary tree.")
	problemTen()
	fmt.Println("Problem 11: Can we solve Problem-10 without recursion?")
	problemEleven()
	fmt.Println("Problem 12: Give an algorithm for finding the deepest node of the binary tree.")
	problemTwelve()
	// fmt.Println("Problem 13: Give an algorithm for deleting an element (assuming data is given) from binary tree.")
	// problemThirteen()
	fmt.Println("Problem 14: Give an algorithm for finding the number of leaves in the binary tree without using recursion.")
	problemFourteen()
	fmt.Println("Problem 15: Give an algorithm for finding the number of full nodes in the binary tree without using recursion.")
	problemFifteen()

	fmt.Println("Problem 16: Give an algorithm for finding the number of half nodes (nodes with only one child) in the binary tree without using recursion.")
	problemSixteen()
	fmt.Println("Problem 17: Given two binary trees, return true if they are structurally identical.")
	problemSeventeen()
}

func main() {
	executeExercises()
}
