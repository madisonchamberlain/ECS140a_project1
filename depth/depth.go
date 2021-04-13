package depth

import (
	"hw1/expr"
)

// Depth should return the maximum number of AST nodes between the root of the
// tree and any leaf (literal or variable) in the tree.
func Depth(e expr.Expr) uint {
	var depth uint = 1
	// find out what type e is...

	switch e.(type) {
	// if e is a Literal its depth is 1
	case expr.Literal:
		return depth
	// if e is a variable its depth is 1
	case expr.Var:
		return depth
	// if e is a unary its depth is 1 + Depth(Expression)
	case expr.Unary:
		e := e.(expr.Unary)
		// recurse on the expression portion
		return depth + Depth(e.X)
	case expr.Binary:
		e := e.(expr.Binary)
		// recurse on the both expression portions
		x := Depth(e.X)
		y := Depth(e.Y)
		// add the larger side to the current depth count 
		if x > y{
			return depth + x
		} else {
			return depth + y
		}
	// if its none of those types then panic
	default:
		panic("unsupported expression")
	}
}
