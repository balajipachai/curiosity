package main

import (
	"fmt"

	"example.com/trees"
)

// The function executes various operations on a binary search tree, including inserting elements,
// traversing the tree in different orders, finding the minimum and maximum nodes, searching for nodes,
// and removing nodes.
func executeTreeOperations() {
	tree := &trees.BinarySearchTree{}
	tree.InsertElement(0)
	tree.InsertElement(1)
	tree.InsertElement(2)
	tree.InsertElement(3)
	tree.InsertElement(4)
	tree.InsertElement(5)
	tree.InsertElement(6)
	tree.InsertElement(7)
	tree.InsertElement(8)
	tree.InsertElement(9)
	tree.InsertElement(10)
	tree.String()
	fmt.Println("*****************IN-ORDER***********************")
	tree.InOrderTraverseTree(trees.PrintTraversedValue)
	fmt.Println("\n*****************PRE-ORDER***********************")
	tree.PreOrderTraverseTree(trees.PrintTraversedValue)
	fmt.Println("\n*****************POST-ORDER***********************")
	tree.PostOrderTraverseTree(trees.PrintTraversedValue)
	fmt.Println("\n*****************MIN-NODE***********************")
	fmt.Printf("%v", *tree.MinNode())
	fmt.Println("\n*****************MAX-NODE***********************")
	fmt.Printf("%v", *tree.MaxNode())
	fmt.Println("\n*****************SEARCH-NODE(5)***********************")
	fmt.Println(tree.SearchNode(5))
	fmt.Println("\n*****************SEARCH-NODE(15)***********************")
	fmt.Println(tree.SearchNode(15))
	fmt.Println("\n*****************REMOVE-NODE (15) => DOESN'T EXISTS***********************")
	tree.RemoveNode(15)
	tree.String()
	fmt.Println("\n*****************REMOVE-NODE (5) => IT EXISTS***********************")
	tree.RemoveNode(5)
	tree.String()
}

// The main function calls the executeTreeOperations function.
func main() {
	executeTreeOperations()
}
