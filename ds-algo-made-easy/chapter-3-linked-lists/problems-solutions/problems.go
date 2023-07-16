package problemAndsolutions

import (
	"fmt"

	"example.com/singlylist"
)

const (
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
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

/*
The function `FindNthNodeInOneScan` finds the nth node in a singly linked list in one scan.
Algorithm:
1. Use two pointers pNthNode and pTemp.
2. Initially, both point to head node of the list.
3. pNthNode starts moving only after pTemp has made n moves.
4. From there both move forward until pTemp reaches the end of the list.
5. As a result pNthNode
points to nth node from the end of the linked list.
*/
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

// The function "FindLoopInListBruteForceApproach" checks if a given singly linked list contains a loop
// using a brute force approach.
func FindLoopInListBruteForceApproach(list *singlylist.SinglyLinkedList) bool {
	currentNode := list.HeadNode
	nodeVisited := make(map[*singlylist.ListNode]bool)

	for currentNode != nil {
		if nodeVisited[currentNode] {
			return true
		}
		nodeVisited[currentNode] = true
		currentNode = currentNode.NextNode
	}
	return false
}

/*
The function uses Floyd's algorithm to determine if a loop exists in a singly linked list.
Algorithm:
1. Use 2 pointers
2. Hare & Tortoise, Hare moves 2 nodes at a time
3. Tortoise moves 1 node at a time
4. When Tortoise == Hare, it implies there is a loop in the list
5. After entire list traversal, if nowhere, Tortoise != Hare then there is no loop in the list
*/
func FindLoopInListUsingFlyodAlgorithm(list *singlylist.SinglyLinkedList) bool {
	tortoise := list.HeadNode
	hare := list.HeadNode

	for hare != nil && hare.NextNode != nil {
		tortoise = tortoise.NextNode
		hare = hare.NextNode.NextNode
		if tortoise == hare {
			return true
		}
	}
	return false
}

// The function determines whether a given singly linked list is a snake or a snail.
func GivenListIsASnakeOrSnail(list *singlylist.SinglyLinkedList) string {
	isSnail := FindLoopInListUsingFlyodAlgorithm(list)
	if isSnail {
		return "Snail"
	}
	return "Snake"
}

// The function `PrintListHavingLoop` prints the elements of a singly linked list until it encounters a
// node that has been visited before, indicating the presence of a loop in the list.
func PrintListHavingLoop(list *singlylist.SinglyLinkedList) {
	currentNode := list.HeadNode
	nodeVisited := make(map[*singlylist.ListNode]bool)

	for currentNode != nil {
		fmt.Printf(colorMagenta+"\t%d\t", currentNode.Data)
		if nodeVisited[currentNode] {
			break
		}
		nodeVisited[currentNode] = true
		currentNode = currentNode.NextNode
	}
	fmt.Printf("\n")
	fmt.Println(colorReset)
}

// The function FindStartingNodeOfLoopInLinkedList finds the starting node of a loop in a linked list
// using the Floyd's cycle-finding algorithm.
func FindStartingNodeOfLoopInLinkedList(list *singlylist.SinglyLinkedList) *singlylist.ListNode {
	tortoise := list.HeadNode
	hare := list.HeadNode
	var loopStartNode *singlylist.ListNode
	for hare != nil && hare.NextNode != nil {
		tortoise = tortoise.NextNode
		hare = hare.NextNode.NextNode

		if hare == tortoise {
			// This implies loop Exists
			/*
				After finding the loop in the linked list, we initialize the slowPtr to the head of the linked list. From that point onwards both slowPtr and fastPtr move only one node at a time. The point at which they meet is the start of the loop. Generally we use this method for removing the loops.
			*/
			tortoise := list.HeadNode
			for tortoise != hare {
				tortoise = tortoise.NextNode
				hare = hare.NextNode
			}
			loopStartNode = tortoise
			return loopStartNode
		}
	}
	return loopStartNode
}

// In the Floyd cycle finding algorithm, does it work if we use steps 2 and 3 instead of 1 and 2
// The function "FindLoopBySteps2And3" checks if a singly linked list contains a loop by using two
// pointers, one moving one step at a time and the other moving two steps at a time, and returns
// true if they meet.
func FindLoopBySteps2And3(list *singlylist.SinglyLinkedList) bool {
	tortoise := list.HeadNode
	hare := list.HeadNode.NextNode.NextNode

	for hare != nil && hare.NextNode != nil {
		tortoise = tortoise.NextNode
		hare = hare.NextNode.NextNode

		if hare == tortoise {
			return true
		}
	}
	return false
}

// Check whether the given linked list is NULL-terminated. If there is a cycle, find the length of the loop.
// The function `LinkedListLoopLength` calculates the length of a loop in a singly linked list, if a
// loop exists.
func LinkedListLoopLength(list *singlylist.SinglyLinkedList) int {
	tortoise := list.HeadNode
	hare := list.HeadNode.NextNode.NextNode
	length := 0
	for hare != nil && hare.NextNode != nil {
		tortoise = tortoise.NextNode
		hare = hare.NextNode.NextNode

		if hare == tortoise {
			fmt.Println("Loop exists")
			// Loop Exists
			hare = hare.NextNode
			for hare != tortoise {
				hare = hare.NextNode
				length++
			}
			return length
		}
	}
	return length
}
