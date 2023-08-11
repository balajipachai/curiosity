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

// The function `treeTraversals` creates a binary tree and performs various traversals on it.
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

// The function problemOne creates a binary tree and prints its elements, then finds and prints the
// maximum element using recursion.
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

// The function problemTwo creates a binary tree and prints its elements, then finds and prints the
// maximum element using level order traversal.
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

// The function problemThree creates a binary tree, prints the tree, and searches for specific elements
// in the tree.
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

// The function problemFour creates a binary tree, prints the tree, and searches for elements using
// level order traversal.
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

// The function problemFive creates a binary tree and performs various operations on it, such as
// printing the tree and performing level order traversal.
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

// The function problemSix creates a binary tree, prints the tree, and then finds the size (number of
// nodes) of the binary tree recursively.
// The function problemSix creates a binary tree, prints the tree, and then finds the size (number of
// nodes) of the binary tree recursively.
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

// The function problemSeven creates a binary tree with values 1, 2, 3, 4, 5, 6, and 7, prints the
// tree, and then finds the size (number of nodes) of the tree using level order traversal.
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

// The function problemEight creates a binary tree and prints its elements in reverse order using level
// order traversal.
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

// The function problemNine creates a binary tree, prints the tree before deleting it, deletes the
// tree, and then prints the tree after deletion.
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

// The function problemTen creates two binary trees and performs various operations on them, including
// adding elements, printing the tree, and calculating the height of the tree.
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

// The function problemEleven creates two binary trees and prints their level order traversal, tree
// structure, and height using both non-recursive and level order methods.
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

// The function problemTwelve creates a binary tree, performs a level order traversal, prints the tree,
// and finds the deepest node using level order traversal.
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

// The function problemFourteen creates a binary tree, performs a level order traversal, prints the
// tree, and finds the number of leaf nodes using level order traversal.
func problemFourteen() {
	newTree := getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	newTree.LevelOrderTraversal()
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Finding number of leaf nodes using (LevelOrder)" + colorReset)
	fmt.Println(colorCyan, "\tDeepest node is = ", newTree.NumberOfLeafNodesUsingLevelOrder(), colorReset)
	printDottedLine()
}

// The function problemFifteen creates a binary tree, performs a level order traversal, prints the
// tree, and then finds the number of full nodes in the tree using level order traversal.
func problemFifteen() {
	newTree := getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	newTree.LevelOrderTraversal()
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Finding number of full nodes using (LevelOrder)" + colorReset)
	fmt.Println(colorCyan, "\tNumber of Full Nodes is = ", newTree.NumberOfFullNodesUsingLevelOrder(), colorReset)
	printDottedLine()
}

// The function problemSixteen creates a binary tree, prints the tree, and then calculates and prints
// the number of half nodes in the tree using level order traversal.
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

// The function problemSeventeen checks if two trees are structurally identical using both a level
// order approach and a recursive approach.
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

// The function "eighteen" creates a binary tree, prints the tree, and calculates the diameter of the
// tree recursively.
func eighteen() {
	newTree := getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Finding Diameter of the binary tree(Recursively)" + colorReset)
	fmt.Println(colorCyan, "\tDiameter of the tree is = ", newTree.Diameter(), colorReset)
	printDottedLine()
}

// The function "nineteen" creates a binary tree, prints the tree, and finds the level with the maximum
// sum using level order traversal.
func nineteen() {
	newTree := getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Finding Level with maximum sum in the binary tree(LevelOrder)" + colorReset)
	maxSum, level := newTree.LevelWithMaximumSumUsingLevelOrder()
	fmt.Println(colorCyan, "\tLevel = ", level, " has maximum sum = ", maxSum, colorReset)
	printDottedLine()
}

// The function "twenty" creates a binary tree, prints the tree, and then prints all root-to-leaf paths
// recursively.
func twenty() {
	newTree := getATree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Printing all root-to-leaf-paths(Recursively)" + colorReset)
	newTree.AllRootToLeafPaths()
	printDottedLine()
}

// The function "twentyOne" creates a binary tree, prints the tree, prints all root-to-leaf paths, and
// checks if a path exists with a given sum.
func twentyOne() {
	newTree := getATree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Printing all root-to-leaf-paths(Recursively)" + colorReset)
	newTree.AllRootToLeafPaths()
	fmt.Println("\t" + colorYellow + "Checks if a path exists in Binary Tree with sum = 18 (Recursively)" + colorReset)
	fmt.Println(colorCyan, "\tDoes a path exists with sum = 18 ===> ", newTree.ExistenceOfPathWithGivenSum(18), colorReset)
	fmt.Println(colorCyan, "\tDoes a path exists with sum = 7 ===> ", newTree.ExistenceOfPathWithGivenSum(7), colorReset)
	printDottedLine()
}

