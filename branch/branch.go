package branch

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// bacnchCount should count the number of branching statements in the function.
// package, name and declarations are key (last doccument)
// code parses and gives ast
// go through ast to count branching statements: if, for, switch, range,...

func branchCount(fn *ast.FuncDecl) uint {
	var count uint = 0
	// inspector defined for each node in the AST
	// assume it will go through every node in the AST
	// the function is called every time the inspector visits a node in the tree

	ast.Inspect(fn, func(node ast.Node) bool {
		// get node type 
		switch node.(type){
		// increase for if, for, switch, typeswitch, range 
		case *ast.IfStmt:
			count += 1
		case *ast.ForStmt:
			count += 1
		case *ast.SwitchStmt:
			count += 1
		case *ast.TypeSwitchStmt:
			count += 1
		case *ast.RangeStmt:
			count += 1
		}
		// always return true 
		return true
	})
	return count
}

// ComputeBranchFactors returns a map from the name of the function in the given
// Go code to the number of branching statements it contains.
func ComputeBranchFactors(src string) map[string]uint {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	// f = ast file object, the one you are going to manipulate
	if err != nil {
		panic(err)
	}

	m := make(map[string]uint)
	for _, decl := range f.Decls {
		switch fn := decl.(type) {
		case *ast.FuncDecl:
			m[fn.Name.Name] = branchCount(fn)
		}
	}

	return m
}
