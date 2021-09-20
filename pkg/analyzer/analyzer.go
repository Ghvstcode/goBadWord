package analyzer

import (
	"flag"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var (
	badwords     string
	Type         string
	Name         string
	flagSet      flag.FlagSet
	count        int
	skipTests    bool
	userBadWords []string
)

var defaultBadWords = []string{"fuck", "shit", "damn"}

func init() {
	flagSet.BoolVar(&skipTests, "skipTests", true, "should the linter execute on test files as well")
	userBadWords = flagSet.Args()
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

func remove(s badWord, i int) badWord {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func appendWithoutDuplicates(bw badWord, nw map[token.Pos]string) badWord {
	//Range Over badWord slice.
	for _, el := range bw {
		//if bad word is equals to new word return s
		eq := reflect.DeepEqual(el, nw)

		if eq {
			return bw
		}
	}
	//if not, append the new word to bad word slice & return it.
	bw = append(bw, nw)
	return bw
}

//isTestFunc checks if a function is a test function.
func isTestFunc(n ast.Node) bool {
	f, ok := n.(*ast.FuncDecl)
	if !ok {
		return false
	}

	return strings.HasPrefix(f.Name.Name, "Test")
}

func run(pass *analysis.Pass) (interface{}, error) {
	fmt.Println(badwords)
	var s badWord
	var num int
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			_, val := checkWords(n)
			if n == nil {
				return true
			}

			if skipTests && isTestFunc(n) {
				fmt.Println("Hit Mw")
				return true
			}

			num = +len(s)
			for _, valx := range val {
				//Check for duplicates before appending.
				s = appendWithoutDuplicates(s, valx)
			}

			for i, sValx := range s {
				v, ok := sValx[n.Pos()]
				if ok {
					//Removeitem from general array
					s = remove(s, i)
					pass.Reportf(n.Pos(), "Bad Word Found - %s", v)
				}
			}
			return true
		})
	}
	fmt.Printf("RESULT: a total of %d bad words were found", num)
	return nil, nil
}

type badWord []map[token.Pos]string

type treeVisitor struct {
	badWordArray   badWord
	badWordCounter int
}

func checkWords(fn ast.Node) (int, badWord) {
	if fn == nil {
		return 0, nil
	}

	v := treeVisitor{}
	ast.Walk(&v, fn)
	return v.badWordCounter, v.badWordArray
}

func isBadWord(word string) bool {
	//Check if user provided args
	if len(userBadWords) > 0 {
		//If yes, remove file paths/name.
		//Check for strings that contains "/" or ".go"
	}
	//Takes in a word, loops through the array of bad words,
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
	//case *ast.Comment:
	//	fmt.Println("I am a Comment", n.Text)
	//case *ast.CommentGroup:
	//	fmt.Println("I am a comment Group", n.Text())
	case *ast.FuncType:
		val := n.Params
		for i := 0; i < val.NumFields(); i++ {
			for _, x := range val.List {
				for _, n := range x.Names {
					b := isBadWord(n.String())
					if b {
						v.badWordCounter++
						v.addWordToSlice(n.String(), n.Pos())
					}
				}
			}
		}
	//case *ast.FuncDecl:
	//	fmt.Println("I am a function name-", n.Name.Name)
	default:
	}
	return v
}
