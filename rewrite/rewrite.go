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
	"strconv"
)


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
				// check the package is expr
				if fun.Name == "ParseAndEval" && fun.Package == "expr"{

					// fun.Name = "Changed"

					//ensure two elements in the argument list 
					if len(e.Args) == 2 {

						//e.Args[0] = e.Args[1]

						switch arg := e.Args[0].(type) {
						case *ast.BasicLit:
							// the argument is a literal, but we have to check that it is a string 
							if arg.Kind == token.STRING{
								
								//fmt.Printf(arg.Value)
								// shoud look like: "1 + 2"

								// need to remove the quotes before 
								asString, error = strconv.Unquote(arg.Value)

								if error == nil{
									// parse to expression format 
									expression, err := expr.Parse(asString)
									if err == nil{
										// simplify the expression 
										simp := simplify.Simplify(expression, expr.Env{})

										// format back to string
										str := expr.Format(simp)

										// put the quotes back
										withQuotes := strconv.Quote(str)

										// assign the value back to the argument
										arg.Value = str
									}

								}
							}
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
