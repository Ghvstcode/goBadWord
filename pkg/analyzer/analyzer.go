package analyzer

import (
	"encoding/json"
	"flag"
	"go/ast"
	"go/token"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var (
	flagSet  flag.FlagSet
	badwords string
)

type badWord []map[token.Pos]string

type treeVisitor struct {
	badWordArray badWord
}
type defaultBadWords []string

func loadDefaultBadWords() string {
	// Open our jsonFile
	jsonFile, _ := os.OpenFile("../../badwords.json", os.O_RDWR, 0644)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result defaultBadWords
	_ = json.Unmarshal(byteValue, &result)
	return strings.Join(result, ", ")
}

func init() {
	v := loadDefaultBadWords()
	flagSet.StringVar(&badwords, "bad-words", v, "Specify the bad word the linter should look out for e.g. -badWord=\"fuck, damn, shit\")")
}

func NewAnalyzer() *analysis.Analyzer {
	an := &analysis.Analyzer{
		Name:  "goBadWords",
		Doc:   "Find occurrence of curse words or specified bad words",
		Run:   run,
		Flags: flagSet,
	}
	return an
}

func remove(s badWord, i int) badWord {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func appendWithoutDuplicates(bw badWord, nw map[token.Pos]string) badWord {
	// Range Over badWord slice.
	for _, el := range bw {
		// if bad word is equals to new word return s
		eq := reflect.DeepEqual(el, nw)

		if eq {
			return bw
		}
	}
	// if not, append the new word to bad word slice & return it.
	bw = append(bw, nw)
	return bw
}

func run(pass *analysis.Pass) (interface{}, error) {
	var s badWord
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			val := checkWords(n)

			if n == nil {
				return true
			}

			for _, valx := range val {
				// Check for duplicates before appending.
				s = appendWithoutDuplicates(s, valx)
			}
			for i, sValx := range s {
				v, ok := sValx[n.Pos()]
				if ok {
					// Remove item from general array
					s = remove(s, i)
					pass.Reportf(n.Pos(), "Bad Word Found - %s", v)
				}
			}
			return true
		})
	}
	return nil, nil
}

func checkWords(n ast.Node) badWord {
	if n == nil {
		return nil
	}

	v := treeVisitor{}
	ast.Walk(&v, n)
	return v.badWordArray
}

// isBadWord Takes in a string, loops through the array of bad words,
// If the received parameter is in the array return true.
// Convert String to array that can be looped over.
func isBadWord(word string) bool {
	nbw := strings.Split(badwords, ",")
	for _, rbw := range nbw {
		if word == rbw || strings.Contains(word, rbw) {
			return true
		}

	}
	return false
}

func (v *treeVisitor) addWordToSlice(badWord string, position token.Pos) {
	newBadWordMap := make(map[token.Pos]string)
	newBadWordMap[position] = badWord
	v.badWordArray = append(v.badWordArray, newBadWordMap)
}

func (v *treeVisitor) genericCheckAndAdd(word string, position token.Pos) {
	var b bool
	b = isBadWord(word)
	if b {
		v.addWordToSlice(word, position)
	}
}

// Visit visits all nodes.
func (v *treeVisitor) Visit(n ast.Node) ast.Visitor {
	switch n := n.(type) {
	case *ast.BasicLit:
		v.genericCheckAndAdd(n.Value, n.Pos())
	case *ast.BinaryExpr:
		v.BinaryExpr(n)
	case *ast.BranchStmt:
		v.genericCheckAndAdd(n.Label.Name, n.Pos())
	case *ast.Comment:
		v.Comment(n)
	case *ast.DeferStmt:
		if n.Pos() != token.NoPos {
			v.genericCheckAndAdd("defer", n.Pos())
		}
	case *ast.Ellipsis:
		if n.Pos() != token.NoPos {
			v.genericCheckAndAdd("...", n.Pos())
		}
	case *ast.Field:
		for _, fi := range n.Names {
			v.genericCheckAndAdd(fi.Name, n.Pos())
		}
	case *ast.ForStmt:
		if n.Pos() != token.NoPos {
			v.genericCheckAndAdd("for", n.Pos())
		}
	case *ast.IfStmt:
		if n.Pos() != token.NoPos {
			v.genericCheckAndAdd("if", n.Pos())
		}
	case *ast.InterfaceType:
		if n.Pos() != token.NoPos {
			v.genericCheckAndAdd("interface", n.Pos())
		}
	case *ast.KeyValueExpr:
		if n.Pos() != token.NoPos {
			v.genericCheckAndAdd(":", n.Pos())
		}
	case *ast.MapType:
		if n.Pos() != token.NoPos {
			v.genericCheckAndAdd("map", n.Pos())
		}
	case *ast.Package:
		if n.Pos() != token.NoPos {
			v.genericCheckAndAdd(n.Name, n.Pos())
		}
	case *ast.FuncDecl:
		v.genericCheckAndAdd(n.Name.String(), n.Pos())
	case *ast.GoStmt:
		if n.Pos() != token.NoPos {
			v.genericCheckAndAdd("go", n.Pos())
		}
	case *ast.ReturnStmt:
		if n.Pos() != token.NoPos {
			v.genericCheckAndAdd("return", n.Pos())
		}
	case *ast.SelectStmt:
		if n.Pos() != token.NoPos {
			v.genericCheckAndAdd("select", n.Pos())
		}
	case *ast.StarExpr:
		if n.Pos() != token.NoPos {
			v.genericCheckAndAdd("*", n.Pos())
		}
	case *ast.StructType:
		if n.Pos() != token.NoPos {
			v.genericCheckAndAdd("struct", n.Pos())
		}
	case *ast.SwitchStmt:
		if n.Pos() != token.NoPos {
			v.genericCheckAndAdd("switch", n.Pos())
		}
	case *ast.TypeSpec:
		if n.Pos() != token.NoPos {
			v.genericCheckAndAdd("=", n.Pos())
		}
	case *ast.FuncType:
		v.FuncType(n)
	default:
	}
	return v
}

// For binary expressions, You will want to escape when specifying a bad word in the terminal
func (v *treeVisitor) BinaryExpr(n *ast.BinaryExpr) {
	xv, ok := n.X.(*ast.BasicLit)
	if ok {
		v.genericCheckAndAdd(xv.Value, n.Pos())
	}
	yv, ok := n.Y.(*ast.BasicLit)
	if ok {
		v.genericCheckAndAdd(yv.Value, n.Pos())
	}

}

// Comment handles a node that is a comment type
func (v *treeVisitor) Comment(n *ast.Comment) {
	cm := strings.Split(n.Text, "//")

	for _, cmEl := range cm {
		v.genericCheckAndAdd(cmEl, n.Pos())
	}

}

// FuncType handles a node which is a functype
func (v *treeVisitor) FuncType(n *ast.FuncType) {
	val := n.Params
	for i := 0; i < val.NumFields(); i++ {
		for _, x := range val.List {
			for _, n := range x.Names {
				v.genericCheckAndAdd(n.String(), n.Pos())
			}
		}
	}

}
