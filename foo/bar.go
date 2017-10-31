package foo

var sideEffect int

// Add sums two numbers - simple!
func Add(x, y int) int {
	sum := x + y
	sideEffect += sum
	return sum
}
