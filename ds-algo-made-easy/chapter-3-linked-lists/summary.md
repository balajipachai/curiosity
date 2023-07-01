# Basic Operations on a list

- Traversing
- Inserting
- Deleting
---
# Singly Linked List

## Insert

1. To insert a node in the beginning

    - [Update the next pointer of the new node, to point to the current head](singlylist/singly_list.go#L61)
    - [Update the head pointer to point to the new node](singlylist/singly_list.go#L62)

2. To insert a node in the end

    - New node's next pointer points to nil
    - Last nodes next pointer points to new node

3. To insert a node in the middle

    - Let us say we want to add a node at position 4
    - Traverse the list and stop at position 3
    - The node at position 3, becomes our `positionNode`
    - if `positionNode.nextNode != nil` this means we are inserting in the middle i.e. the `positionNode` is not the  last node
    
        - [newNode's next pointer points to positionNode's next pointer](singlylist/singly_list.go#L74)
        - [positionNodes's next pointer points to new node](singlylist/singly_list.go#L75)

    -  else we are inserting at the last node

        - [positionNodes's next pointer points to new node](singlylist/singly_list.go#L78)

---

## Delete

1. To delete a node from the beginning

    - [Move the headNode pointer to the next node](singlylist/singly_list.go#L106)

2. To delete the last node / To delete the node from the middle

    - Let us say we want to add a node at position 4
    - Traverse the list and stop at position 3
    - The node at position 3, becomes our `positionNode`
    - if `positionNode.nextNode.nextNode != nil` this means we are deleting from the middle i.e. the `positionNode` is not the  last node

        - [positionNodes's next pointer points to positionNodes.nextNode.nextNode](singlylist/singly_list.go#L118)

    - else we are deleting the last node

        - [positionNode's next pointer points to nil](singlylist/singly_list.go#L121)
---


# Doubly Linked List

## Insert

1. To insert a node in the beginning

    - [Update the next pointer of the new node, to point to the current head](doublylist/doubly_list.go#L84)
    - [Update the previous pointer of the new node to point the current head](doublylist/doubly_list.go#L85)
    - [Update the head pointer to point to the new node](doublylist/doubly_list.go#L86)

2. To insert a node in the end

    - [New node's next pointer points to nil](doublylist/doubly_list.go#L108)
    - [New node's previous pointer points to positionNode](doublylist/doubly_list.go#L107)

3. To insert a node in the middle

    - Let us say we want to add a node at position 4
    - Traverse the list and stop at position 3
    - The node at position 3, becomes our `positionNode`
    - The node to insert becomes `nodeToInsert = newNode`
    - if `positionNode.nextNode != nil` this means we are inserting in the middle i.e. the `positionNode` is not the  last node
    
        - [Node to insert's next pointer points to positionNode's next pointer](doublylist/doubly_list.go#L100)
        - [Position node's next node's previous pointer pointer points to node to insert](doublylist/doubly_list.go#L101)
        - [Node to insert's previous node points to position node](doublylist/doubly_list.go#L102)
        - [Position node's next node points to the node to insert i.e. newNode](doublylist/doubly_list.go#L103)

    -  else we are inserting at the last node

        - [To insert a node in the end](summary.md#L68)

---

## Delete

1. To delete a node from the beginning

    - [Move the headNode pointer to the next node](doublylist/doubly_list.go#L124)
    - [Make the head node's next node's previous node pointing to nil](doublylist/doubly_list.go#L125)

2. To delete the last node / To delete the node from the middle

    - Let us say we want to add a node at position 4
    - Traverse the list and stop at position 3
    - The node at position 3, becomes our `positionNode`
    - The node to delete becomes: `nodeToDelete := positionNode.nextNode`
    - if `positionNode.nextNode.nextNode != nil` this means we are deleting from the middle i.e. the `positionNode` is not the  last node

        - [positionNodes's next pointer points to nodeToDelete.nextNode](doublylist/doubly_list.go#L138)
        - [nodeToDelete's next node's previous node pointer points to the position node](doublylist/doubly_list.go#L139)

    - else we are deleting the last node

        - [positionNode's next pointer points to nil](doublylist/doubly_list.go#L143)
---
