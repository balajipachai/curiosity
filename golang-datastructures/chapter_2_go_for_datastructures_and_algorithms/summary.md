# Arrays

* Arrays are the most common datastructure that one come across in various programming languages.
* Arrays store homogeneous data types.
* In Go, len() returns the length of the array.
* Arrays can be traversed using a traditional for loop OR
* Arrays can be traversed by using the `range` keyword.
* While using `range`, both `index` & `value` can be returned.
* In Go, arrays have a fixed size.

# Slices

* Slices in Go can be thought of as dynamic arrays.
* Using len() & cap() functions we get the length and capacity of slices.
* `_` is the blank identifier, when you don't want to use some returned values make use of underscore.
* Difference between len() & cap()
*	len() tells you how many elements are currently stored in the slice.
*	cap() tells you the maximum number of elements the slice can hold without reallocation.

# Maps

* Maps in GO can be thought of as JSON objects or Hashtables that have a key and an associated value.