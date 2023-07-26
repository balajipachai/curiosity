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
	fmt.Println("\t" + colorYellow + "Print height of the tree (BOTH LEFT & RIGHT AT SAME LEVEL) [Non-Recursive]" + colorReset)
	fmt.Println(colorCyan, "\tTree height is = ", binaryTree.HeightOfBinaryTreeNonRecursive(), colorReset)
	printDottedLine()

	// Second Tree
	newTree := getATree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	newTree.LevelOrderTraversal()
	fmt.Println("\t"+colorYellow+"Printing the Tree", colorReset)
	newTree.PrintTree()
	fmt.Println("\t" + colorYellow + "Print height of the tree (LEFT SUBTREE HEIGHT > RIGHT SUBTREE HEIGHT) [Recursive]" + colorReset)
	fmt.Println(colorCyan, "\tTree height is = ", newTree.HeithtOfBinaryTree(), colorReset)
	fmt.Println("\t" + colorYellow + "Print height of the tree (LEFT SUBTREE HEIGHT > RIGHT SUBTREE HEIGHT) [Non-Recursive]" + colorReset)
	fmt.Println(colorCyan, "\tTree height is = ", newTree.HeightOfBinaryTreeNonRecursive(), colorReset)
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
}

func main() {
	executeExercises()
}
