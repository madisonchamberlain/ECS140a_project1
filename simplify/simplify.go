package simplify

import (
	"hw1/expr"
)

// Simplify should return the simplified expresion
func Simplify(e expr.Expr, env expr.Env) expr.Expr {
	// determine type for e
	switch e.(type) {

	case expr.Literal:
		// if e is a literal; simply return it
		return e

	case expr.Var:
		// if e is a variable in the map; return the val from the map
		if _, ok := env[e.(expr.Var)]; ok {
			return expr.Literal(e.Eval(env))
		}

		// if e is not in the map; return e
		return e

	case expr.Unary:
		e := e.(expr.Unary)
		// if the expression is a literal; simplify further
		if _, ok := Simplify(e.X, env).(expr.Literal); ok {
			//return the value with the sign taken into account
			return expr.Literal(e.Eval(env))
		} else {
			// if its a variable, the negative/positive must stay in char form
			return e
		}
	case expr.Binary:
		// type case e to Binary
		e := e.(expr.Binary)

		// recurse on the both expression portions
		x := Simplify(e.X, env)
		y := Simplify(e.Y, env)

		// set both sides to their simplified values
		e.X = x
		e.Y = y

		//if neither sides are literals then return original
		if _, ok := x.(expr.Literal); !ok {
			if _, ok := y.(expr.Literal); !ok {
				return e
			}
		}

		// if x literal and y not; simplify x, keep y
		if _, ok := x.(expr.Literal); ok {
			if _, ok := y.(expr.Literal); !ok {
				if x.Eval(env) == 0 {
					// return Y if adding 1
					if e.Op == '+' {
						return e.Y
					}
					if e.Op == '*' {
						// return 0 if multiply by 0
						return expr.Literal(0)
					}
				}
				// if x evaluates to 1 and you are multiplying or dividing, return y
				if x.Eval(env) == 1 {
					if e.Op == '*' {
						return e.Y
					}
				} else {
					return e
				}
			}
		}

		// if x literal and y not; simplify x, keep y
		if _, ok := x.(expr.Literal); !ok {
			if _, ok := y.(expr.Literal); ok {
				if y.Eval(env) == 0 {
					// return x if adding 0
					if e.Op == '+' {
						return e.X
					}
					// return 0 if multiply by 0
					if e.Op == '*' {
						return expr.Literal(0)
					}
				}
				// if y is one and op is multiply return x
				if y.Eval(env) == 1 {
					if e.Op == '*' {
						return e.X
					}
				} else {
					return e
				}
			}
		}

		// if both sides are Literals then evaluate the expression
		return expr.Literal(e.Eval(env))

	default:
		panic("unsupported expression")
	}
}
