package problemAndsolutions

import (
	"fmt"

	"example.com/singlylist"
)

// The function `FindNthNodeFromEnd` finds the nth node from the end of a singly linked list.
func FindNthNodeFromEnd(list *singlylist.SinglyLinkedList, n int) *singlylist.ListNode {
	listLength := list.ListLength()
	/**
	* nth node from the end implies `listLength - n`th node from the beginning
	 */
	if n > listLength {
		return nil
	}
	nthNodePosition := listLength - n + 1
	pos := 0
	var nthNode *singlylist.ListNode
	for node := list.HeadNode; node != nil; node = node.NextNode {
		pos++
		if pos == nthNodePosition {
			nthNode = node
			break
		}
	}
	return nthNode
}

// The function finds the nth node from the end of a singly linked list using a hash table.
func FindNthNodeFromEndUsingHashTable(list *singlylist.SinglyLinkedList, n int) *singlylist.ListNode {
	listLength := list.ListLength()
	if n > listLength {
		return nil
	}
	nthNodePosition := listLength - n + 1
	nthNode := singlylist.HashTable[nthNodePosition]
	return nthNode
}

// The function `FindNthNodeInOneScan` finds the nth node in a singly linked list in one scan.
func FindNthNodeInOneScan(list *singlylist.SinglyLinkedList, n int) *singlylist.ListNode {

	listLength := list.ListLength()
	/**
	* nth node from the end implies `listLength - n`th node from the beginning
	 */
	if n > listLength {
		return nil
	}

	pTemp := list.HeadNode
	pNthNode := list.HeadNode

	count := 0
	for node := pTemp; node != nil; node = node.NextNode {
		count++
		pTemp = node
		if count > n {
			pNthNode = pNthNode.NextNode
		}
	}
	return pNthNode
}

// TODO How do we find the end of the linked list?
// Yet to find the answer for this
func FindLoopInListBruteForceApproach(list *singlylist.SinglyLinkedList) bool {
	headNode := list.HeadNode
	currentNode := headNode
	nodeTraversalCount := make(map[*singlylist.ListNode]int)
	isLoop := false

	for node := currentNode; node != nil; node = node.NextNode {
		for innerNode := node.NextNode; innerNode != nil; innerNode = innerNode.NextNode {
			nodeTraversalCount[innerNode]++
			if nodeTraversalCount[innerNode] > 1 {
				fmt.Println("inside first if", innerNode)
				if node == innerNode {
					fmt.Println("inside node == innerNode")
					isLoop = true
					break
				}
			}
		}
	}
	return isLoop
}
