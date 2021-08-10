package wordcount

import (
	"flag"
	"fmt"
	"go/ast"
	"strconv"

	"golang.org/x/tools/go/analysis"
)

var (
	Type    string
	Name    string
	flagSet flag.FlagSet
	count   int
)

func init() {
	flagSet.StringVar(&Type, "Type", "string", "Type of the word to be counted")
	flagSet.StringVar(&Name, "Id", "string", "Name of the word to be counted")
}

func NewAnalyzer() *analysis.Analyzer {
	an := &analysis.Analyzer{
		Name:  "addLink",
		Doc:   "reports integer additions",
		Run:   run,
		Flags: flagSet,
	}

	return an
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			// check whether it is  a Binary Expression
			if Type == "string" {
				be, ok := n.(*ast.BasicLit)
				if !ok {
					return true
				}
				v, err := strconv.Unquote(be.Value)
				if v == Name {
					count++
					return true
				}
				if err == nil {
					return true
				}
				fmt.Printf("%s:\t%b\n", pass.Fset.Position(n.Pos()), count)
				return true
			}
			return true
		})
		return nil, nil
	}
	return nil, nil
}

//func run(pass *analysis.Pass) (interface{}, error) {
//	for _, file := range pass.Files {
//		ast.Inspect(file, func(n ast.Node) bool {
//			// check whether it is  a Binary Expression
//			if Type == "string" {
//				be, ok := n.(*ast.BasicLit)
//				if !ok {
//					return true
//				}
//				v, err := strconv.Unquote(be.Value)
//				if err != nil {
//					return true
//				}
//				if v = Name{
//
//				}
//			}
//			be, ok := n.(*ast.BasicLit)
//			if !ok {
//				return true
//			}
//
//			if be.Op != token.ADD {
//				return true
//			}
//
//			if _, ok := be.X.(*ast.BasicLit); !ok {
//				return true
//			}
//
//			if _, ok := be.Y.(*ast.BasicLit); !ok {
//				return true
//			}
//
//			isInteger := func(expr ast.Expr) bool {
//				t := pass.TypesInfo.TypeOf(expr)
//				if t == nil {
//					return false
//				}
//
//				bt, ok := t.Underlying().(*types.Basic)
//				if !ok {
//					return false
//				}
//
//				if (bt.Info() & types.IsInteger) == 0 {
//					return false
//				}
//
//				return true
//			}
//
//			// check that both left and right hand side are integers
//			if !isInteger(be.X) || !isInteger(be.Y) {
//				return true
//			}
//
//			pass.Reportf(be.Pos(), "integer addition found %q",
//				render(pass.Fset, be))
//			return true
//		})
//	}
//	return nil, errors.New("not implemented yet")
//}
