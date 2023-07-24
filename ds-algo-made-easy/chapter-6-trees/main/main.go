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

func problemFive() {
	binaryTree := &trees.BinaryTree{}
	fmt.Println("\t"+colorYellow+"Add 1, 2, 3, 4, 5, 6 & 7 into the Tree", colorReset)
	binaryTree.Insert(1)
	binaryTree.Insert(2)
	binaryTree.Insert(3)
	binaryTree.Insert(4)
	binaryTree.Insert(5)
	binaryTree.Insert(6)
	binaryTree.Insert(7)
	fmt.Println("\t"+colorYellow+"Printing the Tree:", colorReset)
	binaryTree.PrintTree()
	printDottedLine()
	fmt.Println("\t"+colorYellow+"Printing the LevelOrderTraversal", colorReset)
	binaryTree.LevelOrderTraversal()
	printDottedLine()
}

func treeTraversals() {
	binaryTree := &trees.BinaryTree{}
	fmt.Println("\t"+colorYellow+"Add 1, 2, 3, 4, 5, 6, 7, 8 & 9 into the Tree", colorReset)
	binaryTree.Insert(1)
	binaryTree.Insert(2)
	binaryTree.Insert(3)
	binaryTree.Insert(4)
	binaryTree.Insert(5)
	binaryTree.Insert(6)
	binaryTree.Insert(7)
	binaryTree.Insert(8)
	binaryTree.Insert(9)
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

func executeExercises() {
	fmt.Println("TREE EXERCISES")
	fmt.Println("Tree Traversals: PreOrder | InOrder | PostOrder")
	treeTraversals()
	fmt.Println("Problem 5: Give an algorithm for inserting an element into binary tree")
	problemFive()
}

func main() {
	executeExercises()
}
