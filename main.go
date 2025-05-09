// Package mathmod provides basic math utilities.
package mathmod

import "golang.org/x/exp/constraints"

// Number is a type constraint for numeric types.
type Number interface {
    constraints.Integer | constraints.Float
}

// Add returns the sum of a and b.
//
// See also: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
    return a + b
}