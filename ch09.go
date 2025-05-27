package ch09

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds two numerical values together and returns the result
//
// For more information see [Math is fun]
//
// [Math is fun]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