// The function `twentyTwoAndTwentyThree` prints a binary tree, and then calculates the sum of all
// elements in the tree using both recursive and level order algorithms.
func twentyTwoAndTwentyThree() {
	newTree := getATree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Give an algorithm for finding the sum of all elements in binary tree (Recursively)" + colorReset)
	fmt.Println(colorCyan, "\tSum\t=\t ", newTree.Add(), colorReset)
	fmt.Println("\t" + colorYellow + "Give an algorithm for finding the sum of all elements in binary tree (LevelOrder)" + colorReset)
	fmt.Println(colorCyan, "\tSum\t=\t ", newTree.AddUsingLevelOrder(), colorReset)
}

// The function `twentyFour` creates a binary tree, prints it, mirrors it, and then prints the mirrored
// tree.
func twentyFour() {
	newTree := getATree([]int{1, 2, 3, 4, 5})
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	fmt.Println(colorCyan, "\tMirroring the tree", colorReset)
	newTree.Mirror()
	newTree.PrintTree()
	printDottedLine()
}

// The function `twentyFive` creates two binary trees, mirrors one of them, and then checks if the
// original tree and the mirrored tree are mirrors of each other.
func twentyFive() {
	newTree := getATree([]int{1, 2, 3, 4, 5})
	originalTree := getATree([]int{1, 2, 3, 4, 5})

	newTree.Mirror()
	fmt.Println("\t"+colorYellow+"Printing Original Tree", colorReset)
	originalTree.PrintTree()
	mirroTree := newTree
	fmt.Println("\t"+colorYellow+"Printing the Mirrored Tree", colorReset)
	mirroTree.PrintTree()
	fmt.Println(colorCyan, "\tAre mirrors: (OriginalTree & MirroredTree)\t", originalTree.AreMirrors(mirroTree), colorReset)
	fmt.Println(colorCyan, "\tAre mirrors: (OriginalTree & OriginalTree)\t", originalTree.AreMirrors(originalTree), colorReset)
	printDottedLine()
}

// The function "twentySix" creates a binary search tree, searches for specific nodes,
// prints the tree, and finds the lowest common ancestor of two nodes.
func twentySix() {
	newTree := getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	node1 := newTree.SearchElement(6)
	node2 := newTree.SearchElement(7)
	node3 := newTree.SearchElement(10)
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	fmt.Println(colorCyan, "\tLCA of node1 & node2 (Same Subtree):\t", newTree.LCA(node1, node2), colorReset)
	fmt.Println(colorCyan, "\tLCA of node1 & node3 (Different Subtree):\t", newTree.LCA(node1, node3), colorReset)
	printDottedLine()
}

// TODO Doesn't work as expected, will fix this
func twentySeven() {
	newTree := getATree([]int{})
	fmt.Println("\t"+colorYellow+"Printing the Tree before building", colorReset)
	newTree.PrintTree()
	fmt.Println("\t"+colorYellow+"Building the Tree", colorReset)
	inOrder := []int{4, 2, 5, 1, 6, 3}
	preOrder := []int{1, 2, 4, 5, 3, 6}
	newTree.BuildBinaryTree(preOrder, inOrder, 0, 0)
	fmt.Println("\t"+colorYellow+"Printing the Tree after building", colorReset)
	newTree.PrintTree()
}

