package analyzer

import (
	"flag"
	"fmt"
	"go/ast"
	"go/token"

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
			val := checkWords(n)
			fmt.Println(val)
			for _, valx := range val {
				fmt.Println("VALX-", valx)
				v, ok := valx[n.Pos()]
				if ok {
					fmt.Println("POSITION:", n.Pos(), v)

				}
				return true
				//fmt.Println("POSITION:", n.Pos(), valx[n.Pos()])

			}
			if len(val) > 0 {
				return true
				//pass.Reportf(n.Pos(), "Value Of all Words %d", val)
			}

			fmt.Printf("RESULT %s:\t%b\n", pass.Fset.Position(n.Pos()), count)
			return true
		})
		return nil, nil
	}
	return nil, nil
}

type badWord []map[token.Pos]string

type treeVisitor struct {
	badWordArray badWord
}

func checkWords(fn ast.Node) badWord {
	if fn == nil {
		return nil
	}

	v := treeVisitor{}
	ast.Walk(&v, fn)
	return v.badWordArray
}

func isBadWord(word string) bool {
	//Takes in a word, loops through the array of curse words,
	//If the received parameter is in the array return true.
	return true
}
func (v *treeVisitor) addWordToSlice(badWord string, position token.Pos) {
	newBadWordMap := make(map[token.Pos]string)
	newBadWordMap[position] = badWord
	v.badWordArray = append(v.badWordArray, newBadWordMap)
}

func (v *treeVisitor) Visit(n ast.Node) ast.Visitor {
	switch n := n.(type) {
	case *ast.Comment:
		fmt.Println("I am a Comment", n.Text)
	case *ast.CommentGroup:
		fmt.Println("I am a comment Group", n.Text())
	case *ast.FuncType:
		for i, val := range n.Params.List {
			b := isBadWord(val.Names[i].Name)
			if b {
				v.addWordToSlice(val.Names[i].Name, n.Pos())
			}
			//For each parameter, check to see if it is a curse word
			//
		}
	case *ast.FuncDecl:
		fmt.Println("I am a function name-", n.Name.Name)
		//Names[len(n.Recv.List[len(n.Recv.List)-1].Names)]
		//fmt.Println("I am a function object name-2", n.Recv.List[len(n.Recv.List)-1])
	default:
		//fmt.Println("Type Not set")
	}
	return v
}
