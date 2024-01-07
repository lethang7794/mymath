// Package mymath provides utilities to make it easy do math
package mymath

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add returns sum of x and y.
//
// More information about it can be found at [Addition]
//
// [Addition]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](x, y T) T {
	return x + y
}
