package sets

const (
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
)

type Set struct {
	elements []int
}

// The `MakeSet` function is initializing a set with a given size. It creates a slice of integers with
// the specified size and then appends the elements from 0 to size-1 to the slice. This function is
// used to create a set with a specified number of elements, where each element is initially its own
// root.
// The multiplier is either 1 or -1
func (set *Set) MakeSet(size, multiplier int) {
	set.elements = make([]int, size)
	for i := 0; i < size; i++ {
		set.elements = append(set.elements, i*multiplier)
	}
}

// The `SlowFind` function is used to find the root of an element in the set. It takes an integer `x`
// as input and returns the root of `x` if it exists in the set, otherwise it returns -1.
func (set *Set) SlowFind(x int) int {
	size := len(set.elements)
	if !(x >= 0 && x < size) {
		return -1
	}
	if set.elements[x] == x {
		return x
	} else {
		return set.SlowFind(set.elements[x])
	}
}

// The `QuickUnion` function is used to perform the union operation on two elements in the set. It
// takes two integers `root1` and `root2` as input, which represent the roots of the elements to be
// unioned.
func (set *Set) QuickUnion(root1, root2 int) {
	size := len(set.elements)
	if set.SlowFind(root1) == set.SlowFind(root2) {
		return
	}
	if !((root1 >= 0 && root1 < size) && (root2 >= 0 && root2 < size)) {
		return
	}
	set.elements[root1] = root2
}

// The `QuickFind` function is used to find the root of an element in the set using the quick find
// algorithm. It takes an integer `x` as input and returns the root of `x` if it exists in the set,
// otherwise it returns -1.
func (set *Set) QuickFind(x int) int {
	size := len(set.elements)
	if !(x >= 0 && x < size) {
		return -1
	}
	if set.elements[x] == -1 {
		return x
	} else {
		return set.QuickFind(set.elements[x])
	}
}

// The `UnionBySize` function is used to perform the union operation on two elements in the set using
// the union by size algorithm. It takes two integers `root1` and `root2` as input, which represent the
// roots of the elements to be unioned.
func (set *Set) UnionBySize(root1, root2 int) {
	size := len(set.elements)
	if set.QuickFind(root1) == set.QuickFind(root2) {
		return
	}

	if !((root1 >= 0 && root1 < size) && (root2 >= 0 && root2 < size)) {
		return
	}

	set.elements[root2] += set.elements[root1]
	set.elements[root1] = root2
}

func (set *Set) PathCompressedFind(x int) int {
	size := len(set.elements)
	if !(x >= 0 && x < size) {
		return -1
	}
	if set.elements[x] <= 0 {
		return x
	} else {
		set.elements[x] = set.PathCompressedFind(set.elements[x])
		return set.elements[x]
	}
}
