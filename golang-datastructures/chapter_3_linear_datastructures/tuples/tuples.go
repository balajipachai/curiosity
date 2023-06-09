package tuples

// The Square function takes a float64 input and returns its square.
func Square(x float64) float64 {
	return x * x
}

// The function calculates the product of the input value and its cube.
func XToThePower4(x, cube float64) float64 {
	return x * cube
}

// The function Cube calculates the cube and square of a given float64 input.
func Cube(x float64) (cube float64) {
	cube = x * Square(x)
	return cube
}
