# Questions

## 1. Where can you use double linked list? Please provide an example
- Couldn't think of one at the moment, will touch down on this later.

## 2. Which method on linked list can be used for printing out node values?
- The `IterateList` method on linked list can be used for printing out node value.

## 3. Which queue was shown with channels from the Go language?
- Synchronized Queue

## 4. Write a method that returns multiple values. What data structure can be used for returning multiple values?

```go
package main

import "fmt"

func squareAndCubeOf(number float64) (square, cube float64) {
    square := number * number
    cube = square * number
    return square, cube
}

func main() {
    square, cube := squareAndCubeOf(10)
    fmt.Println("Square of 10:\n\t", square)
    fmt.Println("Cube of 10:\n\t", cube)
}
```
`Tuple` datastructure is used for returning multiple values.

## 5. Can set have duplicate elements?

- No

## 6. Write a code sample showing the union and intersection of two sets.

```go

package main

type Set struct {
	integerMap map[int]bool
}

func (set *Set) Intersection(anotherSet *Set) *Set {
	intersection := &Set{}
	intersection.New()
	for kS1 := range set.integerMap {
		for kS2 := range anotherSet.integerMap {
			if kS1 == kS2 {
				intersection.integerMap[kS2] = true
			}
		}
	}
	return intersection
}

func (set *Set) Union(anotherSet *Set) *Set {
	union := &Set{}
	union.New()
	for kS1, _ := range set.integerMap {
		union.integerMap[kS1] = true
	}

	for kS2, _ := range anotherSet.integerMap {
		union.integerMap[kS2] = true
	}
	return union
}

func main() {
    setA := Set{}
	setB := Set{}

    setA.AddElement(1)
	setA.AddElement(2)
	setA.AddElement(3)
	setA.AddElement(4)

    setB.AddElement(6)
	setB.AddElement(7)
	setB.AddElement(8)
	setB.AddElement(4)

    fmt.Println("\nFinds the intersection of SetA & SetB")
	fmt.Println("\tSetA Intersection SetB: ", setA.Intersection(&setB))

    fmt.Println("\nFinds the union of SetA & SetB")
	fmt.Println("\tSetA Union SetB: ", setA.Union(&setB))
}
```

## 7. In a linked list, which method is used to find the node between two values?

```go

package main

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) NodeBetweenValues(firstProperty, secondProperty int) *Node {
	nodeBetweenValues := &Node{}
	for node := linkedList.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty && node.nextNode.property == secondProperty {
				nodeBetweenValues = node
				break
			}
		}
	}
	return nodeBetweenValues
}
```

## 8. We have elements that are not repeated and unique. What is the correct data structure that represents the collection?
- SET

## 9. In Go, how do you generate a random integer between the values 3 and 5?
- Using the `math/rand` package

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 3 and 5
	randomNumber := rand.Intn(3) + 3
	fmt.Println(randomNumber)
}
```

## 10. Which method is called to check if an element of value 5 exists in the Set?

```go
// ContainsElement is called to check if an element exists in the Set, where the value can be anything
func (set *Set) ContainsElement(element int) bool {
	_, exists := set.integerMap[element]
	return exists
}
```