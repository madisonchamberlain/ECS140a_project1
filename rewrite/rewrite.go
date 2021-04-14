package rewrite

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"hw1/expr"
	"hw1/simplify"
	//"log"
	//"strconv"
)

// check the package is expr

// rewriteCalls should modify the passed AST
func rewriteCalls(node ast.Node) {
	ast.Inspect(node, func(node ast.Node) bool {
		switch e := node.(type) {
		// must be an expression call
		case *ast.CallExpr:
			// must be a function call
			switch fun := e.Fun.(type) {
			case *ast.Ident:
				// check that the function name is ParseAndEval
				if fun.Name == "ParseAndEval" {
					//ensure two elements in the argument list
					if len(e.Args) == 2 {
						e.Args[0] = e.Args[1]
						switch arg := e.Args[0].(type) {
						case *ast.BasicLit:
							asString := arg.Value
							// convert the string to an expression
							expression, err := expr.Parse(asString)
							if err == nil {
								// call simplify to simplify the expression
								//simp := simplify.Simplify(expression, e.Args[1])
								simp := simplify.Simplify(expression, e.Args[1].(expr.Env))

								// convert back to a string(expr.Format)
								str := expr.Format(simp)
								// set the new arg value
								e.Args[0].(*ast.BasicLit).Value = str
							}
						}
					}
				}
			}
		}
		return true
	})
}

func SimplifyParseAndEval(src string) string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	rewriteCalls(f)

	var buf bytes.Buffer
	format.Node(&buf, fset, f)
	return buf.String()
}
