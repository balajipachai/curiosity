package main

import (
	"container/heap"
	"container/list"
	"fmt"
)

/**
* List is a sequential datastructure.
* In Go, lists can be accessed from the container/list package
* Different functions available on the list object are:
* func (*list.List).Back() *list.Element
* func (*list.List).Front() *list.Element
* func (*list.List).Init() *list.List
* func (*list.List).InsertAfter(v any, mark *list.Element) *list.Element
* func (*list.List).InsertBefore(v any, mark *list.Element) *list.Element
* func (*list.List).Len() int
* func (*list.List).MoveAfter(e *list.Element, mark *list.Element)
* func (*list.List).MoveBefore(e *list.Element, mark *list.Element)
* func (*list.List).MoveToBack(e *list.Element)
* func (*list.List).MoveToFront(e *list.Element)
* func (*list.List).PushBack(v any) *list.Element
* func (*list.List).PushBackList(other *list.List)
* func (*list.List).PushFront(v any) *list.Element
* func (*list.List).PushFrontList(other *list.List)
* func (*list.List).Remove(e *list.Element) any
 */

// The function creates a linked list of integers and prints its elements.
func createListAndPrint() {
	intList := list.New().Init()
	intList.PushBack(5)
	intList.PushBack(8)
	intList.PushBack(16)

	fmt.Println("Printing integer lists")
	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}

/*
The function "tuples" returns the square and cube of a given number and an error if any.
A tuple is a finite sorted list of elements. It is a datastructure that groups data.
*/
func tuples(num int) (square, cube int, err error) {
	square = num * num
	cube = square * num
	return
}

/**
* An integer heap
* Package heap provides heap operations for any type that implements heap Interface.
**** A heap is a tree with the property that each node is the minimum-valued node in its subtree.
 */
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() any {
	oldSlice := *h
	n := len(oldSlice)
	x := oldSlice[n-1]
	*h = oldSlice[0 : n-1]
	return x
}

/*
Composite Pattern: Answer Question 1
*/
// The below code defines an interface named IComposite with a single method named perform().
// @property perform - perform is a method signature that belongs to the IComposite interface. Any
// struct that implements this interface must also implement the perform method. The purpose of the
// perform method is specific to the implementation of the struct that implements the IComposite
// interface.
type IComposite interface {
	perform()
}

// The below code defines a struct type called Leaflet with a single field named "name" of type string.
// @property {string} name - The `name` property is a string type field that represents the name of a
// `Leaflet` object.
type Leaflet struct {
	name string
}

// The `perform()` function is a method of the `Leaflet` struct that implements the `IComposite`
// interface. It simply prints out the name of the leaflet object.
func (leaf *Leaflet) perform() {
	fmt.Println("Leaflet " + leaf.name)
}

// The type "Branch" represents a tree structure with leafs and branches, identified by a name.
// @property {[]Leaflet} leafs - A slice of Leaflet structs that belong to this branch.
// @property {string} name - The name property is a string that represents the name of the branch.
// @property {[]Branch} branches - `branches` is a property of the `Branch` struct which is a slice of
// `Branch` type. It represents the sub-branches of the current branch. Each sub-branch is also a
// `Branch` struct with its own set of properties.
type Branch struct {
	leafs    []Leaflet
	name     string
	branches []Branch
}

// The `perform()` function is a method of the `Branch` struct that implements the `IComposite`
// interface. It prints out the name of the branch, its leafs, and sub-branches. It then iterates over
// each leaf and sub-branch and calls their respective `perform()` methods recursively. This allows for
// a tree-like structure to be printed out, with each branch and leaf being displayed in a hierarchical
// manner.
func (branch *Branch) perform() {
	fmt.Println("Branch: " + branch.name)
	fmt.Println("leafs: ", branch.leafs)
	fmt.Println("branches: ", branch.branches)

	for _, leaf := range branch.leafs {
		leaf.perform()
	}

	for _, branch := range branch.branches {
		fmt.Println("inside branch.branches")
		fmt.Println("branch: ", branch)
		branch.perform()
	}
}

// The `add` function is a method of the `Branch` struct that takes a `Leaflet` object as input and
// adds it to the `leafs` slice of the `Branch` object. It does this by using the `append` function to
// add the `Leaflet` object to the end of the `leafs` slice.
func (branch *Branch) add(leaf Leaflet) {
	branch.leafs = append(branch.leafs, leaf)
}

// The `addBranch` method is a function of the `Branch` struct that takes a `Branch` object as input
// and adds it to the `branches` slice of the `Branch` object. It does this by using the `append`
// function to add the `Branch` object to the end of the `branches` slice.
func (branch *Branch) addBranch(newBranch Branch) {
	branch.branches = append(branch.branches, newBranch)
}

// The `getLeaflets()` method is a function of the `Branch` struct that returns a slice of `Leaflet`
// objects. It does this by simply returning the `leafs` property of the `Branch` object.
func (branch *Branch) getLeaflets() []Leaflet {
	return branch.leafs
}

// The main function demonstrates the usage of various Go language features such as tuples, integer
// heap, and composite pattern.
func main() {
	createListAndPrint()
	square, cube, err := tuples(5)
	if err != nil {
		fmt.Errorf("%w error occured", err)
	}
	fmt.Printf("Square and Cube of 5 is %d and %d\n", square, cube)
	fmt.Println("*************************Integer Heap******************************")
	h := &intHeap{10, 34, 5, 87}
	// Initialize heap
	heap.Init(h)
	heap.Push(h, 100)
	fmt.Printf("minimum of heap is %d\n", (*h)[0])
	fmt.Println("heap Len: ", h.Len())
	fmt.Println(h)
	fmt.Println("heap Less(0, 3) => : ", h.Less(0, 3))
	h.Swap(0, 3)
	fmt.Println(h)
	for h.Len() > 0 {
		fmt.Printf("%d\t", h.Pop())
	}
	fmt.Println("*************************Integer Heap******************************")
	fmt.Println()

	fmt.Println("*************************Composite Pattern******************************")
	branch := &Branch{name: "branch 1"}
	leaf1 := Leaflet{name: "leaf 1"}
	leaf2 := Leaflet{name: "leaf 2"}

	branch2 := Branch{name: "branch 2"}

	branch.add(leaf1)
	branch.add(leaf2)

	leaf3 := Leaflet{name: "leaf 3"}
	leaf4 := Leaflet{name: "leaf 4"}
	branch2.add(leaf3)
	branch2.add(leaf4)

	branch.addBranch(branch2)
	branch.perform()

	fmt.Println("*************************Composite Pattern******************************")

}

/**
Questions:
1. k-way merge algorithms?
2. Borrey-Moore algorithms?
3. Ukkonne algorithms?

Questions and exercises
1. Give an example where you can use a composite pattern.
Ans: Explained via code

2. For an array of 10 elements with a random set of integers, identify the maximum and minimum. Calculate the complexity of the algorithm.
Ans: O(n)

3. To manage the state of an object, which structural pattern is relevant?
Ans: Flyweight is used to manage the state of an object with high variation.

4. A window is sub-classed to add a scroll bar to make it a scrollable window. Which pattern is applied in this scenario?
Ans: Decorator Pattern.

5. Find the complexity of a binary tree search algorithm.
Ans: O(log n)

6. Identify the submatrices of 2x2 in a 3x3 matrix. What is the complexity of the algorithm that you have used?
Ans:

7. Explain with a scenario the difference between brute force and backtracking algorithms.
Ans: 
*/
