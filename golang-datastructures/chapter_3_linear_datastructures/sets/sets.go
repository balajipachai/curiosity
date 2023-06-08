package sets

import "fmt"

// The Set type is a struct that contains a map of integers to booleans.
// @property integerMap - integerMap is a property of the Set struct which is a map that stores integer
// keys and boolean values. It is used to keep track of the elements present in the set. If an integer
// key is present in the map, it means that the corresponding element is present in the set. If the
// boolean
type Set struct {
	integerMap map[int]bool
}

// It is a method of the `Set` struct that initializes the `integerMap` property of the `Set` struct as
// an empty map of integers to booleans. It is used to create a new instance of the `Set` struct with
// an empty `integerMap`.
func (set *Set) New() {
	set.integerMap = make(map[int]bool)
}

// It is a method of the `Set` struct that checks if
// an element is present in the set. It takes an integer `element` as input and returns a boolean
// value. It checks if the `element` is present in the `integerMap` property of the `Set` struct by
// using the `element` as the key to access the corresponding value in the map. If the key exists in
// the map, it means that the element is present in the set and the method returns `true`. Otherwise,
// it returns `false`.
func (set *Set) ContainsElement(element int) bool {
	_, exists := set.integerMap[element]
	return exists
}

// It is a method of the `Set` struct that adds an element to
// the set. It takes an integer `element` as input and checks if the element is already present in the
// set by calling the `ContainsElement` method. If the element is not present in the set, it adds the
// element to the `integerMap` property of the `Set` struct by using the `element` as the key and
// setting the corresponding value to `true`. If the element is already present in the set, it prints a
// message saying "Cannot add duplicate element".
func (set *Set) AddElement(element int) {
	if !set.ContainsElement(element) {
		set.integerMap[element] = true
	} else {
		fmt.Println("\tCannot add duplicate element")
	}
}

// It is a method of the `Set` struct that deletes an element
// from the set. It takes an integer `element` as input and checks if the element is present in the set
// by calling the `ContainsElement` method. If the element is present in the set, it deletes the
// element from the `integerMap` property of the `Set` struct by using the `delete` function and
// passing the `element` as the key. If the element is not present in the set, it prints a message
// saying "Cannot delete non-existing element".
func (set *Set) DeleteElement(element int) {
	if set.ContainsElement(element) {
		delete(set.integerMap, element)
	} else {
		fmt.Println("\tCannot delete non-existing element")
	}
}

// The `Intersection` method is a method of the `Set` struct that takes another `Set` as input and
// returns a new `Set` that contains the intersection of the two sets. It creates a new `Set` called
// `intersection` and iterates over the keys of both sets using nested loops. If a key is present in
// both sets, it adds the key to the `integerMap` property of the `intersection` set with a value of
// `true`. Finally, it returns the `intersection` set.
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

// `func (set *Set) Union(anotherSet *Set) *Set` is a method of the `Set` struct that takes another
// `Set` as input and returns a new `Set` that contains the union of the two sets. It creates a new
// `Set` called `union` and iterates over the keys of both sets using two separate loops. It adds all
// the keys of the first set to the `integerMap` property of the `union` set with a value of `true`.
// Then, it adds all the keys of the second set to the `integerMap` property of the `union` set with a
// value of `true`. Finally, it returns the `union` set.
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
