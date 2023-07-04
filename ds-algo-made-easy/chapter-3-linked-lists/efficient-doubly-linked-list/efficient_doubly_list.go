package efficientdoublylist

type EDLListNode struct {
	data            int
	nextAndPrevNode *EDLListNode
}

type EfficientDoublyLinkedList struct {
	headNode *EDLListNode
}

func (efficientDoublyLinkedList *EfficientDoublyLinkedList) ListLength() int {

}
