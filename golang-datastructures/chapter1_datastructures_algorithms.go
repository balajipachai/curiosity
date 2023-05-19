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

}

/**
Questions:
1. k-way merge algorithms?
*/
