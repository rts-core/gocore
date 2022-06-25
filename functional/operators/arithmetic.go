package operators

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

// Add simple addition operation
func Add[T Number](a, b T) T {
	return a + b
}

// Mul simple multiplication operation
func Mul[T Number](a, b T) T {
	return a * b
}

// Div simple division operation (a/b)
func Div[T Number](a, b T) T {
	return a / b
}

// Sub simple subtraction operation (a-b)
func Sub[T Number](a, b T) T {
	return a - b
}

// Xor simple bitwise xor operation
func Xor[T constraints.Integer](a, b T) T {
	return a ^ b
}