// The function "thirtyOne" creates a binary tree, prints the tree, calculates the vertical sum of the
// tree, and prints the sum.
func thirtyOne() {
	newTree := getATree([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	hashTable := make(map[int]int)
	fmt.Println(colorCyan, "\tGetting the Vertical Sum:\t")
	newTree.VerticalSumInBinaryTree(0, hashTable)
	for _, value := range hashTable {
		fmt.Printf("\tSum = %d\n", value)
	}
	fmt.Println(colorReset)
	printDottedLine()
}

// The function calculates the depth of a generic tree.
func thirtyNine() {
	parentArray := []int{-1, 0, 1, 6, 6, 0, 0, 2, 7}
	fmt.Println("\t"+colorYellow+"Calculating depth of generic tree", colorReset)
	fmt.Println(colorCyan, "\tGeneric depth:\t", trees.GenericTreeDepth(parentArray), colorReset)
	printDottedLine()
}

// The function `fortyTwo` creates three trees and checks if they are isomorphic.
func fortyTwo() {
	tree1 := getATree([]int{1, 2, 3, 4, 5})
	fmt.Println("\t"+colorYellow+"Printing the Tree (Tree1)", colorReset)
	tree1.PrintTree()
	tree2 := getATree([]int{6, 7, 8, 9, 10})
	fmt.Println("\t"+colorYellow+"Printing the Tree (Tree2)", colorReset)
	tree2.PrintTree()
	fmt.Println("\t"+colorYellow+"Printing the Tree (Tree3)", colorReset)
	tree3 := getATree([]int{10, 11, 12, 13, 14, 14})
	tree3.PrintTree()

	fmt.Println("\t"+colorCyan+"To check if Tree1 & Tree2 are Isomorphic:\t", tree1.IsIsomorphic(tree2))
	fmt.Println("\t"+colorCyan+"To check if Tree1 & Tree3 are Isomorphic:\t", tree1.IsIsomorphic(tree3))
	fmt.Println(colorReset)
	printDottedLine()
}

func fortyThree() {
	tree1 := getATree([]int{1, 2, 3, 4, 5})
	fmt.Println("\t"+colorYellow+"Printing the Tree (Tree1)", colorReset)
	tree1.PrintTree()
	tree2 := getATree([]int{6, 7, 8, 9, 10})
	fmt.Println("\t"+colorYellow+"Printing the Tree (Tree2)", colorReset)
	tree2.PrintTree()
	fmt.Println("\t"+colorYellow+"Printing the Tree (Tree3)", colorReset)
	tree3 := getATree([]int{10, 11, 12, 13, 14, 14})
	tree3.PrintTree()

	fmt.Println("\t"+colorCyan+"To check if Tree1 & Tree2 are QuasiIsomorphic:\t", tree1.IsQuasiIsomorphic(tree2))
	fmt.Println("\t"+colorCyan+"To check if Tree1 & Tree3 are QuasiIsomorphic:\t", tree1.IsQuasiIsomorphic(tree3))
	fmt.Println("\t"+colorCyan+"To check if Custom Trees added via code are QuasiIsomorphic:\t", trees.CheckQuasiIsomorphism())
	fmt.Println(colorReset)
	printDottedLine()
}

// The function "executeExercises" executes a series of tree-related exercises and prints the results.
func executeExercises() {
	fmt.Println("BINARY TREE EXERCISES")
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
	fmt.Println("Problem 18: Give an algorithm for finding the diameter of the binary tree. The diameter of a tree (sometimes called the width) is the number of nodes on the longest path between two leaves in the tree.")
	eighteen()
	fmt.Println("Problem 19: Give an algorithm for finding the level that has the maximum sum in the binary tree.")
	nineteen()
	fmt.Println("Problem 20: Given a binary tree, print out all its root-to-leaf paths.")
	twenty()
	fmt.Println("Problem 21: Give an algorithm for checking the existence of path with given sum. That means, given a sum, check whether there exists a path from root to any of the nodes.")
	twentyOne()
	fmt.Println("Problem 22: Give an algorithm for finding the sum of all elements in binary tree. \nProblem 23: Can we solve Problem-22 without recursion?")
	twentyTwoAndTwentyThree()
	fmt.Println("Problem 24: Give an algorithm for converting a tree to its mirror. Mirror of a tree is another tree with left and right children of all non-leaf nodes interchanged. The trees below are mirrors to each other.")
	twentyFour()
	fmt.Println("Problem 25: Given two trees, give an algorithm for checking whether they are mirrors of each other.")
	twentyFive()
	fmt.Println("Problem 26: Give an algorithm for finding LCA (Least Common Ancestor) of two nodes in a Binary Tree.")
	twentySix()
	fmt.Println("Problem 27: Give an algorithm for constructing binary tree from given Inorder and Preorder traversals.")
	twentySeven()
	// Randomly Picking Up Problems
	fmt.Println("Problem 31: Give an algorithm for finding the vertical sum of a binary tree.")
	thirtyOne()
	fmt.Println("Problem 39: Given a parent array P, where P[i] indicates the parent of ith node in the tree (assume parent of root node is indicated with –1). Give an algorithm for finding the height or depth of the tree.")
	thirtyNine()
	fmt.Println("Problem 42: Given two trees how do we check whether the trees are isomorphic to each other or not?")
	fortyTwo()
	fmt.Println("Problem 43: Given two trees how do we check whether they are quasi-isomorphic to each other or not?")
	fortyThree()
}

func main() {
	executeExercises()
}
